package main

import (
	"flag"
	"log"
	urlutil "net/url"
	"url-health/router"
	"url-health/scheduler"
	"url-health/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Check for flags
	configPath := flag.String("path", "", "A path variable to the configuration file for URLs")
	flag.Parse()

	var sleep int = 900
	if *configPath != "" {
		config, err := utils.ReadYaml(*configPath)
		if err != nil {
			log.Fatalln(err.Error())
		}

		list, s, err := utils.ParseConfig(config)
		if err != nil {
			log.Fatalln(err.Error())
		}

		scheduler.SetList(list)
		sleep = s
	} else {
		var emptyMap = make(map[urlutil.URL]scheduler.Status)
		scheduler.SetList(emptyMap)
	}

	scheduler.MakeScheduler(sleep)

	r := gin.Default()
	r = router.DefineRoutes(r)
	r.Run()
}
