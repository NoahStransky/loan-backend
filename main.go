package main

import (
	"flag"
	"fmt"

	"bernie.top/demyst/loan-backend/src"
	"github.com/spf13/viper"
)

var configFile string

func main() {
	flag.StringVar(&configFile, "config", "./config/dev.json", "help message for flag")
	flag.Parse()
	viper.SetConfigFile(configFile)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	src.Serve()
}
