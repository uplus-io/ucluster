/*
 * Copyright (c) 2019 uplus.io
 */

package v1

import (
	"github.com/uplus-io/ucluster/v1/model"
)

type PacketDispatcher interface {
	Dispatch(packet model.Packet) error
	Register(packetType model.PacketType, handler PacketHandler) error
}

type PacketHandler func(packet model.Packet) error
