package v1

import (
	"github.com/uplus-io/ucluster/v1/model"
)

type PacketQueue struct {
	messages map[string]*PacketMessage
}

func Put(packetId string, message PacketMessage) {
	//todo: socket sync message
}

type PacketMessage struct {
	messageCh chan *model.Packet
}
