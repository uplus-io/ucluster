package ucluster

import "github.com/uplus-io/ucluster/proto"

type ClusterDelegate interface {
	MessageEncode(bytes []byte) []byte
	MessageDecode(bytes []byte) []byte
}

type UClusterConfig struct {
	Group      string         `json:"group" yaml:"group"`
	Id         int32          `json:"id" yaml:"id"`
	Name       string         `json:"name" yaml:"name"`
	Role       proto.NodeRole `json:"role" yaml:"role"`
	Ip         string         `json:"ip" yaml:"ip"`
	SocketPort int32          `json:"socket_port" yaml:"socket_port"`
	HttpPort   int32          `json:"http_port" yaml:"http_port"`

	ClusterDelegate ClusterDelegate `json:"-" yaml:"-"`

	configPath string
}

func (p *UClusterConfig) Configure(config interface{}) {

}
func (p *UClusterConfig) GetConfigPath() string {
	return p.configPath
}
func (p *UClusterConfig) SetConfigPath(path string) {
	p.configPath = path
}
