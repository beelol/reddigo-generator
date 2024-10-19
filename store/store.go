package store

import (
	"log"

	"github.com/spf13/viper"
)

func LoadProgress() {
	viper.SetConfigName("progress")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No previous progress found, starting fresh.")
	}
}

func SaveProgress() {
	err := viper.WriteConfigAs("progress.json")
	if err != nil {
		log.Fatal("Error saving progress: ", err)
	}
}

func MarkCompleted(endpointID string) {
	viper.Set(endpointID, true)
}

func GetEndpoints() []string {
	var endpoints []string
	for key := range viper.AllSettings() {
		if !viper.GetBool(key) {
			endpoints = append(endpoints, key)
		}
	}
	return endpoints
}
