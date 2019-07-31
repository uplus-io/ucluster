package ucluster

import "github.com/uplus-io/ucluster/model"

type DataOperations interface {
	Push(dataArray []*model.DataBody) error
	Pull([]*model.DataBody) error
}

type DataPipeline interface {
}
