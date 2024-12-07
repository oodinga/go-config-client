package goconfig

import (
	"reflect"
	"testing"
)

func TestUnmarshalConfig(t *testing.T) {
	// Test Case 1: Valid JSON
	validJSON := []byte(`{"name": "test", "port": 8080}`)

	expectedConfig := &Config{
		Name:     "",
		Profiles: []string{"local,dev"},
		Label:    "",
		Version:  "",
		PropertySources: []PropertySource{
			{
				Name: "",
				Source: map[string]interface{}{
					"server.port": "8080",
				},
			},
		},
	}

	config, err := UnmarshalConfig(validJSON)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if !reflect.DeepEqual(config, expectedConfig) {
		t.Errorf("expected %v, got %v", expectedConfig, config)
	}

	// Test Case 2: Invalid JSON
	invalidJSON := []byte(`{"name": "test", "port": "invalid"}`) // invalid value for 'port'

	_, err = UnmarshalConfig(invalidJSON)
	if err == nil {
		t.Errorf("expected an error, but got none")
	}
}
