package pubsub

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/websocket/v2"
)

var (
	PUBLISH     = "publish"
	SUBSCRIBE   = "subscribe"
	UNSUBSCRIBE = "unsubscribe"
)

type Client struct {
	Id         string
	Connection *websocket.Conn
	// interested topics subscriptions
}

type PubSub struct {
	Clients       []Client
	Subscriptions []Subscription
}

type Message struct {
	Action  string          `json:"action"`
	Topic   string          `json:"topic"`
	Message json.RawMessage `json:"message"`
}

type Subscription struct {
	Topic  string
	Client *Client
}

func (ps *PubSub) AddClient(client *Client) *PubSub {
	if ps.ClientExist(client.Id) {
		// if client is already in clients list do notting
		return ps
	}
	ps.Clients = append(ps.Clients, *client)
	//fmt.Println("adding new client to the list", client.Id)

	payload := []byte("hello client Id:" + client.Id) // not important to send the id back
	client.Connection.WriteMessage(1, payload)

	return ps
}

func (ps *PubSub) RemoveClient(client *Client) *PubSub {
	for i, sub := range ps.Subscriptions {
		if sub.Client.Id == client.Id {
			ps.Subscriptions = append(ps.Subscriptions[:i], ps.Subscriptions[i+1:]...)
		}
	}

	for i, c := range ps.Clients {
		if c.Id == client.Id {
			ps.Clients = append(ps.Clients[:i], ps.Clients[i+1:]...)
		}
	}

	return ps
}

func (ps *PubSub) ClientExist(id string) bool {
	for _, client := range ps.Clients {
		if client.Id == id {
			return true
		}
	}
	return false
}

func (ps *PubSub) GetSubscriptions(topic string, client *Client) []Subscription {
	var subscriptionList []Subscription

	for _, subscription := range ps.Subscriptions {

		if client != nil {

			if subscription.Client.Id == client.Id && subscription.Topic == topic {
				subscriptionList = append(subscriptionList, subscription)
			}

		} else {
			if subscription.Topic == topic {
				subscriptionList = append(subscriptionList, subscription)
			}
		}

	}

	return subscriptionList
}

func (ps *PubSub) Subscribe(client *Client, topic string) *PubSub {

	clientSubscriptions := ps.GetSubscriptions(topic, client)

	if len(clientSubscriptions) > 0 {
		// client is already subscribed to this topic
		return ps
	}

	newSubscription := Subscription{
		Topic:  topic,
		Client: client,
	}

	ps.Subscriptions = append(ps.Subscriptions, newSubscription)

	return ps
}

func (ps *PubSub) Publish(topic string, message []byte) {
	subscriptions := ps.GetSubscriptions(topic, nil)

	for _, sub := range subscriptions {
		//sub.Client.Connection.WriteMessage(1, message)
		err := sub.Client.Send(message)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (c *Client) Send(message []byte) error {
	return c.Connection.WriteMessage(1, message)
}

func (ps *PubSub) UnSubscribe(client *Client, topic string) *PubSub {
	subLen := len(ps.Subscriptions)
	for i, s := range ps.Subscriptions {
		if s.Topic == topic && s.Client.Id == client.Id {
			ps.Subscriptions[i] = ps.Subscriptions[subLen-1]
			ps.Subscriptions = ps.Subscriptions[:subLen-1]
		}
	}

	return ps
}

func (ps *PubSub) HandleReceiveMessage(client Client, messageType int, payload []byte) *PubSub {
	m := Message{}
	err := json.Unmarshal(payload, &m)
	if err != nil {
		fmt.Println("This is not correct message format")
		return ps
	}

	//if err := client.Connection.WriteMessage(messageType, message); err != nil {
	//	log.Println(err)
	//	return nil
	//}
	//fmt.Println("New message from Client: ", message, messageType)

	switch m.Action {
	case PUBLISH:
		//fmt.Println("THis is publish new message")
		//ps.Publish(m.Topic, m.Message)
		break
	case SUBSCRIBE:
		ps.Subscribe(&client, m.Topic)
		//fmt.Println("new subscribe to topic", m.Topic, len(ps.Subscriptions), client.Id)
		break
	default:
		break
	case UNSUBSCRIBE:
		ps.UnSubscribe(&client, m.Topic)
		//fmt.Println("Unsubscribed to topic", m.Topic, len(ps.Subscriptions), client.Id)
		break
	}

	return ps
}
