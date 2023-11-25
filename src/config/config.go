package config

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func (c *Config) GetWorkflow(w string) (YAMLWorkflow, error) {
	for name, workflow := range c.Workflows {
		if name == w {
			return workflow, nil
		}
	}
	return YAMLWorkflow{}, fmt.Errorf("workflow %s not found", w)
}

func (c *Config) GetImage(i string) (YAMLImage, error) {
	for name, image := range c.Images {
		if name == i {
			return image, nil
		}
	}
	return YAMLImage{}, fmt.Errorf("image %s not found", i)
}

func (c *Config) GetTest(t string) (YAMLTest, error) {
	for name, test := range c.Tests {
		if name == t {
			return test, nil
		}
	}
	return YAMLTest{}, fmt.Errorf("test %s not found", t)
}

func (c *Config) GetWorkflowNames() []string {
	var names []string
	for name := range c.Workflows {
		names = append(names, name)
	}
	return names
}

func LoadConfig() (*Config, error) {
	path, ok := os.LookupEnv("VELOCITY_CONFIG")
	if !ok {
		path = "velocity.yml"
	}
	// Switch statement to find out what filepath starts with
	// if filepath starts with http:// or https://
	// then use ReadConfigFromURL
	// else use ReadConfigFromFile
	switch {
	case len(path) > 8 && (path[:7] == "http://" || path[:8] == "https://"):
		return ReadConfigFromURL(path)
	default:
		return ReadConfigFromFile(path)
	}
}

func ReadConfigFromURL(url string) (*Config, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve file '%s': status code '%d' error '%v'", url, response.StatusCode, err)
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return ParseConfig(bytes, NewMultiParser(&YAMLParser{}, &JSONParser{}))

}

func ReadConfigFromFile(filepath string) (*Config, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file '%s': %v", filepath, err)
	}
	return ParseConfig(file, NewMultiParser(&YAMLParser{}, &JSONParser{}))
}

func ParseConfig(config []byte, parser MultiParser) (*Config, error) {
	c := &Config{}
	err := parser.Parse(config, c)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML file: %v", err)
	}
	err = c.Populate()
	if err != nil {
		return nil, fmt.Errorf("error populating YAML file: %v", err)
	}
	err = c.Validate()
	if err != nil {
		return nil, fmt.Errorf("error validating YAML file: %v", err)
	}
	return c, nil
}
