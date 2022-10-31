package handler

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/pradeepneosoft/websocket-poc/models"
)

const (
	subscribe   = "subscribe"
	unsubscribe = "unsubscribe"
)

var Subscriptions = map[string][]models.Client{}

func ProcessMessage(client models.Client, payload []byte) {
	m := models.Message{}
	if err := json.Unmarshal(payload, &m); err != nil {
		Send(&client, "Server: Invalid payload")
	}

	switch m.Action {

	case subscribe:
		Subscribe(&client, m.Topic)

	case unsubscribe:
		Unsubscribe(&client, m.Topic)

	default:
		Send(&client, "Server: Action unrecognized")
	}

}

func Publish(topic string, message []byte) {
	//
	for _, client := range Subscriptions[topic] {
		Send(&client, string(message))
	}

}

func Subscribe(client *models.Client, topic string) {

	sort.Slice(Subscriptions[topic], func(i, j int) bool {
		return Subscriptions[topic][i].ID <= Subscriptions[topic][j].ID
	})

	idx := sort.Search(len(Subscriptions[topic]), func(i int) bool {
		return string(Subscriptions[topic][i].ID) >= client.ID
	})

	// if idx < len(Subscriptions[topic]) && Subscriptions[topic][idx].ID == client.ID {
	// 	fmt.Println("Found : ", idx)

	// } else {
	// 	fmt.Println("not Found : ", idx)
	// 	Subscriptions[topic] = append(Subscriptions[topic], *client)

	// }
	if !(idx < len(Subscriptions[topic]) && Subscriptions[topic][idx].ID == client.ID) {
		fmt.Println("not Found : ", idx)
		Subscriptions[topic] = append(Subscriptions[topic], *client)
	}
	//  else {

	// 	fmt.Println("Found : ", idx)

	// }

	fmt.Println(Subscriptions)
}

func Unsubscribe(client *models.Client, topic string) {

	for i, val := range Subscriptions[topic] {
		if val.ID == client.ID {
			if i == len(Subscriptions[topic]) {
				Subscriptions[topic] = Subscriptions[topic][:len(Subscriptions)-1]

			} else {
				Subscriptions[topic] = append((Subscriptions[topic])[:i], (Subscriptions[topic])[i+1:]...)
				break
			}
		}
	}
	fmt.Println(Subscriptions)
}

func RemoveClient(client models.Client) {

	for key, sub := range Subscriptions {
		for i := 0; i < len(sub); i++ {
			if sub[i].ID == client.ID {
				if i == len(Subscriptions[key])-1 {
					Subscriptions[key] = Subscriptions[key][:len(Subscriptions[key])-1]
				} else {
					Subscriptions[key] = append(Subscriptions[key][:i], Subscriptions[key][i+1:]...)
					break
				}
			}
		}
	}
	fmt.Println(Subscriptions)

}
func Send(client *models.Client, message string) {
	client.Connection.WriteMessage(1, []byte(message))
}
