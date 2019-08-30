package network

import (
	"fmt"
	"net"
)

type TcpTemplate struct {
	ip        string
	port      int32
	conn      *net.TCPConn
	splitable Splitable
}

func NewTcpTemplate(ip string, port int32, splitable Splitable) *TcpTemplate {
	return &TcpTemplate{ip: ip, port: port, splitable: splitable}
}

func (p *TcpTemplate) Connect() error {
	server := fmt.Sprintf("%s:%d", p.ip, p.port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		return err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	p.conn = conn
	return nil
}

func (p *TcpTemplate) Disconnect() {
	p.conn.Close()
}

func (p *TcpTemplate) Send(bytes []byte) []byte {
	readWriter := NewReadWriter(p.conn)
	_, err := readWriter.Writer().Write(bytes)
	if err != nil {
		fmt.Printf("write data error:%v", err.Error())
		return nil
	}

	reader := readWriter.Reader(p.splitable)
	dataChan := make(chan []byte, 1)

	go func() {
		for {
			if reader.HasNext() {
				data := reader.Read()
				dataChan <- data
				return
			}
		}
	}()

	reply := <-dataChan
	return reply
}
