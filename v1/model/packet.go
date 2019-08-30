package model

import (
	"github.com/uplus-io/ugo/proto"
	"github.com/uplus-io/ugo/utils"
)

func NewPacket(mode PacketMode, typ PacketType, from int32, to []int32, content []byte) *Packet {
	packet := &Packet{
		Version: 1,
		Id:      utils.GenId(),
		Mode:    mode,
		Type:    typ,
		From:    from,
		Content: content,
	}
	if mode == PacketMode_Multicast {
		packet.Receivers = to
	} else {
		packet.To = to[0]
	}
	return packet
}

func NewTCPPacket(typ PacketType, from, to int32, content []byte) *Packet {
	return NewPacket(PacketMode_TCP, typ, from, []int32{to}, content)
}

func NewUDPPacket(typ PacketType, from, to int32, content []byte) *Packet {
	return NewPacket(PacketMode_UDP, typ, from, []int32{to}, content)
}

func PackSystemMessage(mode PacketMode, from, to int32, messageType SystemMessageType, message proto.ProtoMessage) *Packet {
	var messageData []byte
	if message != nil {
		messageData, _ = proto.Marshal(message)
	}
	systemMessage := &SystemMessage{Type: messageType, Sender: from, Content: messageData}
	systemMessageData, _ := proto.Marshal(systemMessage)
	if mode == PacketMode_TCP {
		return NewTCPPacket(PacketType_System, from, to, systemMessageData)
	} else {
		return NewUDPPacket(PacketType_System, from, to, systemMessageData)
	}
}

func PackTCPSystemMessage(from, to int32, messageType SystemMessageType, message proto.ProtoMessage) *Packet {
	return PackSystemMessage(PacketMode_TCP, from, to, messageType, message)
}

func PackUDPSystemMessage(from, to int32, messageType SystemMessageType, message proto.ProtoMessage) *Packet {
	return PackSystemMessage(PacketMode_TCP, from, to, messageType, message)
}

func UnpackSystemMessage(packet *Packet, systemMessage *SystemMessage, message proto.ProtoMessage) (err error) {
	err = proto.Unmarshal(packet.Content, systemMessage)
	if err != nil {
		return
	}
	err = proto.Unmarshal(Content, message)
	if err != nil {
		return
	}
	return
}
