package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
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

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.Run(":" + viper.GetString("SERVER_PORT"))

}
