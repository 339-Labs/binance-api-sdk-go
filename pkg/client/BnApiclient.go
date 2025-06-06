package client

import (
	"github.com/339-Labs/binance-api-sdk-go/config"
	"github.com/339-Labs/binance-api-sdk-go/internal/common"
)

type BnApiClient struct {
	BnRestClient *common.BnRestClient
}

func (p *BnApiClient) Init(config *config.BnConfig, intType string) *BnApiClient {
	p.BnRestClient = new(common.BnRestClient).Init(config, intType)
	return p
}

func (p *BnApiClient) Post(signature bool, url string, params map[string]string) (string, error) {
	resp, err := p.BnRestClient.DoPost(signature, url, params)
	return resp, err
}

func (p *BnApiClient) Get(signature bool, url string, params map[string]string) (string, error) {
	resp, err := p.BnRestClient.DoGet(signature, url, params)
	return resp, err
}
