package client

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws/response/market"
)

// Responsible to handle Trade data from WebSocket
type KLineWebSocketClient struct {
	ws.WebSocketClientBase
}

// Initializer
func (p *KLineWebSocketClient) Init(host string) *KLineWebSocketClient {
	p.WebSocketClientBase.Init(host, "/linear-swap-ws")
	return p
}

// Set callback handler
func (p *KLineWebSocketClient) SetHandler(
	connectedHandler ws.ConnectedHandler,
	responseHandler ws.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Request latest 300 trade data
func (p *KLineWebSocketClient) Request(symbol string, clientId string, period string, from int64, to int64) {
	topic := fmt.Sprintf("market.%s.kline.%s", symbol, period)
	req := fmt.Sprintf("{\"req\": \"%s\", \"from\":%d, \"to\":%d, \"id\": \"%s\" }", topic, from, to, clientId)

	p.Send(req)

	applogger.Info("WebSocket requested, topic=%s, clientId=%s", topic, clientId)
}

// Subscribe latest completed trade in tick by tick mode
func (p *KLineWebSocketClient) Subscribe(symbol string, period string, clientId string) {
	topic := fmt.Sprintf("market.%s.kline.%s", symbol, period)
	sub := fmt.Sprintf("{\"sub\": \"%s\", \"id\": \"%s\"}", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe trade
func (p *KLineWebSocketClient) UnSubscribe(symbol string, period string, clientId string) {
	topic := fmt.Sprintf("market.%s.kline.%s", symbol, period)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\", \"id\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *KLineWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubKLineResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
