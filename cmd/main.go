package main

import (
	"LaZhStudyEnglishWords/config"
	"fmt"
	"log"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//ctx := context.Background()

	// config
	Config, err := config.GetConfig()
	fmt.Println(Config, err)

	return nil
}
