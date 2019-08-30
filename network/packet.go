package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/uplus-io/ucluster/proto"
	"github.com/uplus-io/ugo/goproto"
)

const (
	CONST_HEADER_LENGTH int = 4
	HEADER_FORMAT           = "UP.%d"
)

func Pack(message proto.Message) ([]byte, error) {
	var err error
	var length,headerLength int32

	data, err := goproto.Marshal(&message)
	length = int32(len(data))

	header := fmt.Sprintf(HEADER_FORMAT, message.Command)
	headerData := []byte(header)
	headerLength = int32(len(headerData))

	length += headerLength
	length += 8

	buffer := new(bytes.Buffer)
	err = binary.Write(buffer, binary.LittleEndian, &length)
	err = binary.Write(buffer, binary.LittleEndian, &headerLength)
	err = binary.Write(buffer, binary.LittleEndian, &headerData)
	err = binary.Write(buffer, binary.LittleEndian, &data)
	return buffer.Bytes(), err
}

func Unpack(source []byte) (proto.Message, error) {
	reader := bytes.NewReader(source)
	var length int32
	var headerLength int32
	var headerData []byte
	var data []byte
	var err error
	err = binary.Read(reader, binary.LittleEndian, &length)
	err = binary.Read(reader, binary.LittleEndian, &headerLength)
	headerData = make([]byte,headerLength)
	err = binary.Read(reader, binary.LittleEndian, &headerData)
	length -= headerLength
	length -= 8
	data = make([]byte,length)
	err = binary.Read(reader, binary.LittleEndian, &data)
	message := &proto.Message{}
	err = goproto.Unmarshal(data, message)
	return *message, err
}
