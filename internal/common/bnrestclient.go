package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/339-Labs/binance-api-sdk-go/config"
	"github.com/339-Labs/binance-api-sdk-go/constants"
	"github.com/339-Labs/binance-api-sdk-go/internal"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type BnRestClient struct {
	ApiKey       string
	ApiSecretKey string
	BaseUrl      string
	HttpClient   http.Client
	Signer       *Signer
	SignType     string
}

func (p *BnRestClient) Init(config *config.BnConfig, intType string) *BnRestClient {
	p.ApiKey = config.ApiKey
	p.ApiSecretKey = config.SecretKey
	p.BaseUrl = config.BaseUrl
	if intType == string(constants.Spot) {
		p.BaseUrl = config.SpotUrl
	} else if intType == string(constants.SWAP) {
		p.BaseUrl = config.SwapUrl
	}
	p.Signer = new(Signer).Init(config.SecretKey)
	p.HttpClient = http.Client{
		Timeout: time.Duration(config.TimeoutSecond) * time.Second,
	}
	p.SignType = config.SignType
	return p
}

func (p *BnRestClient) DoPost(signature bool, uri string, params map[string]string) (string, error) {
	body := internal.BuildGetParams(params)
	//fmt.Println(body)

	if signature && len(params) > 0 {
		mac := hmac.New(sha256.New, []byte(p.ApiSecretKey))
		_, err := mac.Write([]byte(body))
		if err != nil {
			return "", nil
		}
		v := url.Values{}
		v.Set("signature", fmt.Sprintf("%x", (mac.Sum(nil))))
		body = fmt.Sprintf("%s&%s", body, v.Encode())
	}
	requestUrl := p.BaseUrl + uri + body

	buffer := strings.NewReader("{}")
	request, err := http.NewRequest(constants.POST, requestUrl, buffer)

	internal.Headers(request, p.ApiKey)
	if err != nil {
		return "", err
	}
	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}

func (p *BnRestClient) DoGet(signature bool, uri string, params map[string]string) (string, error) {
	body := internal.BuildGetParams(params)
	//fmt.Println(body)

	if signature && len(params) > 0 {
		mac := hmac.New(sha256.New, []byte(p.ApiSecretKey))
		_, err := mac.Write([]byte(body))
		if err != nil {
			return "", nil
		}
		v := url.Values{}
		v.Set("signature", fmt.Sprintf("%x", (mac.Sum(nil))))
		body = fmt.Sprintf("%s&%s", body, v.Encode())
	}

	requestUrl := p.BaseUrl + uri + body
	request, err := http.NewRequest(constants.GET, requestUrl, nil)
	if err != nil {
		return "", err
	}
	internal.Headers(request, p.ApiKey)

	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}
