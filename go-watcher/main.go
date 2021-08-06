package main

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/fsnotify/fsnotify"
)

const (
	url           = "/Users/tarcisio/Desktop/teste.txt" // mapear volume pra arquivo local pra teste
	topic         = "topic/secret"
	addressBroker = "tcp://localhost:1883"
)

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

	// anonymous function
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				// log.Println("modified file:", event.Op)
				if event.Op == fsnotify.Create || event.Op == fsnotify.Rename || event.Op == fsnotify.Write {
					log.Println("modified created:", event.Name)
					client.Publish(topic, 1, true, event.Name)
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	// client.Publish(topic, 0, false, "nasa")
	err = watcher.Add(url)
	if err != nil {
		log.Fatal(err)
	}
	<-done

}
