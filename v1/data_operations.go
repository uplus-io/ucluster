package v1

import "github.com/uplus-io/ucluster/v1/model"

type DataOperations interface {
	Migrate(from, to int32, startRing int32, endRing int32) error
	Push(from, to int32, dataArray []*model.DataBody)
	Pull(from, to int32, dataArray []*model.DataBody) error
}

type DataPipeline interface {
}
