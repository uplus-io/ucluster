package ucluster

import "github.com/uplus-io/ucluster/model"

type ClusterPipeline struct {
	pipeline PacketPipeline
}

func NewClusterPipeline(pipeline PacketPipeline) *ClusterPipeline {
	return &ClusterPipeline{pipeline: pipeline}
}

func (p *ClusterPipeline) SyncSend(packet *model.Packet) *model.Packet {
	channel := p.pipeline.InSyncWrite(packet)
	return <-channel.Read()
}
func (p *ClusterPipeline) ASyncSend(packet *model.Packet) {
	p.pipeline.InWrite(packet)
}
