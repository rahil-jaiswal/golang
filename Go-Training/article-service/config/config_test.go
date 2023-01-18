package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateConfig(t *testing.T) {
	filename := "config.json"
	configuration := TranslateConfig(filename)
	if configuration.Debug != true {
		t.Error("Expected", true, "Got", configuration.Debug)
	}
	if configuration.Database.Port != "3306" {
		t.Error("Expected", 3306, "Got", configuration.Database.Port)
	}
}

func TestTranslateConfig_2(t *testing.T) {
	filename := "../config.json"
	configuration := TranslateConfig(filename)
	assert.Nil(t, configuration)
}

func TestTranslateConfig_3(t *testing.T) {
	sampleConfig := Configuration{
		Debug: true,
		Server: Server{
			Address: ":9090",
		},
		Context: Context{
			Timeout: 2,
		},
		Database: Database{
			Host:     "localhost",
			Port:     "3306",
			User:     "root",
			Password: "password",
			Name:     "article",
		},
	}
	filename := "config.json"
	configuration := TranslateConfig(filename)
	assert.Equal(t, sampleConfig, *configuration)
}
