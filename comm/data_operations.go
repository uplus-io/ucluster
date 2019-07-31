package comm

import (
	"github.com/uplus-io/uengine/model"
)

type DataOperations interface {
	Push(dataArray []*model.DataBody) error
	Pull([]*model.DataBody) error
}

type DataPipeline interface {
}
