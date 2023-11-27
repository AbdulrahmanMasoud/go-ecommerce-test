package main

import (
	"ecommerce/config"
	"fmt"
	"github.com/spf13/viper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	configs := configSet()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app_name": viper.Get("App.Name"),
		})
	})
	r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// Set the configurations from the config folder and
func configSet() config.Config {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("config") // path to look for the config file in
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("We can't read the configs")
	}

	// Unmarshal the configs from the Config struct to be able to read
	var configs config.Config
	err := viper.Unmarshal(&configs)
	if err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	return configs
}
