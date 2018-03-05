package cli

import (
	"fmt"
	"math"
	"net/url"
	"sort"
	"strings"
	"sync"

	"github.com/home-IoT/home-weather/internal/config"

	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/home-IoT/home-weather/gen/client"
	"github.com/home-IoT/home-weather/gen/client/operations"
	"github.com/home-IoT/home-weather/gen/models"
	"github.com/home-IoT/home-weather/internal/log"
	"gopkg.in/yaml.v2"
)

const fiveMinutesInMillis = 5 * 60 * 1000

var jURLOnce sync.Once
var jupiterURL *url.URL

type sensorData struct {
	ID          string  `yaml:"id,omitempty"`
	Name        string  `json:"name"`
	Temperature float64 `yaml:"temperature"`
	Humidity    float64 `yaml:"humidity"`
	HeatIndex   float64 `yaml:"heatIndex,omitempty"`
	Stale       bool    `yaml:"stale,omitempty"`
}

type readingStats struct {
	Median  float64 `yaml:"median"`
	Average float64 `yaml:"average"`
	Min     float64 `yaml:"min"`
	Max     float64 `yaml:"max"`
}

type readingSummary struct {
	Count     int          `yaml:"count"`
	TempStats readingStats `yaml:"temperature"`
	HumStats  readingStats `yaml:"humidity"`
}

type readingResults struct {
	Readings []*sensorData   `yaml:"reading"`
	Summary  *readingSummary `yaml:"summary,omitempty"`
}

// ListSensors lists the available sensors
func ListSensors() {
	list := getAllSensorIDs()

	if len(list) == 0 {
		fmt.Println("sensors: []")
	} else {
		fmt.Println("sensors:")
		for _, id := range list {
			fmt.Printf("  - %s\n", id)
		}
	}
}

// ReadSensors reads data from the sensors and prints the results
func ReadSensors(sensorList string, fullOutput bool, allSensors bool) {
	sensorList = strings.TrimSpace(sensorList)

	var sensorIDs []string
	switch {
	case allSensors:
		sensorIDs = getAllSensorIDs()

	case sensorList == "":
		log.Exitf(1, "No valid sensor ID is given.")

	default:
		sensorIDs = strings.Split(sensorList, ",")
	}

	ids := map[string]bool{}
	dataChannel := make(chan *models.SensorData)
	for _, id := range sensorIDs {
		if _, ok := ids[id]; !ok {
			ids[id] = true
			go readSensor(id, dataChannel)
		}
	}

	sensorDataList := []*sensorData{}
	for i := 0; i < len(ids); i++ {
		sensorData := <-dataChannel
		if sensorData != nil {
			csd := consumeSensorData(sensorData)
			if !fullOutput {
				csd.ID = ""
			}
			sensorDataList = append(sensorDataList, csd)
		}
	}

	if len(sensorDataList) == 0 {
		log.Debugf("Did not read any sensor data.")
		return
	}

	results := processReadings(sensorDataList, fullOutput)

	if yamlData, err := yaml.Marshal(&results); err != nil {
		log.Debugf("%v", err)
		log.Fatalf("Cannot encode summary data as YAML.")
	} else {
		fmt.Printf("%s", yamlData)
	}
}

func getAllSensorIDs() []string {
	httpClient := newSensorsHTTPClient()
	params := operations.NewGetSensorsListParams()
	resp, err := httpClient.Operations.GetSensorsList(params)
	handleJupiterError(err)

	list := resp.Payload.Data
	ids := make([]string, len(list))

	for i, card := range list {
		ids[i] = *card.ID
	}

	return ids
}

func processReadings(sensorDataList []*sensorData, fullOutput bool) readingResults {
	results := readingResults{}
	results.Readings = sensorDataList
	if fullOutput && len(sensorDataList) > 1 {
		results.Summary = &readingSummary{}
		results.Summary.Count = len(sensorDataList)
		results.Summary.HumStats.Max = 0
		results.Summary.HumStats.Min = 100
		results.Summary.TempStats.Min = 1000
		results.Summary.TempStats.Max = 0

		temps := make([]float64, len(sensorDataList))
		humids := make([]float64, len(sensorDataList))

		for index, sd := range sensorDataList {
			temps[index] = sd.Temperature
			humids[index] = sd.Humidity

			// temp summary
			if sd.Temperature < results.Summary.TempStats.Min {
				results.Summary.TempStats.Min = sd.Temperature
			}
			if sd.Temperature > results.Summary.TempStats.Max {
				results.Summary.TempStats.Max = sd.Temperature
			}
			results.Summary.TempStats.Average += sd.Temperature

			// humidity summary
			if sd.Humidity < results.Summary.HumStats.Min {
				results.Summary.HumStats.Min = sd.Humidity
			}
			if sd.Humidity > results.Summary.HumStats.Max {
				results.Summary.HumStats.Max = sd.Humidity
			}
			results.Summary.HumStats.Average += sd.Humidity
		}

		results.Summary.TempStats.Average /= float64(len(sensorDataList))
		results.Summary.TempStats.Average = math.Trunc(results.Summary.TempStats.Average*100) / 100
		results.Summary.HumStats.Average /= float64(len(sensorDataList))
		results.Summary.HumStats.Average = math.Trunc(results.Summary.HumStats.Average*100) / 100

		results.Summary.TempStats.Median = median(temps)
		results.Summary.HumStats.Median = median(humids)
	}

	return results
}

func median(data []float64) float64 {
	sort.Float64s(data)

	var result float64

	switch {
	case len(data) == 1:
		result = data[0]

	case len(data)%2 == 0:
		j := len(data) / 2
		i := j - 1
		result = (data[i] + data[j]) / 2

	default:
		i := len(data) / 2
		result = data[i]
	}

	result = math.Trunc(result*100) / 100

	return result
}

func readSensor(sensorID string, c chan *models.SensorData) {
	httpClient := newSensorsHTTPClient()
	params := operations.NewGetSensorDataParams()
	params.SensorID = strings.TrimSpace(sensorID)
	resp, err := httpClient.Operations.GetSensorData(params)
	handleJupiterError(err)
	c <- resp.Payload.Data
}

func newSensorsHTTPClient() *client.Jupiter {
	jURLOnce.Do(initJupiterURL)
	transport := httptransport.New(jupiterURL.Host, jupiterURL.Path, nil)
	return client.New(transport, strfmt.Default)
}

func initJupiterURL() {
	if jupiterURL == nil {
		jupiterURL = config.GetJupiterURL()
	}
}

func handleJupiterError(err error) {
	if err != nil {
		log.Debugf("%v", err)
		switch err.(type) {
		case *url.Error:
			log.Exitf(1, "Cannot access Jupiter.")

		default:
			log.Exitf(1, "Erorr reading sensor data.")
		}
	}
}

func consumeSensorData(data *models.SensorData) *sensorData {
	if data == nil {
		return nil
	}

	result := sensorData{}

	result.ID = *data.ID
	result.Name = *data.Name
	result.Stale = (*data.DeltaTime > fiveMinutesInMillis)
	result.Temperature = *data.Temperature
	result.Humidity = *data.Humidity
	result.HeatIndex = math.Trunc(data.HeatIndex*100) / 100

	return &result
}
