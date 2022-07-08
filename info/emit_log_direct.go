package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"test_task/storage/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	id := rand.Intn(10)
	fmt.Println(id)
	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	code, body, err := bodyFrom(id)
	if err != nil {
		fmt.Println(body)
		log.Println("error getting rest response", err)
	}
	b, err := json.Marshal(body)
	if err != nil {
		log.Println("error getting rest response", err)
	}

	if code == 200 {
		err = ch.Publish(
			"logs_direct", // exchange
			"info",        // routing key
			false,         // mandatory
			false,         // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(b),
			})
		failOnError(err, "Failed to publish a message")
	} else if code == 400 {
		err = ch.Publish(
			"logs_direct", // exchange
			"error",       // routing key
			false,         // mandatory
			false,         // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(b),
			})
		failOnError(err, "Failed to publish a message")
	} else if code == 500 {
		_, body, err := bodyFrom(id)
		if err != nil {
			fmt.Println(body)
			log.Println("error getting rest response", err)
		}
		err = ch.Publish(
			"logs_direct", // exchange
			"error",       // routing key
			false,         // mandatory
			false,         // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(b),
			})
		failOnError(err, "Failed to publish a message")
	}

	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(id int) (int, *models.Response, error) {
	var response models.Response
	str := strconv.Itoa(id)
	resp, err := http.Get("http://localhost:8080/storage/phone/" + str)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	fmt.Println(resp)

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return resp.StatusCode, nil, err
	}
	return resp.StatusCode, &response, nil
}
