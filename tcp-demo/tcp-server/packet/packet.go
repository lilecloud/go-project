package packet

import (
	"bytes"
	"errors"
)

// commandId 的种类
const (
	// 链接请求
	connCommand uint8 = iota + 1
	// 消息请求
	sbmitCommand
)

const (
	// 链接请求
	connAck uint8 = iota + 100
	// 消息请求
	submitAck
)

// cmdId 一个字节
func Decode(packet []byte) (Packet, error) {
	cmdId := uint8(packet[0])
	body := packet[1:]

	switch cmdId {
	case connCommand:
		connPack := ConnPacket{}
		err := connPack.Decode(body)
		if err != nil {
			return nil, err
		}
		return &connPack, err
	case connAck:
		cnnAck := ConnAckPacket{}
		err := cnnAck.Decode(body)
		if err != nil {
			return nil, err
		}
		return &cnnAck, err
	case sbmitCommand:
		submitPacket := SubmitPacket{}
		err := submitPacket.Decode(body)
		if err != nil {
			return nil, err
		}
		return &submitPacket, err
	case submitAck:
		submitAck := SubmitAckPacket{}
		err := submitAck.Decode(body)
		if err != nil {
			return nil, err
		}
		return &submitAck, err
	default:
		return nil, errors.New("unKnown commandId")
	}
}

func Encode(p Packet) ([]byte, error) {
	var cmdId uint8

	switch p.(type) {
	case *ConnPacket:
		cmdId = connCommand
	case *ConnAckPacket:
		cmdId = connAck
	case *SubmitPacket:
		cmdId = sbmitCommand
	case *SubmitAckPacket:
		cmdId = submitAck
	default:
		return nil, errors.New("the packet type not support")
	}

	bs, err := p.Encode()
	if err != nil {
		return nil, err
	}
	return bytes.Join([][]byte{{cmdId}, bs}, nil), nil
}

type Packet interface {
	Decode([]byte) error     // 将 [] byte 转成packet
	Encode() ([]byte, error) // 将packet转换成[]byte
}

type ConnPacket struct {
	Id      string
	Payload []byte
}

// 将packet转换成[]byte
func (c *ConnPacket) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(c.Id[:8]), c.Payload}, nil), nil
}

// 将 [] byte 转成packet
func (c *ConnPacket) Decode(data []byte) error {
	c.Id = string(data[:8])
	c.Payload = data[8:]
	return nil
}

type ConnAckPacket struct {
	Id     string
	Result uint8
}

// 将packet转换成[]byte
func (c *ConnAckPacket) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(c.Id[:8]), {c.Result}}, nil), nil
}

// 将 [] byte 转成packet
// result 1 成功 0 失败
func (c *ConnAckPacket) Decode(data []byte) error {
	c.Id = string(data[:8])
	c.Result = uint8(data[8])
	if c.Result == 0 {
		return errors.New("connect error")
	}
	return nil
}

type SubmitPacket struct {
	Id      string
	Payload []byte
}

// 将packet转换成[]byte
func (c *SubmitPacket) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(c.Id[:8]), c.Payload}, nil), nil
}

// 将 [] byte 转成packet
func (c *SubmitPacket) Decode(data []byte) error {
	c.Id = string(data[:8])
	c.Payload = data[8:]
	return nil
}

type SubmitAckPacket struct {
	Id     string
	Result uint8
}

// 将packet转换成[]byte
func (c *SubmitAckPacket) Encode() ([]byte, error) {
	return bytes.Join([][]byte{[]byte(c.Id[:8]), {c.Result}}, nil), nil
}

// 将 [] byte 转成packet
func (c *SubmitAckPacket) Decode(data []byte) error {
	c.Id = string(data[:8])
	c.Result = uint8(data[8])
	return nil
}
