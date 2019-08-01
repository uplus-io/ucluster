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

	warehouse            *Warehouse
	communication        clusterCommunication
	localDataOperations  DataOperations
	remoteDataOperations DataOperations

	delegate        ClusterDelegate        //用户集群实现委托
	dataDelegate    ClusterDataDelegate    //用户数据实现委托
	messageDelegate ClusterMessageDelegate //用户集群消息实现委托

	exit chan bool

	packetDispatcher PacketDispatcher //数据包处理器
	transport        Transport        //协议实现
	pipeline         Pipeline         //暴露的数据通信管道
	packetPipeline   PacketPipeline   //数据包处理器管道

	clusterHealth model.ClusterHealth
	nodeHealth    model.NodeHealth
	nodeStatus    model.NodeStatus

	launched bool

	sync.RWMutex
}

func NewCluster(config UClusterConfig) *Cluster {
	//init log
	log.DebugLoggerEnable(true)
	cluster := &Cluster{
		config:         config,
		packetPipeline: NewClusterPacketPipeline(),
		exit:           make(chan bool),
		clusterHealth:  model.ClusterHealth_CH_Unknown,
		nodeHealth:     model.NodeHealth_Suspect,
		nodeStatus:     model.NodeStatus_Unknown,
	}
	cluster.packetDispatcher = NewClusterPacketDispatcher(cluster)
	cluster.pipeline = NewClusterPipeline(cluster.packetPipeline)
	cluster.communication = newClusterCommunicationImplementor(cluster.pipeline)
	cluster.warehouse = NewWarehouse(cluster)
	cluster.localDataOperations = newLocalDataOperations(cluster.communication, cluster.dataDelegate)

	//user impl
	cluster.delegate = config.Delegate
	cluster.dataDelegate = config.DataDelegate
	cluster.messageDelegate = config.MessageDelegate
	return cluster
}

func (p *Cluster) Listen() {
	go p.launchGossip()
	go p.packetInLoop()
	go p.packetOutLoop()
	p.nodeHealth = model.NodeHealth_Alive
	exitSingal := <-p.exit
	log.Infof("cluster receive quit signal[%v],ready exit", exitSingal)
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
	transportConfig := &TransportConfig{
		Seeds:          p.config.Seeds,
		Secret:         p.config.Secret,
		BindIp:         p.config.BindIp,
		BindPort:       p.config.BindPort,
		AdvertisePort:  p.config.AdvertisePort,
		EventListener:  NewClusterEventListener(p.warehouse),
		PacketListener: NewClusterPacketListener(p.packetPipeline)}

	p.transport = NewTransportGossip(transportConfig)

	//server launch
	transportInfo := p.transport.Serving()
	p.id = transportInfo.Id

	p.launched = true
	log.Debugf("cluster node[%d] started %v", p.id, p.launched)

	localInfo := p.delegate.LocalNodeStorageInfo()
	if localInfo == nil {
		log.Warnf("local node storage info is nil,ignore cluster join")
	} else {
		p.JoinNode(p.id, int(localInfo.PartitionSize), int(localInfo.ReplicaSize))
	}
	p.checkWarehouse()
}

func (p *Cluster) checkWarehouse() {
	p.contactCluster()
	p.warehouse.Readying(func(status WarehouseStatus) {
		transportInfo := p.transport.Me()
		repository := model.ParseRepository(transportInfo.Addr.String())
		storageInfo := p.delegate.LocalNodeStorageInfo()
		if storageInfo == nil {
			log.Warnf("local node storage info is nil,ignore cluster data migrate")
		} else {
			parts := storageInfo.Partitions
			for _, part := range parts {
				center := p.warehouse.GetCenter(repository.DataCenter)
				next := center.NextOfRing(uint32(part.Id))
				request := &model.DataMigrateRequest{StartRing: part.Id, EndRing: int32(next)}
				for _, node := range center.Nodes() {
					if node.Id != transportInfo.Id {
						p.communication.MigrateRequest(p.id, node.Id, request)
					}
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
			err := p.communication.NodeStorageInfoReply(p.id, node.Id, p.delegate.LocalNodeStorageInfo())
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

func (p *Cluster) packetInLoop() {
	for {
		packet := <-p.packetPipeline.InRead()
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
		packet := <-p.packetPipeline.OutRead()
		log.Debugf("received packet[%s]", packet.String())
		go func() {
			err := p.packetDispatcher.Dispatch(*packet)
			if err != nil {
				p.packetPipeline.OutWrite(packet)
			}
		}()
	}
}
