package model

type WebSocketV2AuthenticationRequest struct {
	Op               string `json:"op"`
	AuthType         string `json:"type"`
	AccessKey        string `json:"AccessKeyId"`
	SignatureMethod  string `json:"SignatureMethod"`
	SignatureVersion string `json:"SignatureVersion"`
	Timestamp        string `json:"Timestamp"`
	Signature        string `json:"Signature"`
}

func (p *WebSocketV2AuthenticationRequest) Init() *WebSocketV2AuthenticationRequest {

	p.Op = "auth"
	p.AuthType = "api"
	p.SignatureMethod = "HmacSHA256"
	p.SignatureVersion = "2"

	return p
}
