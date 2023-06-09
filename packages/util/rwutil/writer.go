// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package rwutil

import (
	"bytes"
	"encoding"
	"errors"
	"io"
	"time"

	"github.com/iotaledger/hive.go/serializer/v2"
	"github.com/iotaledger/hive.go/serializer/v2/marshalutil"
	iotago "github.com/iotaledger/iota.go/v3"
)

type Writer struct {
	Err error
	w   io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func NewBytesWriter() *Writer {
	return NewWriter(new(bytes.Buffer))
}

func (ww *Writer) Bytes() []byte {
	buf, ok := ww.w.(*bytes.Buffer)
	if !ok {
		panic("writer expects bytes buffer")
	}
	if ww.Err != nil {
		// writing to bytes buffer should never have failed
		panic(ww.Err)
	}
	return buf.Bytes()
}

func (ww *Writer) Write(writer interface{ Write(w io.Writer) error }) *Writer {
	if ww.Err == nil {
		ww.Err = writer.Write(ww.w)
	}
	return ww
}

func (ww *Writer) WriteN(val []byte) *Writer {
	if ww.Err == nil {
		ww.Err = WriteN(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteAddress(val iotago.Address) *Writer {
	if ww.Err == nil {
		buf, _ := val.Serialize(serializer.DeSeriModeNoValidation, nil)
		ww.WriteN(buf)
	}
	return ww
}

func (ww *Writer) WriteBool(val bool) *Writer {
	if ww.Err == nil {
		ww.Err = WriteBool(ww.w, val)
	}
	return ww
}

//nolint:govet
func (ww *Writer) WriteByte(val byte) *Writer {
	if ww.Err == nil {
		ww.Err = WriteByte(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteBytes(val []byte) *Writer {
	if ww.Err == nil {
		ww.Err = WriteBytes(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteDuration(val time.Duration) *Writer {
	return ww.WriteInt64(int64(val))
}

func (ww *Writer) WriteFromBytes(bytes interface{ Bytes() []byte }) *Writer {
	if ww.Err == nil {
		ww.WriteBytes(bytes.Bytes())
	}
	return ww
}

func (ww *Writer) WriteInt8(val int8) *Writer {
	if ww.Err == nil {
		ww.Err = WriteInt8(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteInt16(val int16) *Writer {
	if ww.Err == nil {
		ww.Err = WriteInt16(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteInt32(val int32) *Writer {
	if ww.Err == nil {
		ww.Err = WriteInt32(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteInt64(val int64) *Writer {
	if ww.Err == nil {
		ww.Err = WriteInt64(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteMarshaled(m encoding.BinaryMarshaler) *Writer {
	if ww.Err == nil {
		if m == nil {
			ww.Err = errors.New("nil marshaler")
			return ww
		}
		var buf []byte
		buf, ww.Err = m.MarshalBinary()
		ww.WriteBytes(buf)
	}
	return ww
}

func (ww *Writer) WriteMessageType(msgType byte) *Writer {
	return ww.WriteByte(msgType)
}

type serializable interface {
	Serialize(serializer.DeSerializationMode, interface{}) ([]byte, error)
}

func (ww *Writer) WriteSerialized(s serializable) *Writer {
	if ww.Err == nil {
		if s == nil {
			ww.Err = errors.New("nil deserializer")
			return ww
		}
		var buf []byte
		buf, ww.Err = s.Serialize(serializer.DeSeriModeNoValidation, nil)
		ww.WriteBytes(buf)
	}
	return ww
}

func (ww *Writer) WriteSize(val int) *Writer {
	return ww.WriteSize32(uint32(val))
}

func (ww *Writer) WriteSize32(val uint32) *Writer {
	if ww.Err == nil {
		ww.Err = WriteSize32(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteString(val string) *Writer {
	if ww.Err == nil {
		ww.Err = WriteString(ww.w, val)
	}
	return ww
}

type marshalUtilWriter interface {
	WriteToMarshalUtil(mu *marshalutil.MarshalUtil)
}

func (ww *Writer) WriteToMarshalUtil(m marshalUtilWriter) *Writer {
	if ww.Err == nil {
		mu := marshalutil.New()
		m.WriteToMarshalUtil(mu)
		ww.WriteN(mu.Bytes()[:mu.WriteOffset()])
	}
	return ww
}

func (ww *Writer) WriteUint8(val uint8) *Writer {
	if ww.Err == nil {
		ww.Err = WriteUint8(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteUint16(val uint16) *Writer {
	if ww.Err == nil {
		ww.Err = WriteUint16(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteUint32(val uint32) *Writer {
	if ww.Err == nil {
		ww.Err = WriteUint32(ww.w, val)
	}
	return ww
}

func (ww *Writer) WriteUint64(val uint64) *Writer {
	if ww.Err == nil {
		ww.Err = WriteUint64(ww.w, val)
	}
	return ww
}
