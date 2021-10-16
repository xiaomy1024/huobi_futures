package client

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws/response/market"
)

// Responsible to handle Depth data from WebSocket
type IncreDepthWebSocketClient struct {
	ws.WebSocketClientBase
}

// Initializer
func (p *IncreDepthWebSocketClient) Init(host string) *IncreDepthWebSocketClient {
	p.WebSocketClientBase.Init(host, "/linear-swap-ws")
	return p
}

// Set callback handler
func (p *IncreDepthWebSocketClient) SetHandler(
	connectedHandler ws.ConnectedHandler,
	responseHandler ws.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Subscribe latest market by price order book in snapshot mode at 1-second interval.
func (p *IncreDepthWebSocketClient) Subscribe(symbol string, size string, clientId string) {
	topic := fmt.Sprintf("market.%s.depth.size_%s.high_freq", symbol, size)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe market by price order book
func (p *IncreDepthWebSocketClient) UnSubscribe(symbol string, size string, clientId string) {
	topic := fmt.Sprintf("market.%s.depth.size_%s.high_freq", symbol, size)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *IncreDepthWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubDepthResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
