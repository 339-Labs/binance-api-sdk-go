package constants

type InstType string

const (
	/*
	 * http headers
	 */
	ContentType     = "Content-Type"
	BnAccessKey     = "X-MBX-APIKEY"
	ApplicationJson = "application/json"
	ApplicationFrom = "application/x-www-form-urlencoded"

	EN_US  = "en_US"
	ZH_CN  = "zh_CN"
	LOCALE = "locale="

	/*
	 * http methods
	 */
	GET  = "GET"
	POST = "POST"

	/*
	 * websocket
	 */
	WsAuthMethod        = "GET"
	WsAuthPath          = "session.logon"
	WsOpLogin           = "login"
	WsOpUnsubscribe     = "unsubscribe"
	WsOpSubscribe       = "subscribe"
	TimerIntervalSecond = 5
	ReconnectWaitSecond = 60

	/*
	 * SignType
	 */
	RSA    = "RSA"
	SHA256 = "SHA256"

	/*
	   * 产品类型
	   SPOT：币币
	   MARGIN：币币杠杠
	   SWAP：永续合约
	   FUTURES：交割合约
	   OPTION：期权
	*/
	Spot    InstType = "SPOT"
	MARGIN  InstType = "MARGIN"
	SWAP    InstType = "SWAP"
	Inverse InstType = "FUTURES"
	Option  InstType = "OPTION"
)
