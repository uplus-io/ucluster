/*
 * Copyright (c) 2019 uplus.io
 */

package ucluster

import "github.com/uplus-io/ucluster/model"

type MessageDispatcher interface {
	Dispatch(packet model.Packet) error
	register(packetType model.PacketType, handler PacketHandler) error
}

type PacketHandler func(packet model.Packet) error
