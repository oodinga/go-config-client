// GoConfig is inspired by sprin-boot config module.It allows developers to use external configs.
package goconfig

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

// Config represents the configuration object
type Config struct {
	Name            string           `json:"name"`
	Profiles        []string         `json:"profiles"`
	Label           interface{}      `json:"label"`
	Version         string           `json:"version"`
	State           interface{}      `json:"state"`
	PropertySources []PropertySource `json:"propertySources"`
}

type PropertySource struct {
	Name   string                 `json:"name"`
	Source map[string]interface{} `json:"source"`
}

type ConfigSettings struct {
	Profiles  []string
	ServerURL string
	AppName   string
	Optional  bool
}

var settings ConfigSettings

// The Load function loads configs from the remote config server. It should be called as early as possible in your application.
// The call is also done by the autoload package. This is to allow ease of use.
func Load() {
	settings = ConfigSettings{
		Profiles:  strings.Split(os.Getenv("app.config.profiles.active"), ","),
		AppName:   os.Getenv("app.name"),
		ServerURL: os.Getenv("app.config.server.url"),
		Optional:  os.Getenv("app.config.optional") == "true",
	}
	log.Printf("Active profiles [%v]", strings.Join(settings.Profiles, ","))
	loadConfigs()
}

// loadConfigs loops through the Profiles set in the environment variables and loads its configs for
// the application. If an error occours, it is reported via #reportError(err error) function
func loadConfigs() {
	for _, profile := range settings.Profiles {
		var err error
		config, err := fetchConfigs(profile)

		if err != nil {
			reportError(err)
			continue
		}
		setEnvVariables(config)
	}
}

// Helper function used to check log errors that occurr during fetching of configs.
// If the configs was set as optional, the errors will be excused, otherwise an exit code of 1 is executed
// and the function is exited
func reportError(err error) {
	if !settings.Optional {
		log.Fatal(err)
	}

	log.Print("Error fetching config:", err)
}

// Perharps the most critical fuction, it receives a config object, and uses the ,
// source objects of the Soource objects values to set environment variables
func setEnvVariables(config *Config) {
	for _, property := range config.PropertySources {
		for key, value := range property.Source {
			os.Setenv(key, fmt.Sprint(value))
		}
	}
}

// fetchConfigs is a helper function used to fetch configs from a config server of a given profile
// IT receives profile and returns a Config object or an error
func fetchConfigs(profile string) (*Config, error) {
	log.Printf("Fetching config from %s/%s/%s\n", settings.ServerURL, settings.AppName, profile)
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s", settings.ServerURL, settings.AppName, profile))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return UnmarshalConfig(body)
}

// Reds bytes and Unmarshals into Config object.
// The service makes use of json.Unmarshal package.
func UnmarshalConfig(data []byte) (*Config, error) {
	var r Config
	err := json.Unmarshal(data, &r)
	return &r, err
}
