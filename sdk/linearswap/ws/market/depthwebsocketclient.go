package client

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws/response/market"
)

// Responsible to handle Depth data from WebSocket
type DepthWebSocketClient struct {
	ws.WebSocketClientBase
}

// Initializer
func (p *DepthWebSocketClient) Init(host string) *DepthWebSocketClient {
	p.WebSocketClientBase.Init(host, "/linear-swap-ws")
	return p
}

// Set callback handler
func (p *DepthWebSocketClient) SetHandler(
	connectedHandler ws.ConnectedHandler,
	responseHandler ws.ResponseHandler) {
	p.WebSocketClientBase.SetHandler(connectedHandler, p.handleMessage, responseHandler)
}

// Subscribe latest market by price order book in snapshot mode at 1-second interval.
func (p *DepthWebSocketClient) Subscribe(symbol string, step string, clientId string) {
	topic := fmt.Sprintf("market.%s.depth.%s", symbol, step)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe market by price order book
func (p *DepthWebSocketClient) UnSubscribe(symbol string, step string, clientId string) {
	topic := fmt.Sprintf("market.%s.depth.%s", symbol, step)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *DepthWebSocketClient) handleMessage(msg string) (interface{}, error) {
	result := market.SubDepthResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
