package message_broker

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/thoas/go-funk"
)

// Config ...
type Config struct {
	// Connect url
	URL string

	// Auth user
	User string

	// Auth password
	Password string

	// TLS config
	TLS *TLSConfig

	// RequestTimeout
	RequestTimeout time.Duration

	Debug bool
}

// TLSConfig ...
type TLSConfig struct {
	// Cert file
	CertFilePath string

	// Key file
	KeyFilePath string

	// Root CA
	RootCAFilePath string
}

var (
	natsServer    *nats.Conn
	configuration Config
)

// Connect ...
func Connect(cfg Config) error {
	if cfg.URL == "" {
		return errors.New("natsio: connect URL is required")
	}

	// Connect options
	opts := []nats.Option{
		nats.ReconnectWait(2 * time.Second), // Time to wait before attempting reconnection
		nats.MaxReconnects(-1),              // Unlimited reconnections
	}

	// Has authentication
	if cfg.User != "" {
		opts = append(opts, nats.UserInfo(cfg.User, cfg.Password))
	}

	// If it has TLS
	if cfg.TLS != nil {
		opts = append(opts, nats.ClientCert(cfg.TLS.CertFilePath, cfg.TLS.KeyFilePath))
		opts = append(opts, nats.RootCAs(cfg.TLS.RootCAFilePath))
	}

	nc, err := nats.Connect(cfg.URL, opts...)
	if err != nil {
		msg := fmt.Sprintf("natsio: error when connecting to NATS: %s", err.Error())
		return errors.New(msg)
	}

	fmt.Printf("⚡️[natsio]: connected to %s \n", cfg.URL)

	if cfg.RequestTimeout == 0 {
		cfg.RequestTimeout = requestTimeout
	}

	// Set client
	natsServer = nc
	configuration = cfg

	return nil
}

// GetConn ...
func GetConn() *nats.Conn {
	return natsServer
}

// GetConfig ...
func GetConfig() Config {
	return configuration
}

// mergeAndUniqueArrayStrings ...
func mergeAndUniqueArrayStrings(arr1, arr2 []string) []string {
	var result = make([]string, 0)
	result = append(result, arr1...)
	result = append(result, arr2...)
	result = funk.UniqString(result)
	return result
}

// generateSubjectNames ...
func generateSubjectNames(streamName string, subjects []string) []string {
	var result = make([]string, 0)
	for _, subject := range subjects {
		name := combineStreamAndSubjectName(streamName, subject)
		result = append(result, name)
	}
	return result
}

func combineStreamAndSubjectName(stream, subject string) string {
	return fmt.Sprintf("%s.%s", stream, subject)
}

func generateStreamConfig(stream string, subjects []string) *nats.StreamConfig {
	cfg := nats.StreamConfig{
		Name:         stream,
		Subjects:     subjects,
		Retention:    nats.WorkQueuePolicy,
		MaxConsumers: -1,
		MaxMsgSize:   -1,
		MaxMsgs:      -1,
		NoAck:        false,
	}
	return &cfg
}

// ClientRequest ...
func ClientRequest[REQUEST any, RESPONSE any](subject string, req REQUEST) (RESPONSE, error) {
	var (
		traceId = uuid.New().String()
	)
	if GetConfig().Debug {
		fmt.Println(fmt.Sprintf("[%s] [REQUEST] [%s] with data %v", traceId, subject, req))
	}
	var (
		res     RESPONSE
		natsRes NatsResponse
	)
	msg, err := Request(subject, ToBytes(req))
	if err != nil {
		return res, err
	}
	if err = json.Unmarshal(msg.Data, &natsRes); err != nil {
		return res, err
	}

	if err = json.Unmarshal(natsRes.Data, &res); err != nil {
		return res, err
	}
	if GetConfig().Debug {
		fmt.Println(fmt.Sprintf("[%s] [RESPONSE] [%s] with data %v", traceId, subject, string(natsRes.Data)))
	}
	return res, nil
}
