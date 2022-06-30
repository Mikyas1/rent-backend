package pubsub

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"log"
)

type Ws struct {
	Ps *PubSub
}

func (w Ws) Notify(userId string, eventName string, data interface{}) error {
	var payload = map[string]interface{}{"data": data, "event": eventName}

	b, err := GetBytes(payload)
	if err != nil {
		return err
	}

	w.Ps.Publish(userId, b) // ignored error
	return nil
}

func (w Ws) BulkNotify(userIds []string, eventName string, data interface{}) error {
	var payload = map[string]interface{}{"data": data, "event": eventName}
	b, err := GetBytes(payload)
	if err != nil {
		return err
	}
	for _, userId := range userIds {
		w.Ps.Publish(userId, b) // ignored error
	}
	return nil
}

func (ws Ws) WsEndPoint() func(c *fiber.Ctx) error {
	return websocket.New(func(c *websocket.Conn) {
		client := Client{
			Id:         uuid.New().String(),
			Connection: c,
		}

		ws.Ps.AddClient(&client)
		fmt.Println("New client is connected, total: ", len(ws.Ps.Clients))

		for {
			mt, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			ws.Ps.HandleReceiveMessage(client, mt, msg)
		}
	})
}

func GetBytes(payload interface{}) ([]byte, error) {
	jsonStr, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return []byte(jsonStr), nil
}

// CreateWs creates web socket listening end point
func CreateWs() Ws {
	ps := PubSub{}
	return Ws{
		Ps: &ps,
	}
}
