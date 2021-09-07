package kafka

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

const (
	kafkaContextKey = "kafka"

	// todo move to config
	brokerAddress = "127.0.0.1:9094"
)

func RegisterIn(ctx context.Context) context.Context {
	producer, err := buildProducer(brokerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to build producer")
	}
	return context.WithValue(ctx, kafkaContextKey, producer)
}

func getProducer(ctx context.Context) sarama.SyncProducer {
	producer, ok := ctx.Value(kafkaContextKey).(sarama.SyncProducer)
	if !ok {
		log.Fatal().Msg("failed to get producer")
	}
	return producer
}

func SendMessageWithContext(ctx context.Context, topic string, message interface{}) error {
	producer := getProducer(ctx)
	msg, err := json.Marshal(message)
	if err != nil {
		log.Error().Err(err).Msg("failed to prepare message")
		return err
	}

	kafkaMsg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(msg),
	}
	sendMessage, i, err := producer.SendMessage(kafkaMsg)
	if err != nil {
		log.Error().Err(err).Msg("failed to send message")
		return err
	}
	log.Debug().Msgf("message send: %v %v", sendMessage, i)
	return nil
}

func buildProducer(brokerAddress string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer([]string{brokerAddress}, config)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create producer")
		return nil, err
	}
	return producer, nil
}
