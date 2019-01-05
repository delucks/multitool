package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// This converts the yaml.v2 map[interface{}]interface{} to a map[string]interface{} that can be understood by json
// Shamelessly copied from SO, https://stackoverflow.com/questions/40737122/convert-yaml-to-json-without-struct
func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}

func YamlToJson(_ []string, stdin io.Reader) error {
	d := yaml.NewDecoder(stdin)
	e := json.NewEncoder(os.Stdout)
	var data interface{}
	for {
		err := d.Decode(&data)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("Couldn't decode YAML: %v", err)
		}
		err = e.Encode(convert(data))
		if err != nil {
			return fmt.Errorf("Couldn't encode to JSON: %v", err)
		}
	}
	return nil
}

func JsonToYaml(_ []string, stdin io.Reader) error {
	d := json.NewDecoder(stdin)
	e := yaml.NewEncoder(os.Stdout)
	var data interface{}
	for {
		err := d.Decode(&data)
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("Couldn't decode JSON: %v", err)
		}
		err = e.Encode(convert(data))
		if err != nil {
			return fmt.Errorf("Couldn't encode to YAML: %v", err)
		}
	}
	return nil
}
