package v1

import (
	"github.com/uplus-io/ucluster/v1/model"
	"github.com/uplus-io/ugo/goproto"
	log "github.com/uplus-io/ugo/logger"
	"github.com/uplus-io/ugo/proto"
)

type ClusterPacketDispatcher struct {
	cluster *Cluster

	handlerMap map[model.PacketType]PacketHandler
}

func NewClusterPacketDispatcher(cluster *Cluster) *ClusterPacketDispatcher {
	packetHandler := &ClusterPacketDispatcher{cluster: cluster, handlerMap: make(map[model.PacketType]PacketHandler)}
	packetHandler.registerDefaultHandler()
	return packetHandler
}

func (p *ClusterPacketDispatcher) Dispatch(packet model.Packet) error {
	handler, exist := p.handlerMap[packet.Type]
	if exist {
		return handler(packet)
	} else {
		log.Warnf("not match message handler[%v],drop packet[%v]", packet.Type, packet)
		return nil
	}
}

func (p *ClusterPacketDispatcher) Register(packetType model.PacketType, handler PacketHandler) error {
	_, ok := p.handlerMap[packetType]
	if ok {
		return ErrMessageHandlerExist
	}
	p.handlerMap[packetType] = handler
	return nil
}

func (p *ClusterPacketDispatcher) registerDefaultHandler() {
	p.Register(model.PacketType_System, p.systemHandle)
	p.Register(model.PacketType_Event, p.eventHandle)
	p.Register(model.PacketType_Topic, p.topicHandle)
	p.Register(model.PacketType_Data, p.dataHandle)
}

func (p *ClusterPacketDispatcher) systemHandle(packet model.Packet) error {
	message := &model.SystemMessage{}
	goproto.Unmarshal(packet.Content, message)
	switch message.Type {
	case model.SystemMessageType_NODE_STORAGE_INFO:
		return p.systemNodeInfoHandle(packet)
	case model.SystemMessageType_NODE_STORAGE_INFO_REPLY:
		return p.systemNodeInfoReplyHandle(packet)
	case model.SystemMessageType_DATA_MIGRATE:
		return p.systemDataMigrateHandle(packet)
	case model.SystemMessageType_DATA_MIGRATE_REPLY:
		return p.systemDataMigrateReplyHandle(packet)
	case model.SystemMessageType_DATA_PUSH:
		return p.systemDataPushHandle(packet)
	case model.SystemMessageType_DATA_PUSH_REPLY:
		return p.systemDataPushReplyHandle(packet)
	case model.SystemMessageType_DATA_PULL:
		return p.systemDataPullHandle(packet)
	case model.SystemMessageType_DATA_PULL_REPLY:
		return p.systemDataPullReplyHandle(packet)
	default:
		return System(pipeline, *message)
	}
}
func (p *ClusterPacketDispatcher) systemNodeInfoHandle(packet model.Packet) error {
	storageInfo := LocalNodeStorageInfo()
	return NodeStorageInfoReply(packet.To, packet.From, storageInfo)
}
func (p *ClusterPacketDispatcher) systemNodeInfoReplyHandle(packet model.Packet) error {
	message := &model.SystemMessage{}
	storageInfo := &model.NodeStorageInfo{}
	model.UnpackSystemMessage(&packet, message, storageInfo)
	JoinNode(packet.From, int(storageInfo.PartitionSize), int(storageInfo.ReplicaSize))
	return nil
}
func (p *ClusterPacketDispatcher) systemDataMigrateHandle(packet model.Packet) error {
	message := &model.SystemMessage{}
	request := &model.DataMigrateRequest{}
	model.UnpackSystemMessage(&packet, message, request)
	Migrate(packet.From, packet.To, request.StartRing, request.EndRing)
	return nil
}
func (p *ClusterPacketDispatcher) systemDataMigrateReplyHandle(packet model.Packet) error {
	message := &model.SystemMessage{}
	migrateResponse := &model.DataMigrateResponse{}
	model.UnpackSystemMessage(&packet, message, migrateResponse)
	log.Debugf("migrateReply[%s]", migrateResponse.String())
	return nil
}
func (p *ClusterPacketDispatcher) systemDataPushHandle(packet model.Packet) error {
	message := &model.SystemMessage{}
	pushRequest := &model.PushRequest{}
	model.UnpackSystemMessage(&packet, message, pushRequest)
	Push(packet.From, packet.To, pushRequest.Data)
	return nil
}
func (p *ClusterPacketDispatcher) systemDataPushReplyHandle(packet model.Packet) error {
	message := &model.SystemMessage{}
	pushResponse := &model.PushResponse{}
	model.UnpackSystemMessage(&packet, message, pushResponse)
	log.Debugf("pushReply[%s]", pushResponse.String())
	return nil
}
func (p *ClusterPacketDispatcher) systemDataPullHandle(packet model.Packet) error {
	return nil
}
func (p *ClusterPacketDispatcher) systemDataPullReplyHandle(packet model.Packet) error {
	return nil
}

func (p *ClusterPacketDispatcher) eventHandle(packet model.Packet) error {
	message := model.EventMessage{}
	goproto.Unmarshal(packet.Content, &message)
	return Event(pipeline, message)
}
func (p *ClusterPacketDispatcher) topicHandle(packet model.Packet) error {
	message := model.TopicMessage{}
	goproto.Unmarshal(packet.Content, &message)
	return Topic(pipeline, message)
}
func (p *ClusterPacketDispatcher) dataHandle(packet model.Packet) error {
	message := model.DataMessage{}
	goproto.Unmarshal(packet.Content, &message)
	return Data(pipeline, message)
}
