package producer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"fmt"
)

var p *kafka.Producer
var err error

var topic string

var doneChan = make(chan bool)

type ServerConf struct {
	Producer ProducerConf
}

// ConsumerConf 与Kafka有关的配置信息
type ProducerConf struct {
	Bootstrap string // KafkaAddr kafka_server地址
	Topic     string
}

func init() {

	// 导入配置文件
	// pwd, _ := os.Getwd()
	// println(pwd)
	// pos := len(pwd)
	// length := len("iov-gateway-wstdp-mediators")
	// for pwd[pos-length:pos] != "iov-gateway-wstdp-mediators" {
	// 	pos--
	// }
	// date, _ := ioutil.ReadFile(pwd[:pos] + "/conf/outputConf/producerConf.yml")
	conf := ServerConf{}
	// yaml.Unmarshal(date, &conf)
	// fmt.Println("kafkabootstrap: " + conf.Producer.Bootstrap)
	conf.Producer.Bootstrap = "139.199.31.230:9092"
	conf.Producer.Topic = "test"
	topic = conf.Producer.Topic
	p, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": conf.Producer.Bootstrap})
	// p, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "192.168.56.101"})

	if err != nil {
		fmt.Printf("Failed to create producer.yml: %s", err)
		os.Exit(1)
	}

	fmt.Printf("Created Producer %v", p)
	//todo
	go func() {
		fmt.Printf("routine start!")
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v", m.TopicPartition.Error)
				} else {
					fmt.Printf("Delivered message to topic %s [%d] at offset %v",
						*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
				}
				doneChan <- true
				continue

			default:
				fmt.Printf("Ignored event: %s", ev)
				doneChan <- false
			}
		}
		fmt.Printf("routine end!")
	}()
}

func Send(data string) bool {
	p.ProduceChannel() <- &kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, Value: []byte(data)}
	// wait for delivery report goroutine to finish
	return <-doneChan
}

func Close() {
	close(doneChan)
	p.Close()
}
