/*
 * Copyright (c) 2019 uplus.io
 */

package ucluster

import (
	"github.com/uplus-io/ucluster/model"
)

type Pipeline interface {
	SyncSend(packet *model.Packet) *model.Packet
	ASyncSend(packet *model.Packet)
}

type ClusterPipeline interface {
	InSyncWrite(packet *model.Packet) *SyncChannel
	InWrite(packet *model.Packet)
	InRead() <-chan *model.Packet
	OutWrite(packet *model.Packet)
	OutRead() <-chan *model.Packet
}

type ClusterPacketPipeline struct {
	packetInCh  chan *model.Packet
	packetOutCh chan *model.Packet

	syncChannelMap map[string]*SyncChannel
}

func NewClusterPacketPipeline() *ClusterPacketPipeline {
	return &ClusterPacketPipeline{
		packetInCh:     make(chan *model.Packet),
		packetOutCh:    make(chan *model.Packet),
		syncChannelMap: make(map[string]*SyncChannel),
	}
}

func (p *ClusterPacketPipeline) InSyncWrite(packet *model.Packet) *SyncChannel {
	syncChannel := NewSyncChannel(packet.Id)
	p.syncChannelMap[packet.Id] = syncChannel
	p.InWrite(packet)
	return syncChannel
}

func (p *ClusterPacketPipeline) InWrite(packet *model.Packet) {
	p.packetInCh <- packet
}

func (p *ClusterPacketPipeline) InRead() <-chan *model.Packet {
	return p.packetInCh
}

func (p *ClusterPacketPipeline) OutWrite(packet *model.Packet) {
	syncChannel, exist := p.syncChannelMap[packet.Id]
	if exist {
		syncChannel.Write(packet)
	} else {
		p.packetOutCh <- packet
	}
}

func (p *ClusterPacketPipeline) OutRead() <-chan *model.Packet {
	return p.packetOutCh
}

type SyncChannel struct {
	PacketId string
	packetCh chan *model.Packet
}

func NewSyncChannel(packetId string) *SyncChannel {
	return &SyncChannel{PacketId: packetId, packetCh: make(chan *model.Packet)}
}
func (p *SyncChannel) Write(packet *model.Packet) {
	p.packetCh <- packet
}
func (p *SyncChannel) Read() <-chan *model.Packet {
	return p.packetCh
}
