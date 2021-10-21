package notify

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws/response/notify"
)

// Responsible to handle Trade data from WebSocket
type IsolatedOrdersClient struct {
	ws.WebSocketV2ClientBase
}

// Initializer
func (p *IsolatedOrdersClient) Init(accessKey string, secretKey string, host string) *IsolatedOrdersClient {
	p.WebSocketV2ClientBase.Init(accessKey, secretKey, host, "/linear-swap-notification")
	return p
}

// Set callback handler
func (p *IsolatedOrdersClient) SetHandler(
	authenticationResponseHandler ws.AuthenticationV2ResponseHandler,
	responseHandler ws.ResponseHandler) {
	p.WebSocketV2ClientBase.SetHandler(authenticationResponseHandler, p.handleMessage, responseHandler)
}

// Subscribe latest completed trade in tick by tick mode
func (p *IsolatedOrdersClient) Subscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("orders.%s", symbol)
	sub := fmt.Sprintf("{\"sub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe trade
func (p *IsolatedOrdersClient) UnSubscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("orders.%s", symbol)
	unsub := fmt.Sprintf("{\"unsub\": \"%s\",\"id\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *IsolatedOrdersClient) handleMessage(msg string) (interface{}, error) {
	result := notify.SubOrdersResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
