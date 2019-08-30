package v1

import (
	"github.com/uplus-io/ucluster/v1/model"
)

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
