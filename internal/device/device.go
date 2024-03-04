package device

import (
	"errors"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type DeviceConfig struct {
	Details        DeviceDetails   `yaml:"details"`
	SerialSettings SerialSettings  `yaml:"config"`
	Commands       []DeviceCommand `yaml:"commands"`
}

type DeviceDetails struct {
	Make  string `yaml:"make"`
	Model string `yaml:"model"`
	Type  string `yaml:"type"`
}

type SerialSettings struct {
	BaudRate int     `yaml:"baudRate"`
	DataBits int     `yaml:"dataBits"`
	StopBits float32 `yaml:"stopBits"`
	Parity   string  `yaml:"parity"`
}

type DeviceCommand struct {
	Command string `yaml:"command"`
	Details string `yaml:"details"`
}

func LoadDeviceConfigs() ([]DeviceConfig, error) {
	var configs []DeviceConfig
	var libPath = "./internal/device/lib"

	// Find all YAML files in the lib folder
	files, err := filepath.Glob(filepath.Join(libPath, "*.yaml"))
	if err != nil {
		return nil, err
	}

	// Load and unmarshal each file
	for _, file := range files {
		var config DeviceConfig
		data, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal(data, &config)
		if err != nil {
			return nil, err
		}

		configs = append(configs, config)
	}

	return configs, nil
}

func GetConfigForMakeAndModel(make string, model string) (*DeviceConfig, error) {
	configs, err := LoadDeviceConfigs()
	if err != nil {
		return nil, err
	}

	for _, config := range configs {
		if config.Details.Make == make && config.Details.Model == model {
			return &config, nil
		}
	}

	return nil, errors.New("no matching device config found")
}
