package notify

import (
	"encoding/json"
	"fmt"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws"
	"github.com/xiaomy1024/huobi_futures/sdk/linearswap/ws/response/notify"
)

// Responsible to handle Trade data from WebSocket
type OrdersCrossClient struct {
	ws.WebSocketV2ClientBase
}

// Initializer
func (p *OrdersCrossClient) Init(accessKey string, secretKey string, host string) *OrdersCrossClient {
	p.WebSocketV2ClientBase.Init(accessKey, secretKey, host, "/linear-swap-notification")
	return p
}

// Set callback handler
func (p *OrdersCrossClient) SetHandler(
	authenticationResponseHandler ws.AuthenticationV2ResponseHandler,
	responseHandler ws.ResponseHandler) {
	p.WebSocketV2ClientBase.SetHandler(authenticationResponseHandler, p.handleMessage, responseHandler)
}

// Subscribe latest completed trade in tick by tick mode
func (p *OrdersCrossClient) Subscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("orders_cross.%s", symbol)
	sub := fmt.Sprintf("{\"op\": \"sub\",\"topic\": \"%s\",\"cid\": \"%s\" }", topic, clientId)

	p.Send(sub)

	applogger.Info("WebSocket subscribed, topic=%s, clientId=%s", topic, clientId)
}

// Unsubscribe trade
func (p *OrdersCrossClient) UnSubscribe(symbol string, clientId string) {
	topic := fmt.Sprintf("orders_cross.%s", symbol)
	unsub := fmt.Sprintf("{\"op\": \"unsub\",\"topic\": \"%s\",\"cid\": \"%s\" }", topic, clientId)

	p.Send(unsub)

	applogger.Info("WebSocket unsubscribed, topic=%s, clientId=%s", topic, clientId)
}

func (p *OrdersCrossClient) handleMessage(msg string) (interface{}, error) {
	result := notify.SubOrdersResponse{}
	err := json.Unmarshal([]byte(msg), &result)
	return result, err
}
