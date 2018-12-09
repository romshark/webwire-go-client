package client

import (
	"context"

	"github.com/qbeon/webwire-go"
	"github.com/qbeon/webwire-go/message"
	"github.com/qbeon/webwire-go/payload"
)

func (clt *client) sendNamelessRequest(
	ctx context.Context,
	messageType byte,
	payload payload.Payload,
) (webwire.Reply, error) {
	request := clt.requestManager.Create()

	writer, err := clt.conn.GetWriter()
	if err != nil {
		return nil, err
	}

	if err := message.WriteMsgNamelessRequest(
		writer,
		messageType,
		request.IdentifierBytes,
		payload.Data,
	); err != nil {
		return nil, err
	}

	clt.heartbeat.reset()

	// Block until request either times out or a response is received
	return request.AwaitReply(ctx)
}
