package cli

import (
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/home-IoT/home-weather/gen/client"
	"github.com/home-IoT/home-weather/gen/client/operations"
	"github.com/home-IoT/home-weather/internal/config"
	"github.com/home-IoT/home-weather/internal/log"
	"net/url"
)

// ListSensors lists the available sensors
func ListSensors() {
	httpClient := newSensorsHTTPClient()
	params := operations.NewGetSensorsListParams()
	resp, err := httpClient.Operations.GetSensorsList(params)
	if err != nil {
		log.Debugf("%v", err)
		log.Exitf(1, "Cannot access Jupiter.")
	}

	list := resp.Payload.Data

	if len(list) == 0 {
		fmt.Println("sensors: []")
	} else {
		fmt.Println("sensors:")
		for _, card := range list {
			fmt.Printf("  - %s\n", *card.ID)
		}
	}
}

func newSensorsHTTPClient() *client.Jupiter {
	jURL, err := url.Parse(config.GetJupiterURL())
	if err != nil {
		log.Debugf("%v", err)
		log.Exitf(1, "The Jupiter URL is not valid.")
	}

	transport := httptransport.New(jURL.Host, jURL.Path, nil)
	return client.New(transport, strfmt.Default)
}
