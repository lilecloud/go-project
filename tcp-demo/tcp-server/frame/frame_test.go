package frame

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"testing"
)

func TestMyFrameCodec_Encode(t *testing.T) {
	type args struct {
		frame FramePayload
	}
	tests := []struct {
		name       string
		codec      *MyFrameCodec
		args       args
		wantWriter string
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name:       "test1",
			codec:      &MyFrameCodec{},
			args:       args{frame: FramePayload("this is a test")},
			wantWriter: "this is a test",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			if err := tt.codec.Encode(buffer, tt.args.frame); (err != nil) != tt.wantErr {
				t.Errorf("MyFrameCodec.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			frame, err := tt.codec.Decode(buffer)
			if tt.wantWriter != string(frame) {
				t.Errorf("MyFrameCodec.Encode() error = %v, wantErr %v", err, tt.wantErr)

			}
		})
	}
}

func TestMyFrameCodec_Decode(t *testing.T) {
	codec := &MyFrameCodec{}
	data := []byte{0x0, 0x0, 0x0, 0x9, 'h', 'e', 'l', 'l', 'o'}

	myFrame, err := codec.Decode(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
		return
	}
	if string(myFrame) != "hello" {
		fmt.Println("decode error ", string(myFrame))
	}
}

type returnErrReader struct {
	R         io.Reader
	ErrCount  uint16 // 第几次返回错误
	execCount uint16 // 记录执行的次数
}

func (r *returnErrReader) Read(p []byte) (n int, err error) {
	r.execCount++
	if r.execCount == r.ErrCount {
		return 0, errors.New("want err")
	}
	return r.Read(p)
}

func TestMyFrameCodec_Decode_WantErr(t *testing.T) {
	codec := MyFrameCodec{}
	data := []byte{0x0, 0x0, 0x0, 0x9, 'h', 'e', 'l', 'l', 'o'}

	_, err := codec.Decode(&returnErrReader{
		R:        bytes.NewReader(data),
		ErrCount: 1,
	})
	if err == nil {
		t.Error("not got want error")
	}

	_, err1 := codec.Decode(&returnErrReader{
		R:        bytes.NewReader(data),
		ErrCount: 2,
	})
	if err1 == nil {
		t.Error("want non-nil,actual nil")
	}

}
