package ucluster

import "github.com/uplus-io/ucluster/model"

type RemoteDataOperations struct {
}

func (p *RemoteDataOperations) Push(dataArray []*model.DataBody) error {
	return nil
}
func (p *RemoteDataOperations) Pull([]*model.DataBody) error {
	return nil
}
