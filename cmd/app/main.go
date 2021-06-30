package main

import (
	"log"

	"github.com/iamthe1whoknocks/bazilikgroup_test_task/config"
)

func main() {
	config := config.New()
	err := config.Init()
	if err != nil {
		log.Fatal("Cant init config...", err)

	}

}
