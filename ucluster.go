package ucluster

import "github.com/pkg/errors"

const PROTO_VERSION = 1

type UClusterConfig struct {
	Name          string   `json:"level" yaml:"name"`
	Seeds         []string `json:"seeds" yaml:"seeds"`
	BindIp        []string `json:"bind_ip" yaml:"bind_ip"`
	BindPort      int      `json:"bind_port" yaml:"bind_port"`
	AdvertisePort int      `json:"advertise_port" yaml:"advertise_port"`
	ReplicaCount  int      `json:"replica_count" yaml:"replica_count"`
	Secret        string   `json:"secret" yaml:"secret"`

	Delegate       ClusterDelegate        `json:"-" yaml:"-"`
	PacketDelegate ClusterMessageDelegate `json:"-" yaml:"-"`
}

type UCluster struct {
	config UClusterConfig
}

func Serving(config UClusterConfig) error {
	if config.Delegate == nil {
		return errors.New("cluster[Delegate] not implement")
	}
	if config.PacketDelegate == nil {
		return errors.New("cluster[PacketDelegate] not implement")
	}
	cluster := NewCluster(config)
	cluster.Listen()
	return nil
}
