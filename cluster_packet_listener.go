/*
 * Copyright (c) 2019 uplus.io
 */

package ucluster

import (
	"github.com/uplus-io/ucluster/model"
)

type PacketListener interface {
	OnReceive(packet *model.Packet)
}

type ClusterPacketListener struct {
	Pipeline ClusterPipeline
}

func NewClusterPacketListener(pipeline ClusterPipeline) *ClusterPacketListener {
	return &ClusterPacketListener{Pipeline: pipeline}
}

func (p *ClusterPacketListener) OnReceive(packet *model.Packet) {
	p.Pipeline.OutWrite(packet)
}
