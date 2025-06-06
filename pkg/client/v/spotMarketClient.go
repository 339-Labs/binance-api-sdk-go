package v

import (
	"github.com/339-Labs/binance-api-sdk-go/config"
	"github.com/339-Labs/binance-api-sdk-go/constants"
	"github.com/339-Labs/binance-api-sdk-go/internal/common"
)

type SpotMarketClient struct {
	BnRestClient *common.BnRestClient
}

func (p *SpotMarketClient) Init(config *config.BnConfig) *SpotMarketClient {
	p.BnRestClient = new(common.BnRestClient).Init(config, string(constants.Spot))
	return p
}

func (p *SpotMarketClient) Tickers(params map[string]string) (string, error) {
	resp, err := p.BnRestClient.DoGet(false, "/api/v3/ticker/price", params)
	return resp, err
}
