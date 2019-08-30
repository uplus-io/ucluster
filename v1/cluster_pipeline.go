package v1

import "github.com/uplus-io/ucluster/v1/model"

type ClusterPipeline struct {
	pipeline PacketPipeline
}

func NewClusterPipeline(pipeline PacketPipeline) *ClusterPipeline {
	return &ClusterPipeline{pipeline: pipeline}
}

func (p *ClusterPipeline) SyncSend(packet *model.Packet) *model.Packet {
	channel := InSyncWrite(packet)
	return <-Read()
}
func (p *ClusterPipeline) ASyncSend(packet *model.Packet) {
	InWrite(packet)
}
