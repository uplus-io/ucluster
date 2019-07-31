package ucluster

import "github.com/uplus-io/ucluster/model"

type ClusterDelegate interface {
	LocalNodeStorageInfo() *model.NodeStorageInfo
	LocalNodeStorageStat() *model.NodeStorageStat
	LocalNodeHealthStat() *model.NodeHealthStat
}

type ClusterMessageDelegate interface {
	System(pipeline Pipeline, message model.SystemMessage) error
	Event(pipeline Pipeline, message model.EventMessage) error
	Topic(pipeline Pipeline, message model.TopicMessage) error
	Data(pipeline Pipeline, message model.DataMessage) error
}

type DataIterator func(data *model.DataBody) bool
type ClusterDataDelegate interface {
	ForEach(iterator DataIterator)
	Get(data *model.DataBody) (bool, error)
	Set(data *model.DataBody) error
}

//集群通讯
type clusterCommunication interface {
	SendNodeInfoTo(to int32) error
}

//集群数据通讯
type clusterDataCommunication interface {
	Push(from, to int32, request *model.PushRequest) *model.PushResponse
	PushReply(from, to int32, request *model.PushResponse) error
	Pull(from, to int32, request *model.PullRequest) *model.PullResponse
	PullReply(from, to int32, request *model.PullResponse) error
	MigrateRequest(from, to int32, request *model.DataMigrateRequest) error
	MigrateResponse(from, to int32, request *model.DataMigrateResponse) error
}
