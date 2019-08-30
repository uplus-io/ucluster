/*
 * Copyright (c) 2019 uplus.io
 */

package v1

import (
	"github.com/uplus-io/ugo/core"
	log "github.com/uplus-io/ugo/logger"
	"github.com/uplus-io/ugo/utils"
	"strings"
	"sync"
	"time"
)

const (
	warehouseAvailableInterval = 5 //可用时间间隔 second
)

type WarehouseListener func(status WarehouseStatus)

type WarehouseStatus uint8

const (
	WarehouseStatus_Launching WarehouseStatus = iota
	WarehouseStatus_Normal
	WarehouseStatus_Node_Changed
)

type Warehouse struct {
	cluster    *Cluster
	Centers    *core.Array
	applicants *core.Array
	delegate   ClusterDelegate

	available             bool  //可用
	lastCommunicationTime int64 //最后节点状态变化通讯时间
	clusterReadying       bool  //当前节点服务已就绪标记
	status                WarehouseStatus

	listener WarehouseListener

	sync.RWMutex
}

func NewWarehouse(cluster *Cluster) *Warehouse {
	return &Warehouse{
		cluster:    cluster,
		Centers:    core.NewArray(),
		applicants: core.NewArray(),
		status:     WarehouseStatus_Launching,
	}
}

func GenerateRepositoryId(group string) int32 {
	return utils.StringToInt32(group)
}

func (p *Warehouse) IfAbsentCreateDataCenter(group string) *DataCenter {
	id := GenerateRepositoryId(group)
	return p.Centers.IfAbsentCreate(NewDataCenter(id)).(*DataCenter)
}

func (p *Warehouse) IfPresent(ipv4 string) *DataCenter {
	return nil
}

func (p *Warehouse) GetCenter(dc int32) *DataCenter {
	p.RLock()
	defer p.RUnlock()
	return p.Centers.Id(dc).(*DataCenter)
}

func (p *Warehouse) GetNode(dc int32, nodeId int32) *Node {
	p.RLock()
	defer p.RUnlock()
	center := p.GetCenter(dc)
	if center != nil {
		return center.nodes.Id(nodeId).(*Node)
	}
	return nil
}

func (p *Warehouse) JoinNode(ip string, port int) *Node {
	p.lastCommunicationTime = time.Now().Unix()
	p.status = WarehouseStatus_Node_Changed
	node := NewNode(ip, port, 0)
	if p.clusterReadying {
		storageInfo := LocalNodeStorageInfo()
		NodeStorageInfoReply(id, Id, storageInfo)
	} else {
		p.applicants.Add(node)
		log.Infof("cluster applicant[%d:%s:%d] join", Id, Ip, Port)
	}
	return node
}

func (p *Warehouse) LeaveNode(ip string, port int) *Node {
	p.lastCommunicationTime = time.Now().Unix()

	node := NewNode(ip, port, 0)
	if p.clusterReadying {
		//todo:set node invalid
	} else {
		p.applicants.Delete(Id)
		log.Infof("cluster applicant[%d:%s:%d] leave", Id, Ip, Port)
	}
	return node
}

func (p *Warehouse) AddNode(node *Node, partitionSize int, replicaSize int) error {
	p.Lock()
	defer p.Unlock()
	bits := strings.Split(Ip, ".")
	center := p.IfAbsentCreateDataCenter(bits[0])
	area := IfAbsentCreateArea(bits[1])
	rack := IfAbsentCreateRack(bits[2])
	newNode := IfAbsentCreateNode(Ip, Port)
	DataCenter = Id
	Area = Id
	Rack = Id
	//todo:需要注意 分区数与比重与已存数据不一致问题
	Weight = Weight
	PartitionSize = partitionSize
	ReplicaSize = replicaSize

	addNode(newNode)

	node = newNode
	return nil
}

func (p *Warehouse) Group() {
	p.Lock()
	defer p.Unlock()
	for i := 0; i < p.Centers.Len(); i++ {
		center := p.Centers.Index(i).(*DataCenter)
		Group()
	}
}

func (p *Warehouse) Applicants() *core.Array {
	return p.applicants
}

func (p *Warehouse) Readying(listener WarehouseListener) {
	p.Lock()
	defer p.Unlock()
	p.clusterReadying = true
	p.listener = listener
	go p.selfCheck()
}

func (p *Warehouse) selfCheck() {
	for {
		ts := time.Now().Unix() - p.lastCommunicationTime
		log.Debugf("self check,ts:%d interval:%d", ts, warehouseAvailableInterval)
		if ts >= warehouseAvailableInterval {
			p.status = WarehouseStatus_Normal
			p.listener(p.status)
			break
			//runtime.Gosched()
		}
		time.Sleep(time.Second)
	}
}

func (p *Warehouse) print() {
	for i := 0; i < p.Centers.Len(); i++ {
		center := p.Centers.Index(i).(*DataCenter)
		log.Debugf("%d dataCenter[%d] has %d areas", i, Id, center.Areas.Len())
		for j := 0; j < center.Areas.Len(); j++ {
			area := center.Areas.Index(j).(*Area)
			log.Debugf("    %d area[%d] has %d racks", j, Id, area.Racks.Len())
			for k := 0; k < area.Racks.Len(); k++ {
				rack := area.Racks.Index(k).(*Rack)
				log.Debugf("        %d rack[%d] has %d nodes", j, Id, rack.Nodes.Len())
				for l := 0; l < rack.Nodes.Len(); l++ {
					node := rack.Nodes.Index(l).(*Node)
					log.Debugf("            %d node[id:%d dataCenter:%d area:%d rack:%d] part[size:%d] replica[size:%d]", k, Id, DataCenter, Area, Rack, PartitionSize, ReplicaSize)
					for m := 0; m < node.Partitions.Len(); m++ {
						part := node.Partitions.Index(m).(*Partition)
						log.Debugf("                %d node[%d-%d] part[id:%d index:%d dataCenter:%d area:%d rack:%d] has replicas:%d",
							m, Node, Index, Id, Index, DataCenter, Area, Rack, part.Replicas.Len())
						for n := 0; n < part.Replicas.Len(); n++ {
							replica := part.Replicas.Index(n).(*Partition)
							log.Debugf("                    %d node[%d-%d] replica[id:%d index:%d dataCenter:%d area:%d rack:%d]",
								n, Node, Index, Id, Index, DataCenter, Area, Rack)
						}
					}
				}
			}
		}
	}
}
