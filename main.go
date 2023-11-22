package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	fmt.Println("port is : ", viper.Get("PORT"))
	fmt.Println("Username is : ", viper.Get("USER_NAME"))
}
