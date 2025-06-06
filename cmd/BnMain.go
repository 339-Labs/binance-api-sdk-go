package main

import (
	"fmt"
	"github.com/339-Labs/binance-api-sdk-go/config"
)

func main() {
	config := config.NewBnConfig(config.BnApiKey, config.BnApiSecretKey, 5000, "")
	fmt.Println("Press ENTER to unsubscribe and stop...", config)
	fmt.Scanln()
}
