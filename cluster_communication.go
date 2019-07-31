package ucluster

import (
	"github.com/uplus-io/ucluster/model"
	"github.com/uplus-io/ugo/proto"
)

type clusterCommunicationImplementor struct {
	cluster *Cluster
}

func newClusterCommunicationImplementor(cluster *Cluster) *clusterCommunicationImplementor {
	return &clusterCommunicationImplementor{cluster: cluster}
}

func (p *clusterCommunicationImplementor) SendNodeInfoTo(to int32) error {
	transport := p.cluster.transport

	nodeInfo := p.cluster.delegate.LocalNodeStorageInfo()
	nodeInfoData, err := proto.Marshal(nodeInfo)
	if err != nil {
		return err
	}
	clusterStat := model.NewTCPPacket(model.PacketType_SystemHi, int32(transport.Me().Id), to, nodeInfoData)
	p.cluster.SendAsyncPacket(clusterStat)
	return nil
}
