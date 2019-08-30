package v1

import "github.com/uplus-io/ucluster/v1/model"

type RemoteDataOperations struct {
}

func (p *RemoteDataOperations) Push(dataArray []*model.DataBody) error {
	return nil
}
func (p *RemoteDataOperations) Pull([]*model.DataBody) error {
	return nil
}
