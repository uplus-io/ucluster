package ucluster

import (
	"github.com/uplus-io/ucluster/model"
	"github.com/uplus-io/ugo/proto"
)

type ClusterPacketHandler struct {
	cluster *Cluster

	handlerMap map[model.PacketType]PacketHandler
}

func NewClusterPacketHandler(cluster *Cluster) *ClusterPacketHandler {
	packetHandler := &ClusterPacketHandler{cluster: cluster, handlerMap: make(map[model.PacketType]PacketHandler)}
	packetHandler.registerDefaultHandler()
	return packetHandler
}

func (p *ClusterPacketHandler) Register(packetType model.PacketType, handler PacketHandler) error {
	_, ok := p.handlerMap[packetType]
	if ok {
		return ErrMessageHandlerExist
	}
	p.handlerMap[packetType] = handler
	return nil
}

func (p *ClusterPacketHandler) registerDefaultHandler() {
	p.Register(model.PacketType_System, p.systemHandle)
	p.Register(model.PacketType_Event, p.eventHandle)
	p.Register(model.PacketType_Topic, p.topicHandle)
	p.Register(model.PacketType_Data, p.dataHandle)
}

func (p *ClusterPacketHandler) systemHandle(packet model.Packet) error {
	message := &model.SystemMessage{}
	proto.Unmarshal(packet.Content, message)
	switch message.Type {
	case model.SystemMessageType_NodeInfo:
		return p.systemNodeInfoHandle(packet)
	case model.SystemMessageType_NodeInfoReply:
		return p.systemNodeInfoReplyHandle(packet)
	case model.SystemMessageType_DataMigrate:
		return p.systemDataMigrateHandle(packet)
	case model.SystemMessageType_DataMigrateReply:
		return p.systemDataMigrateReplyHandle(packet)
	case model.SystemMessageType_DataPush:
		return p.systemDataPushHandle(packet)
	case model.SystemMessageType_DataPushReply:
		return p.systemDataPushReplyHandle(packet)
	case model.SystemMessageType_DataPull:
		return p.systemDataPullHandle(packet)
	case model.SystemMessageType_DataPullReply:
		return p.systemDataPullReplyHandle(packet)
	default:
		return p.cluster.packetDelegate.System(p.cluster.pipeline, *message)
	}
}
func (p *ClusterPacketHandler) systemNodeInfoHandle(packet model.Packet) error {
	storageInfo := p.cluster.delegate.LocalNodeStorageInfo()
	storageInfoData, _ := proto.Marshal(storageInfo)
	message := &model.SystemMessage{Type: model.SystemMessageType_NodeInfoReply, Sender: packet.To, Content: storageInfoData}
	messageData, _ := proto.Marshal(message)
	reply := model.NewUDPPacket(model.PacketType_System, packet.To, packet.From, messageData)
	p.cluster.pipeline.ASyncSend(reply)
	return nil
}
func (p *ClusterPacketHandler) systemNodeInfoReplyHandle(packet model.Packet) error {
	message := &model.SystemMessage{}
	storageInfo := &model.NodeStorageInfo{}
	model.UnpackSystemMessage(&packet, message, storageInfo)
	p.cluster.JoinNode(packet.From, int(storageInfo.PartitionSize), int(storageInfo.ReplicaSize))
	return nil
}
func (p *ClusterPacketHandler) systemDataMigrateHandle(packet model.Packet) error {
	return nil
}
func (p *ClusterPacketHandler) systemDataMigrateReplyHandle(packet model.Packet) error {
	return nil
}
func (p *ClusterPacketHandler) systemDataPushHandle(packet model.Packet) error {
	return nil
}
func (p *ClusterPacketHandler) systemDataPushReplyHandle(packet model.Packet) error {
	return nil
}
func (p *ClusterPacketHandler) systemDataPullHandle(packet model.Packet) error {
	return nil
}
func (p *ClusterPacketHandler) systemDataPullReplyHandle(packet model.Packet) error {
	return nil
}

func (p *ClusterPacketHandler) eventHandle(packet model.Packet) error {
	return nil
}

func (p *ClusterPacketHandler) topicHandle(packet model.Packet) error {
	return nil
}

func (p *ClusterPacketHandler) dataHandle(packet model.Packet) error {
	return nil
}

func (p *ClusterPacketHandler) packetHandle(packet model.Packet) error {
	handler, exist := p.handlerMap[packet.Type]
	if exist {
		return handler(packet)
	} else {
		//todo:not found match handler
		return nil
	}
}
