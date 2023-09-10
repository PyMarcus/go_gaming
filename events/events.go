package events

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var response string

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	response = string(msg.Payload())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
}

func PubAndRecv(playerName string, posx, posy float64) string{

	var broker = "broker.emqx.io"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(playerName)
	opts.SetUsername("emqx")
	opts.SetPassword("public")
	opts.SetDefaultPublishHandler(messagePubHandler)
	//opts.SetPingTimeout(60 * time.Second)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	
	message := fmt.Sprintf("%s:%.1f:%.1f", playerName, posx, posy)

	sub(client)
	publish(client, message)

	//client.Disconnect(250)
	
	return response
}

func publish(client mqtt.Client, message string) {
	token := client.Publish("topic/test", 0, false, message)
	token.Wait()

}

func sub(client mqtt.Client) {
	topic := "topic/test"
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
}
