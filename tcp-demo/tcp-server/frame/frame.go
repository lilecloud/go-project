package frame

import (
	"encoding/binary"
	"errors"
	"io"
)

const (
	FRAME_HEAER_LEN uint32 = 4
)

var ErrorShortWrite = errors.New("write short")
var ErrorShortRead = errors.New("read short")

// 每一个tcp包的抽象
type FramePayload []byte

// tcp包编码解码的接口

type FrameCodec interface {
	/**
	将framePayload编码之后，写入writer
	*/
	Encode(io.Writer, FramePayload) error
	// 从reader中读取数据然后，封装入framepayload然后返回
	Decode(io.Reader) (FramePayload, error)
}

type MyFrameCodec struct{}

/**

编码
*/
func (codec *MyFrameCodec) Encode(writer io.Writer, frame FramePayload) error {
	var frameLen int = len(frame)
	var totalLen uint32 = uint32(frameLen) + FRAME_HEAER_LEN

	err := binary.Write(writer, binary.BigEndian, &totalLen)
	if err != nil {
		return err
	}

	n, err := writer.Write([]byte(frame))
	if err != nil {
		return err
	}

	if n != frameLen {
		return ErrorShortWrite
	}
	return nil

}

func (codec *MyFrameCodec) Decode(reader io.Reader) (FramePayload, error) {
	var totalLen uint32
	err := binary.Read(reader, binary.BigEndian, &totalLen)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, totalLen-FRAME_HEAER_LEN)
	n, err := io.ReadFull(reader, buf)
	if err != nil {
		return nil, err
	}

	if n != int(totalLen-FRAME_HEAER_LEN) {
		return nil, ErrorShortRead
	}
	return FramePayload(buf), nil
}
