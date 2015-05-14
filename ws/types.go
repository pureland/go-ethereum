package ws

import "encoding/json"

type Config struct {
	ListenAddress string
	ListenPort    uint
	SecurityToken string
}

type WSRequest struct {
	Id        interface{}     `json:"id"`
	WsVersion string          `json:"wsjson"`
	Method    string          `json:"method"`
	Params    json.RawMessage `json:"params"`
}

type WSSuccessResponse struct {
	Id        interface{} `json:"id"`
	WsVersion string      `json:"wsjson"`
	Method    string      `json:"method"`
	Result    interface{} `json:"result"`
}

type WSErrorResponse struct {
	Id        interface{}    `json:"id"`
	WsVersion string         `json:"wsjson"`
	Method    string         `json:"method"`
	Error     *WSErrorObject `json:"error"`
}

type WSErrorObject struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type MinerStartRequest struct {
    NumThreads int `json:"threads"`
}

type MinerStopRequest struct {
    NumThreads int `json:"threads"`
}

type MinerHashrateResponse struct {
	Hashrate int64 `json:"hashrate"`
}

type ImportPresaleWalletRequest struct {
	Path string `json:"path"`
	Password string `json:"password"`
}

type ImportPresaleWalletResponse struct {
	Address string `json:"address"`
}
