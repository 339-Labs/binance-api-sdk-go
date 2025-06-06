package config

import "github.com/339-Labs/binance-api-sdk-go/constants"

type BnConfig struct {
	BaseUrl       string
	SpotUrl       string
	SwapUrl       string
	WsUrl         string
	SpotWsUrl     string
	SwapWsUrl     string
	ApiKey        string
	SecretKey     string
	TimeoutSecond int
	SignType      string // 可选: "HMAC_SHA256" or "RSA"
}

func NewBnConfig(ApiKey string, SecretKey string, TimeoutSecond int, SignType string) *BnConfig {
	if SignType == "" {
		SignType = constants.SHA256
	}
	return &BnConfig{
		BaseUrl:       "https://api.bitget.com",
		SpotUrl:       "https://api1.binance.com",
		SwapUrl:       "https://fapi.binance.com",
		WsUrl:         "wss://stream.binance.com:443",
		SpotWsUrl:     "wss://stream.binance.com:9443",
		SwapWsUrl:     "wss://ws-fapi.binance.com/ws-fapi/v1",
		ApiKey:        ApiKey,
		SecretKey:     SecretKey,
		TimeoutSecond: TimeoutSecond,
		SignType:      SignType,
	}
}
