package ucluster

import (
	"github.com/uplus-io/ucluster/model"
	"github.com/uplus-io/ugo/hash"
	log "github.com/uplus-io/ugo/logger"
	"sync"
)

const pageSize = 100

type LocalDataOperations struct {
	comm     clusterCommunication
	delegate ClusterDataDelegate

	queue []*model.DataBody
	sync.Mutex
}

func newLocalDataOperations(comm clusterCommunication, delegate ClusterDataDelegate) *LocalDataOperations {
	return &LocalDataOperations{comm: comm, delegate: delegate, queue: make([]*model.DataBody, 0)}
}

//1 3 5 7 9 11
//2
func (p *LocalDataOperations) Migrate(from, to int32, startRing int32, endRing int32) error {
	delegate := p.delegate
	delegate.ForEach(func(data *model.DataBody) bool {
		p.Lock()
		defer p.Unlock()
		ring := hash.Int32(data.Id)
		if startRing < endRing {
			if ring >= startRing && ring < endRing {
				exists, err := delegate.Get(data)
				if err != nil {
					log.Errorf("read local data error:[%v]", err)
				} else if exists {
					p.put(from, to, data, err)
				}
			}
		} else {
			if ring >= startRing {
				exists, err := delegate.Get(data)
				if err != nil {
					log.Errorf("read local data error:[%v]", err)
				} else if exists {
					p.put(from, to, data, err)
				}
			}
		}
		return true
	})
	if len(p.queue) > 0 {
		p.pushData(from,to)
	}
	response := &model.DataMigrateResponse{Completed: true}
	return p.comm.MigrateResponse(to, from, response)
}

func (p *LocalDataOperations) put(from, to int32, data *model.DataBody, err error) {
	p.queue = append(p.queue, data)

	if len(p.queue) >= pageSize {
		p.pushData(from, to)
	}
}

func (p *LocalDataOperations) pushData(from, to int32) {
	p.Lock()
	defer p.Unlock()
	pushRequest := &model.PushRequest{Data: p.queue}

	response := p.comm.Push(to, from, pushRequest)
	if response != nil && response.Success {
		p.queue = make([]*model.DataBody, 0)
	}
}

// A节点将数据(key,value,version)及对应的版本号推送给B节点
// B节点更新A发过来的数据中比自己新的数据
func (p *LocalDataOperations) Push(from, to int32, dataArray []*model.DataBody) {
	result := make([]*model.DataBody, len(dataArray))
	for i, data := range dataArray {
		exist := &model.DataBody{Namespace: data.Namespace, Table: data.Table, Name: data.Name, Key: data.Key}
		exists, err := p.delegate.Get(exist)
		if err != nil {
			log.Errorf("read local data error:[%v]", err)
		} else {
			if !exists || exist.Version < data.Version {
				err := p.delegate.Set(data)
				if err != nil {
					log.Errorf("write local data error:[%v]", err)
				}
				log.Debugf(
					"local data older,update[%s]ring:%d from v:%d to v:%d",
					string(data.Id), uint32(data.Ring), exist.Version, data.Version)
			}
		}
		result[i] = &model.DataBody{Namespace: data.Namespace, Table: data.Table, Id: data.Id, Version: data.Version}
	}
	response := &model.PushResponse{Success: true, Data: result}
	p.comm.PushReply(to, from, response)
}

// A不发送数据的value，仅发送数据的摘要key和version给B。
// B根据版本比较数据，将本地比A新的数据(key,value,version)推送给A
// A更新自己的本地数据
func (p *LocalDataOperations) Pull(from, to int32, dataArray []*model.DataBody) error {
	return nil
}
