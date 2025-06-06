package test

import (
	"github.com/339-Labs/binance-api-sdk-go/config"
	"github.com/339-Labs/binance-api-sdk-go/internal"
	"github.com/339-Labs/binance-api-sdk-go/pkg/client/v"
	"testing"
)

func Test_MixTickers(t *testing.T) {
	cf := config.NewBnConfig(config.BnApiKey, config.BnApiSecretKey, 5000, "")
	sp := new(v.MixMarketClient).Init(cf)
	params := internal.NewParams()
	rsp, err := sp.Tickers(params)
	if err != nil {
		t.Error(err)
	}
	t.Log(rsp)
}
