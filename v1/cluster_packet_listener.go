/*
 * Copyright (c) 2019 uplus.io
 */

package v1

import (
	"github.com/uplus-io/ucluster/v1/model"
)

type PacketListener interface {
	OnReceive(packet *model.Packet)
}

type ClusterPacketListener struct {
	Pipeline PacketPipeline
}

func NewClusterPacketListener(pipeline PacketPipeline) *ClusterPacketListener {
	return &ClusterPacketListener{Pipeline: pipeline}
}

func (p *ClusterPacketListener) OnReceive(packet *model.Packet) {
	OutWrite(packet)
}
