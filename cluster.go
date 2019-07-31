/*
 * Copyright (c) 2019 uplus.io
 */

package ucluster

import (
	"github.com/uplus-io/ucluster/model"
	log "github.com/uplus-io/ugo/logger"
	"github.com/uplus-io/ugo/proto"
	"sync"
)

type Cluster struct {
	id     int32
	config UClusterConfig

	warehouse         *Warehouse
	communication     clusterCommunication
	dataCommunication clusterDataCommunication
	packetHandler     PacketHandler

	pipeline       Pipeline               //用户数据管道
	delegate       ClusterDelegate        //用户集群实现委托
	packetDelegate ClusterMessageDelegate //用户集群消息实现委托

	exit chan bool

	transport       Transport
	clusterPipeline ClusterPipeline

	clusterHealth model.ClusterHealth
	nodeHealth    model.NodeHealth
	nodeStatus    model.NodeStatus

	launched bool

	sync.RWMutex
}

func NewCluster(config UClusterConfig) *Cluster {
	//init log
	log.DebugLoggerEnable(true)
	return &Cluster{
		config:          config,
		clusterPipeline: NewClusterPacketPipeline(),
		exit:            make(chan bool),
		clusterHealth:   model.ClusterHealth_CH_Unknown,
		nodeHealth:      model.NodeHealth_Suspect,
		nodeStatus:      model.NodeStatus_Unknown,
	}
}

func (p *Cluster) Listen() {
	p.communication = newClusterCommunicationImplementor(p)
	p.dataCommunication = newClusterDataCommunicationImplementor(p)
	go p.launchGossip()
	go p.packetInLoop()
	go p.packetOutLoop()
	p.nodeHealth = model.NodeHealth_Alive
	<-p.exit
}

func (p *Cluster) getClusterHealth() model.ClusterHealth {
	return p.clusterHealth
}

func (p *Cluster) getLocalNodeHealth() model.NodeHealth {
	return p.nodeHealth
}

func (p *Cluster) getLocalNodeStatus() model.NodeStatus {
	return p.nodeStatus
}

func (p *Cluster) launchGossip() {
	p.Lock()
	defer p.Unlock()
	p.warehouse = NewWarehouse(p.communication)
	transportConfig := &TransportConfig{
		Seeds:          p.config.Seeds,
		Secret:         p.config.Secret,
		BindIp:         p.config.BindIp,
		BindPort:       p.config.BindPort,
		AdvertisePort:  p.config.AdvertisePort,
		EventListener:  NewClusterEventListener(p.warehouse),
		PacketListener: NewClusterPacketListener(p.clusterPipeline)}

	p.transport = NewTransportGossip(transportConfig)

	//server launch
	transportInfo := p.transport.Serving()
	p.id = transportInfo.Id

	p.launched = true
	log.Debugf("cluster node[%d] started %v", p.id, p.launched)

	localInfo := p.delegate.LocalNodeStorageInfo()
	p.JoinNode(p.id, int(localInfo.PartitionSize), int(localInfo.ReplicaSize))
	p.checkWarehouse()
}

func (p *Cluster) checkWarehouse() {
	p.contactCluster()
	p.warehouse.Readying(func(status WarehouseStatus) {
		transportInfo := p.transport.Me()
		repository := model.ParseRepository(transportInfo.Addr.String())
		storageInfo := p.delegate.LocalNodeStorageInfo()
		parts := storageInfo.Partitions
		for _, part := range parts {
			center := p.warehouse.GetCenter(repository.DataCenter)
			next := center.NextOfRing(uint32(part.Id))
			request := &model.DataMigrateRequest{StartRing: part.Id, EndRing: int32(next)}

			for _, node := range center.Nodes() {
				if node.Id != transportInfo.Id {
					p.dataCommunication.MigrateRequest(node.Id, request)
				}
			}
		}
	})
}

func (p *Cluster) contactCluster() {
	applicants := p.warehouse.Applicants()
	for i := 0; i < applicants.Len(); i++ {
		node := applicants.Index(i).(*Node)
		if p.id != node.Id {
			err := p.communication.SendNodeInfoTo(node.Id)
			if err != nil {
				log.Errorf("contact cluster[%d->%d] error", p.id, node.Id)
			}
		}
	}

}

func (p *Cluster) JoinNode(nodeId int32, partitionSize int, replicaSize int) {
	node := p.transport.Node(nodeId)

	p.warehouse.AddNode(NewNode(node.Addr.String(), int(node.Port), 1), partitionSize, replicaSize)
	p.warehouse.Group()

}

func (p *Cluster) SendAsyncPacket(packet *model.Packet) {
	p.clusterPipeline.InWrite(packet)
}

func (p *Cluster) SendSyncPacket(packet *model.Packet) *model.Packet {
	channel := p.clusterPipeline.InSyncWrite(packet)
	return <-channel.Read()
}

func (p *Cluster) packetInLoop() {
	for {
		packet := <-p.clusterPipeline.InRead()
		log.Debugf("send packet[%s]", packet.String())
		bytes, err := proto.Marshal(packet)
		if err != nil {
			log.Errorf("waiting to send packet marshal error[%s]", packet.String())
			continue
		}
		if packet.Mode == model.PacketMode_TCP {
			err := p.transport.SendToTCP(packet.To, bytes)
			if err != nil {
				log.Errorf("sending tcp packet error[%s]", packet.String())
				continue
			}
		} else if packet.Mode == model.PacketMode_UDP {
			err := p.transport.SendToUDP(packet.To, bytes)
			if err != nil {
				log.Errorf("sending udp packet error[%s]", packet.String())
				continue
			}
		}
		//node := p.warehouse.GetNode(packet.GetDataCenter(), packet.GetTo())
		//if node != nil {
		//}
	}
}

func (p *Cluster) packetOutLoop() {
	for {
		packet := <-p.clusterPipeline.OutRead()
		log.Debugf("received packet[%s]", packet.String())
		go func() {
			var err error
			switch packet.Type {
			case model.PacketType_System:
				message := model.SystemMessage{}
				proto.Unmarshal(packet.Content, &message)
				err = p.packetDelegate.System(message)
			case model.PacketType_Event:
				message := model.EventMessage{}
				proto.Unmarshal(packet.Content, &message)
				err = p.packetDelegate.Event(message)
			case model.PacketType_Topic:
				message := model.TopicMessage{}
				proto.Unmarshal(packet.Content, &message)
				err = p.packetDelegate.Topic(message)
			case model.PacketType_Data:
				message := model.DataMessage{}
				proto.Unmarshal(packet.Content, &message)
				err = p.packetDelegate.Data(message)
			}
			if err != nil {
				p.clusterPipeline.OutWrite(packet)
			}
		}()
	}
}
