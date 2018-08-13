package consumer

import (
	"awesomeProject/judgerServer/judger"
	"awesomeProject/judgerServer/problem"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Task judge task or SPJ or upload file
type Task struct {
	TaskID int
	Data   string
}

// Consumer a consumer
func Consumer() {

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    "139.199.31.230:9092",
		"group.id":             "test",
		"session.timeout.ms":   6000,
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"}})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)
	topic := []string{"test"}
	err = c.SubscribeTopics(topic, nil)

	run := true
	var task Task
	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				err := json.Unmarshal(e.Value, &task)
				if err != nil {
					log.Println(err.Error())
				}
				if task.TaskID == 1 {
					go judger.Handler(task.Data)
				} else if task.TaskID == 2 {
					go problem.AddInputOutput(task.Data)
				} else if task.TaskID == 3 {
					// TODO SPJ
				}
				if e.Headers != nil {
					fmt.Printf("%% Headers: %v\n", e.Headers)
				}
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
				run = false
			default:
				fmt.Printf("Ignored %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()
}
