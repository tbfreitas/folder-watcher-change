package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fsnotify/fsnotify"
)

const (
	url           = "/tmp/files"
	topic         = "topic/secret"
	addressBroker = "tcp://mosquitto:1883"
)

type value struct {
	File   string `json:"file"`
	Action string `json:"action"`
}

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Message %s received on topic %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection Lost: %s\n", err.Error())
}

func main() {

	var broker = addressBroker
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID("go_mqtt_example")
	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		panic(err)
	}

	defer watcher.Close()
	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Op == fsnotify.Create || event.Op == fsnotify.Rename || event.Op == fsnotify.Write {
					log.Println("modified created:", event.Op)
					message := value{File: event.Name, Action: event.Op.String()}
					messageJSON, err := json.Marshal(message)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}
					fmt.Println(messageJSON)
					client.Publish(topic, 1, true, messageJSON)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(url)
	if err != nil {
		log.Fatal(err)
	}
	<-done

}
