package config

import (
	"github.com/home-IoT/home-weather/internal/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

const configDirName = ".home-weather"
const defaultConfigFileName = "config.yml"

// WeatherConfig holds the configuration of the weather CLI
type WeatherConfig struct {
	JupiterURL string `yaml:"jupiterURL,omitempty"`
}

// SetJupiterURL sets the URL of the Jupiter service in the configuration file
func SetJupiterURL(url string) {
	config := WeatherConfig{}
	readConfigIfExists(&config)
	config.JupiterURL = url
	writeWeatherConfig(&config)
}

// GetJupiterURL returns the currently configured URL of the Jupiter service
func GetJupiterURL() string {
	config := WeatherConfig{}
	readConfigIfExists(&config)
	return config.JupiterURL
}

func writeWeatherConfig(data *WeatherConfig) {
	writeConfig(data, defaultConfigFileName)
}

func writeConfig(data interface{}, configName string) {
	if yamlData, err := yaml.Marshal(&data); err != nil {
		log.Fatalf("cannot encode config data to yaml <%e>", err)
	} else {
		homeWeatherPath := ensureConfigFolderExists(nil)

		if err := ioutil.WriteFile(getConfigFilePath(configName), yamlData, 0600); err != nil {
			log.Fatalf("cannot write config to %s <%e>", homeWeatherPath, err)
		}
	}
}

func ensureConfigFolderExists(folder *string) string {
	var homeWeatherPath string
	if folder == nil {
		homeWeatherPath = locateConfigFolder()
	} else {
		homeWeatherPath = *folder
	}

	if _, err := os.Stat(homeWeatherPath); os.IsNotExist(err) {
		if os.Mkdir(homeWeatherPath, 0700) != nil {
			log.Exitf(1, "Cannot create %s.\nError: %e", homeWeatherPath, err)
		}
	}

	return homeWeatherPath
}

func getConfigFilePath(fileName string) string {
	if filepath.IsAbs(fileName) {
		return fileName
	}

	return path.Join(locateConfigFolder(), fileName)
}

func readConfigIfExists(data *WeatherConfig) {
	homeWeatherPath := getConfigFilePath(defaultConfigFileName)
	if content, err := ioutil.ReadFile(homeWeatherPath); err == nil {
		yaml.Unmarshal(content, data)
	}
}

func locateConfigFolder() string {
	switch {

	case fileExists("/home-weather"):
		//from within docker container
		return "/" + configDirName

	default:
		configPath := path.Join(os.Getenv("HOME"), configDirName)
		ensureConfigFolderExists(&configPath)
		return configPath
	}
}

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
