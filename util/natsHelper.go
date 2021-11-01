package util

import (
	"sync"

	"github.com/nats-io/nats.go"
)

var natsOnce sync.Once
var conn *nats.EncodedConn

const CODES_CREATED_EVENT = "codescreated"
const CODES_UPDATED_EVENT = "codesupdated"
const CODES_DELETE_EVENT = "codesdeleted"

func GetNatsClient() (*nats.EncodedConn, error) {
	//Perform connection creation operation only once.
	config := LoadConfig()
	natsOnce.Do(func() {
		// Set client options
		nc, _ := nats.Connect(config.NatsServer)
		conn, _ = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	})
	return conn, nil
}
