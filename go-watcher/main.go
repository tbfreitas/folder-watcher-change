package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	url = "/Users/tarcisio/Desktop/teste.txt"
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

	var broker = "tcp://localhost:1883"
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

	topic := "topic/secret"
	token = client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s\n", topic)

	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("%d", i)
		token = client.Publish(topic, 0, false, text)
		token.Wait()
		time.Sleep(time.Second)
	}

	client.Disconnect(100)
	// watcher, err := fsnotify.NewWatcher()
	// p, errCon := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})

	// if err != nil {
	// 	panic(err)
	// }
	// if errCon != nil {
	// 	log.Fatal(err)
	// }

	// defer watcher.Close()
	// defer p.Close()

	// done := make(chan bool)

	// // anonymous function
	// go func() {
	// 	for {
	// 		select {
	// 		case event, ok := <-watcher.Events:
	// 			if !ok {
	// 				return
	// 			}
	// 			log.Println("event:", event)

	// 			fmt.Printf("OK")

	// 			go func() {
	// 				for e := range p.Events() {
	// 					switch ev := e.(type) {
	// 					case *kafka.Message:
	// 						if ev.TopicPartition.Error != nil {
	// 							fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
	// 						} else {
	// 							fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
	// 						}
	// 					}
	// 				}
	// 			}()

	// 			topic := "t1"
	// 			str := fmt.Sprint(event.Op)
	// 			for _, word := range []string{event.Name, str} {
	// 				p.Produce(&kafka.Message{
	// 					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 					Value:          []byte(word),
	// 				}, nil)
	// 			}
	// 			log.Println("modified file:", event.Name)

	// 		case err, ok := <-watcher.Errors:
	// 			if !ok {
	// 				return
	// 			}
	// 			log.Println("error:", err)
	// 		}
	// 	}
	// }()

	// err = watcher.Add(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// <-done

}
