package kafkaMessage

import (
	"context"
	"time"
	"github.com/segmentio/kafka-go"
)



func SendMessage(ctx context.Context, message string, kafkaData map[string]interface{}) {

	partition := 0

	conn, _ := kafka.DialLeader(ctx, "tcp", kafkaData["kafkaURL"].(string), kafkaData["topic"].(string), partition)

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(
		kafka.Message{Value: []byte(message)},
	)

	conn.Close()

}
