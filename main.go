package main

import (
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/papertrail/go-tail/follower"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Starting loghub-go-tailer")
	verbosePtr := flag.Bool("v", false, "Enables verbose printing")
	brokersPtr := flag.String("brokers", "localhost:9092", "List of comma separated kafka brokers")
	topicPtr := flag.String("topic", "test", "Topic to write tailed file to")
	tailPtr := flag.String("tail", "EXAMPLE_INPUT.log", "File to tail")
	maxRetryPtr := flag.Int("maxRetry", 5, "Kafka max retry limit")
	flag.Parse()

	verbose := *verbosePtr
	brokers := *brokersPtr
	topic := *topicPtr
	tail := *tailPtr
	maxRetry := *maxRetryPtr

	if verbose {
		fmt.Println("Verbose printing enabled")
		fmt.Println("Brokers:", brokers)
		fmt.Println("Topic:", topic)
		fmt.Println("Tail file:", tail)
		fmt.Println("Max Retry:", maxRetry)
	}

	brokerArr := strings.Split(brokers, ",")

	if verbose {
		fmt.Println("Broker list size: ", len(brokerArr))
	}

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = maxRetry
	config.Version = sarama.V1_1_0_0
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokerArr, config)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	t, err := follower.New(tail, follower.Config{
		Whence: io.SeekEnd,
		Offset: 0,
		Reopen: true,
	})

	for line := range t.Lines() {
		strLine := line.String()
		if verbose {
			fmt.Println("Read:", strLine)
		}
		msg := &sarama.ProducerMessage{
			Topic:     topic,
			Value:     sarama.StringEncoder(strLine),
			Timestamp: time.Now(),
		}

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			panic(err)
		}
		if verbose {
			fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
		}
	}

	if err != nil {
		fmt.Println(os.Stderr, err)
	}
}
