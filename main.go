package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	viper.AddConfigPath(".")
	viper.SetConfigName("sample") // Register config file name (no extension) only name for e.g sample
	viper.SetConfigType("json")   // Look for specific type
	viper.ReadInConfig()

	var config ConfigExample

	viper.Unmarshal(&config)

	fmt.Println(config.CompilerOptions.Target)
}

type ConfigExample struct {
	CompilerOptions struct {
		Module string `json:"module"`
		Target string `json:"target"`
	} `json:"compilerOptions"`
	Exclude []string `json:"exclude"`
}
