package ucluster

import "errors"

const PROTO_VERSION = 1

type UClusterConfig struct {
	Name          string   `json:"level" yaml:"name"`
	Seeds         []string `json:"seeds" yaml:"seeds"`
	BindIp        []string `json:"bind_ip" yaml:"bind_ip"`
	BindPort      int      `json:"bind_port" yaml:"bind_port"`
	AdvertisePort int      `json:"advertise_port" yaml:"advertise_port"`
	ReplicaCount  int      `json:"replica_count" yaml:"replica_count"`
	Secret        string   `json:"secret" yaml:"secret"`

	Delegate        ClusterDelegate        `json:"-" yaml:"-"`
	DataDelegate    ClusterDataDelegate    `json:"-" yaml:"-"`
	MessageDelegate ClusterMessageDelegate `json:"-" yaml:"-"`
}

type UCluster struct {
	config  *UClusterConfig
	cluster *Cluster
}

func NewUCluster(config *UClusterConfig) *UCluster {
	return &UCluster{config: config}
}

func (p *UCluster) Serving() (error) {
	err := assertConfig(p.config)
	if err != nil {
		return err
	}
	cluster := NewCluster(*p.config)
	p.cluster = cluster
	cluster.Listen()
	return nil
}

func assertConfig(config *UClusterConfig) error {
	if config.Delegate == nil {
		return errors.New("[Delegate]must implement")
	}
	if config.DataDelegate == nil {
		return errors.New("[DataDelegate]must implement")
	}
	if config.MessageDelegate == nil {
		return errors.New("[MessageDelegate]must implement")
	}
	return nil
}
