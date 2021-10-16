package client

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws/response/market"
)

// Responsible to handle Trade data from WebSocket
type BboWebSocketClient struct {
	ws.WebSocketClientBase
}

// Initializer
func (p *BboWebSocketClient) Init(host string) *BboWebSocketClient {
	p.WebSocketClientBase.Init(host, "/linear-swap-ws")
	return p
}

// Set callback handler
func (p *BboWebSocketClient) SetHandler(
	connectedHandler ws.ConnectedHandler,
	responseHandler ws.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Subscribe latest completed trade in tick by tick mode
func (p *BboWebSocketClient) Subscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.bbo", symbol)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe trade
func (p *BboWebSocketClient) UnSubscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("market.%s.bbo", symbol)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *BboWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubBBOResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
