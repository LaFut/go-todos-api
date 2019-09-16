package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"todos-rest/app"
)

func readConfig() {
	var err error

	viper.SetConfigFile("default.env")
	viper.SetConfigType("props")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		fmt.Println("WARNING: file .env not found")
	} else {
		viper.SetConfigFile(".env")
		viper.SetConfigType("props")
		err = viper.MergeInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// Read parameters from
	for _, key := range viper.AllKeys() {
		viper.BindEnv(key)
	}
}

func main() {
	readConfig()

	var app = app.NewApp()
	app.Init()
	app.Run()
}
