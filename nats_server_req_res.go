package natsio

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// Default timeout 10s
const requestTimeout = 10 * time.Second

// Request ...
func Request(subject string, payload []byte) (*nats.Msg, error) {
	IsDisconnected()
	timeout := requestTimeout
	if GetConfig().RequestTimeout > 0 {
		timeout = GetConfig().RequestTimeout
	}
	msg, err := GetConn().Request(subject, payload, timeout)
	if errors.Is(err, nats.ErrNoResponders) {
		log.Printf("[NATS SERVER]: request - no responders for subject: %s", subject)
	}
	return msg, err
}

func Publish(sub string, payload []byte) error {
	IsDisconnected()
	return GetConn().Publish(sub, payload)
}

func PublishRequest(sub, reply string, data []byte) error {
	IsDisconnected()
	return GetConn().PublishRequest(sub, reply, data)
}

// Reply ...
func Reply(msg *nats.Msg, payload []byte) error {
	IsDisconnected()
	return GetConn().Publish(msg.Reply, payload)
}

// Response ...
func Response(msg *nats.Msg, payload interface{}, message string) error {
	res := NatsResponse{
		Success: false,
		Message: message,
		Data:    nil,
	}
	if message == "" {
		res.Success = true
		res.Data = ToBytes(payload)
	}
	err := Reply(msg, ToBytes(res))
	if err != nil {
		fmt.Println("[ERROR] Response : ", err.Error())
	}
	return err
}

// QueueSubscribe ...
func QueueSubscribe(subject, queue string, cb nats.MsgHandler) {
	IsDisconnected()
	_, err := GetConn().QueueSubscribe(subject, queue, cb)
	if err != nil {
		fmt.Println(fmt.Sprintf("[NATS SERVER] - queue subscribe subject %s, error: %s", subject, err.Error()))
		return
	}
	log.Println("[NATS SERVER] - queue subscribe subject: ", subject)
	return
}

func IsDisconnected() {
	if GetConn() == nil {
		panic("[NATS SERVER] Disconnected")
	}
}
