package client

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws/response/market"
)

// Responsible to handle Trade data from WebSocket
type DetailWebSocketClient struct {
	ws.WebSocketClientBase
}

// Initializer
func (p *DetailWebSocketClient) Init(host string) *DetailWebSocketClient {
	p.WebSocketClientBase.Init(host, "/linear-swap-ws")
	return p
}

// Set callback handler
func (p *DetailWebSocketClient) SetHandler(
	connectedHandler ws.ConnectedHandler,
	responseHandler ws.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Subscribe latest completed trade in tick by tick mode
func (p *DetailWebSocketClient) Subscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.detail", symbol)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe trade
func (p *DetailWebSocketClient) UnSubscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.detail", symbol)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *DetailWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubKLineResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
