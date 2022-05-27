package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	urlutil "net/url"
	"url-health/router"
	"url-health/scheduler"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

// Take in file path and return map of yaml
func readYaml(filepath string) (map[interface{}]interface{}, error) {
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
func parseConfig(data map[interface{}]interface{}) (map[urlutil.URL]scheduler.Status, error) {
	if val, ok := data["urls"]; ok {
		if urls, ok := val.([]interface{}); ok {
			results := make(map[urlutil.URL]scheduler.Status)
			for _, url := range urls {
				// Use URL list and prepend https://
				strURL := fmt.Sprint(url)
				parsedURL, err := urlutil.ParseRequestURI(fmt.Sprintf("https://www.%s", strURL))
				if err != nil {
					return nil, errors.New(fmt.Sprintf("URL (%s) is not a valid format.", strURL))
				}
				// Set all status values as DOWN until first call
				status := scheduler.CheckURL(*parsedURL)
				results[*parsedURL] = status
			}
			log.Printf("%v\n", results)
			return results, nil
		} else {
			return nil, errors.New("Error assigning YAML 'urls' to correct type. List should only include strings.")
		}
	} else {
		return nil, errors.New("Error getting value data[urls] in parse config, YAML is misconfigured.")
	}
}

func main() {
	// Check for flags
	configPath := flag.String("path", "", "A path variable to the configuration file for URLs")
	flag.Parse()

	if *configPath != "" {
		config, err := readYaml(*configPath)
		if err != nil {
			log.Fatalln(err.Error())
		}

		list, err := parseConfig(config)
		if err != nil {
			log.Fatalln(err.Error())
		}

		scheduler.SetList(list)
	} else {
		var emptyMap = make(map[urlutil.URL]scheduler.Status)
		scheduler.SetList(emptyMap)
	}

	fmt.Printf("List available %v\n", scheduler.GetList())

	r := gin.Default()
	r = router.DefineRoutes(r)
	r.Run()
}
