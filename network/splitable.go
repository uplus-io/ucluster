package network

import (
	"bytes"
	"encoding/binary"
)

var DefaultSpliterInstance *DefaultSpliter

type Splitable interface {
	Split(data []byte, isEOF bool) (advance int, token []byte, err error)
}

type DefaultSpliter struct {

}

func init() {
	DefaultSpliterInstance = &DefaultSpliter{}
}

func (p *DefaultSpliter) Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) > CONST_HEADER_LENGTH {
		length := int32(0)
		binary.Read(bytes.NewReader(data[:4]), binary.LittleEndian, &length)
		if int(length) <= len(data) {
			return int(length), data[:int(length)], nil
		}
	}
	return
}


