package mqtt

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func ConnectMqtt() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(os.Getenv("MQTT_URL"))
	// opts.SetClientID("go_mqtt_client")
	opts.SetUsername(os.Getenv("MQTT_USER"))
	opts.SetPassword(os.Getenv("MQTT_PASS"))
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return client
}

func Sub(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
}

func Publish(client mqtt.Client, topic string, qos int, retained bool, body []byte) {
	token := client.Publish(topic, byte(qos), retained, body)
	token.Wait()
	time.Sleep(time.Second)
}

// func Mqtt() {
// 	dbHost := os.Getenv("DB_HOST")
// 	dbPort := os.Getenv("DB_PORT")
// 	apiKey := os.Getenv("API_KEY")

// 	fmt.Printf("Database Host: %s\n", dbHost)
// 	fmt.Printf("Database Port: %s\n", dbPort)
// 	fmt.Printf("API Key: %s\n", apiKey)
// }
