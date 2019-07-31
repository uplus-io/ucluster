package ucluster

import (
	"github.com/uplus-io/ucluster/model"
	log "github.com/uplus-io/ugo/logger"
)

type clusterDataCommunicationImplementor struct {
	pipeline Pipeline
}

func newClusterDataCommunicationImplementor(pipeline Pipeline) *clusterDataCommunicationImplementor {
	return &clusterDataCommunicationImplementor{pipeline: pipeline}
}

func (p *clusterDataCommunicationImplementor) Push(from, to int32, request *model.PushRequest) *model.PushResponse {
	packet := model.PackTCPSystemMessage(from, to, model.SystemMessageType_DataPush, request)

	responseData := p.pipeline.SyncSend(packet)
	systemMessage := &model.SystemMessage{}
	pushResponse := &model.PushResponse{}
	err := model.UnpackSystemMessage(responseData, systemMessage, pushResponse)
	if err != nil {
		log.Errorf("push response unpack error:[%v]", err)
		return nil
	}
	return pushResponse
}

func (p *clusterDataCommunicationImplementor) PushReply(from, to int32, request *model.PushResponse) error {
	packet := model.PackTCPSystemMessage(from, to, model.SystemMessageType_DataPushReply, request)
	p.pipeline.ASyncSend(packet)
	return nil
}

func (p *clusterDataCommunicationImplementor) Pull(from, to int32, request *model.PullRequest) *model.PullResponse {
	packet := model.PackTCPSystemMessage(from, to, model.SystemMessageType_DataPull, request)
	packetReply := p.pipeline.SyncSend(packet)
	systemMessage := &model.SystemMessage{}
	response := &model.PullResponse{}
	err := model.UnpackSystemMessage(packetReply, systemMessage, response)
	if err != nil {
		log.Errorf("pull response unpack error:[%v]", err)
		return nil
	}
	return response
}

func (p *clusterDataCommunicationImplementor) PullReply(from, to int32, request *model.PullResponse) error {
	packet := model.PackTCPSystemMessage(from, to, model.SystemMessageType_DataPullReply, request)
	p.pipeline.ASyncSend(packet)
	return nil
}

func (p *clusterDataCommunicationImplementor) MigrateRequest(from, to int32, request *model.DataMigrateRequest) error {
	packet := model.PackTCPSystemMessage(from, to, model.SystemMessageType_DataMigrate, request)
	p.pipeline.ASyncSend(packet)
	return nil
}

func (p *clusterDataCommunicationImplementor) MigrateResponse(from, to int32, request *model.DataMigrateResponse) error {
	packet := model.PackTCPSystemMessage(from, to, model.SystemMessageType_DataMigrateReply, request)
	p.pipeline.ASyncSend(packet)
	return nil
}
