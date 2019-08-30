package v1

import (
	"github.com/uplus-io/ucluster/v1/model"
	log "github.com/uplus-io/ugo/logger"
)

//集群通讯
type clusterCommunication interface {
	NodeStorageInfo(from, to int32) error
	NodeStorageInfoReply(from, to int32, request *model.NodeStorageInfo) error
	Push(from, to int32, request *model.PushRequest) *model.PushResponse
	PushReply(from, to int32, request *model.PushResponse) error
	Pull(from, to int32, request *model.PullRequest) *model.PullResponse
	PullReply(from, to int32, request *model.PullResponse) error
	MigrateRequest(from, to int32, request *model.DataMigrateRequest) error
	MigrateResponse(from, to int32, request *model.DataMigrateResponse) error
}

type clusterCommunicationImplementor struct {
	pipeline Pipeline
}

func newClusterCommunicationImplementor(pipeline Pipeline) *clusterCommunicationImplementor {
	return &clusterCommunicationImplementor{pipeline: pipeline}
}

func (p *clusterCommunicationImplementor) Push(from, to int32, request *model.PushRequest) *model.PushResponse {
	packet := model.PackTCPSystemMessage(from, to, model.SystemMessageType_DATA_PUSH, request)

	responseData := SyncSend(packet)
	systemMessage := &model.SystemMessage{}
	pushResponse := &model.PushResponse{}
	err := model.UnpackSystemMessage(responseData, systemMessage, pushResponse)
	if err != nil {
		log.Errorf("push response unpack error:[%v]", err)
		return nil
	}
	return pushResponse
}

func (p *clusterCommunicationImplementor) NodeStorageInfo(from, to int32) error {
	packet := model.PackUDPSystemMessage(from, to, model.SystemMessageType_NODE_STORAGE_INFO, nil)
	ASyncSend(packet)
	return nil
}
func (p *clusterCommunicationImplementor) NodeStorageInfoReply(from, to int32, request *model.NodeStorageInfo) error {
	packet := model.PackUDPSystemMessage(from, to, model.SystemMessageType_NODE_STORAGE_INFO_REPLY, request)
	ASyncSend(packet)
	return nil
}

func (p *clusterCommunicationImplementor) PushReply(from, to int32, request *model.PushResponse) error {
	packet := model.PackTCPSystemMessage(from, to, model.SystemMessageType_DATA_PUSH_REPLY, request)
	ASyncSend(packet)
	return nil
}

func (p *clusterCommunicationImplementor) Pull(from, to int32, request *model.PullRequest) *model.PullResponse {
	packet := model.PackTCPSystemMessage(from, to, model.SystemMessageType_DATA_PULL, request)
	packetReply := SyncSend(packet)
	systemMessage := &model.SystemMessage{}
	response := &model.PullResponse{}
	err := model.UnpackSystemMessage(packetReply, systemMessage, response)
	if err != nil {
		log.Errorf("pull response unpack error:[%v]", err)
		return nil
	}
	return response
}

func (p *clusterCommunicationImplementor) PullReply(from, to int32, request *model.PullResponse) error {
	packet := model.PackTCPSystemMessage(from, to, model.SystemMessageType_DATA_PULL_REPLY, request)
	ASyncSend(packet)
	return nil
}

func (p *clusterCommunicationImplementor) MigrateRequest(from, to int32, request *model.DataMigrateRequest) error {
	packet := model.PackTCPSystemMessage(from, to, model.SystemMessageType_DATA_MIGRATE, request)
	ASyncSend(packet)
	return nil
}

func (p *clusterCommunicationImplementor) MigrateResponse(from, to int32, request *model.DataMigrateResponse) error {
	packet := model.PackTCPSystemMessage(from, to, model.SystemMessageType_DATA_MIGRATE_REPLY, request)
	ASyncSend(packet)
	return nil
}
