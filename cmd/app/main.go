package main

import (
	"log"

	"github.com/iamthe1whoknocks/bazilikgroup_test_task/config"
	"github.com/iamthe1whoknocks/bazilikgroup_test_task/pkg/routing"
	"github.com/spf13/viper"
)

func main() {
	config := config.New()
	err := config.Init()
	if err != nil {
		log.Fatal("Cant init config...", err)

	}

	err = routing.Run(viper.GetString("host"), viper.GetString("port"))
	if err != nil {
		log.Fatal("Routing cant start : ", err)
	}

}
