package ucluster

import (
	"github.com/uplus-io/ucluster/model"
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
