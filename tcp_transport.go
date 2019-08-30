package ucluster

import (
	"fmt"
	"github.com/uplus-io/ucluster/network"
	"net"
)

type TcpTransport struct {
	config UClusterConfig
}

func NewTcpTransport(config UClusterConfig) *TcpTransport {
	return &TcpTransport{config: config}
}

func (p *TcpTransport) Listen() {
	address := fmt.Sprintf("%s:%d", p.config.Ip, p.config.SocketPort)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}

	fmt.Printf("tcp transport listen at %s \n", address)
	//接受客户端信息
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}
		//用协程建立连接
		go p.process(conn)
	}
}

func (p *TcpTransport) process(conn net.Conn) {
	defer conn.Close()

	for {
		readWriter := network.NewReadWriter(conn)
		reader := readWriter.Reader(network.DefaultSpliterInstance)
		for reader.HasNext() {
			readData := reader.Read()
			unpackData, _ := network.Unpack(readData)
			//bytes := p.config.ClusterDelegate.MessageDecode(unpackData)
			//message := &proto.Message{}
			//err := goproto.Unmarshal(bytes, message)
			//if err != nil {
			//	todo://处理异常
			//}
			fmt.Println("%v", unpackData)
		}
	}
}
