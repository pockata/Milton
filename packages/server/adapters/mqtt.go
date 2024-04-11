package adapters

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

type MQTTConfig struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type MQTT struct {
	client mqtt.Client
}

func (m MQTT) Setup(Config MQTTConfig) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(
		fmt.Sprintf("tcp://%s:%s", Config.Server, Config.Port),
	)
	opts.SetClientID("MiltonServer")
	opts.SetAutoReconnect(true)
	opts.SetUsername(Config.User)
	opts.SetPassword(Config.Password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = func(client mqtt.Client) {
		fmt.Println("Connected")
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		fmt.Printf("Connection lost: %v", err)
	}

	m.client = mqtt.NewClient(opts)
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		// TODO: Handle error more gracefully
		panic(token.Error())
	}

	topic := "milton/units"
	m.client.Subscribe(topic, 2, nil)
}

func (m MQTT) SendCommand(topic string, cmd string) {
	token := m.client.Publish(topic, 2, false, cmd)
	isReceived := token.WaitTimeout(time.Second * 3)

	if isReceived {
		log.Printf("Sent cmd to `%s` with contents: `%s`", topic, cmd)
	} else {
		log.Printf("[ERROR] Failed to send cmd to %s with contents: %s", topic, cmd)
	}
}

func (m MQTT) SendCommandToUnit(unit string, cmd string) {
	m.SendCommand(fmt.Sprintf("milton/unit/%s", unit), cmd)
}
