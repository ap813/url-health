package utils

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"url-health/scheduler"

	"gopkg.in/yaml.v2"
)

var defaultSleep int = 900

// Take in file path and return map of yaml
func ReadYaml(filepath string) (map[interface{}]interface{}, error) {
	yamlBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Path to file invalid: %v", err))
	}
	data := make(map[interface{}]interface{})
	err = yaml.Unmarshal(yamlBytes, &data)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Failed to read yaml configuration: %v", err))
	}
	return data, nil
}

// Return a slice of url structs from the parsed yaml config
func ParseConfig(data map[interface{}]interface{}) (map[url.URL]scheduler.Status, int, error) {
	results := make(map[url.URL]scheduler.Status)
	var s int
	if sleep, ok := data["sleep"]; ok {
		if sleepUint, ok := sleep.(int); ok {
			s = sleepUint
		} else {
			return nil, 0, errors.New(fmt.Sprintf("Value in config for (sleep) in not valid positive integer. Received: %v", sleep))
		}
	} else {
		s = defaultSleep
	}

	if val, ok := data["urls"]; ok {
		if urls, ok := val.([]interface{}); ok {
			for _, u := range urls {
				// Use URL list and prepend https://
				strURL := fmt.Sprint(u)
				parsedURL, err := url.ParseRequestURI(fmt.Sprintf("https://www.%s", strURL))
				if err != nil {
					return nil, 0, errors.New(fmt.Sprintf("URL (%s) is not a valid format.", strURL))
				}
				// Set all status values as DOWN until first call
				status := scheduler.CheckURL(*parsedURL)
				results[*parsedURL] = status
			}
		} else {
			return nil, 0, errors.New("Error assigning YAML 'urls' to correct type. List should only include strings.")
		}
	} else {
		return nil, 0, errors.New("Error getting value data[urls] in parse config, YAML is misconfigured.")
	}

	return results, s, nil
}
