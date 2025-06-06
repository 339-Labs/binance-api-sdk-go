package v

import (
	"github.com/339-Labs/binance-api-sdk-go/config"
	"github.com/339-Labs/binance-api-sdk-go/constants"
	"github.com/339-Labs/binance-api-sdk-go/internal/common"
)

type MixMarketClient struct {
	BnRestClient *common.BnRestClient
}

func (p *MixMarketClient) Init(config *config.BnConfig) *MixMarketClient {
	p.BnRestClient = new(common.BnRestClient).Init(config, string(constants.SWAP))
	return p
}

func (p *MixMarketClient) Tickers(params map[string]string) (string, error) {
	resp, err := p.BnRestClient.DoGet(false, "/fapi/v2/ticker/price", params)
	return resp, err
}
