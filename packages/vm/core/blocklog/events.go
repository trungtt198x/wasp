package blocklog

import (
	"io"

	"github.com/iotaledger/wasp/packages/kv/codec"
)

// EventLookupKey is a globally unique reference to the event:
// block index + index of the request within block + index of the event within the request
type EventLookupKey [8]byte

func NewEventLookupKey(blockIndex uint32, requestIndex, eventIndex uint16) EventLookupKey {
	ret := EventLookupKey{}
	copy(ret[:4], codec.EncodeUint32(blockIndex))
	copy(ret[4:6], codec.EncodeUint16(requestIndex))
	copy(ret[6:8], codec.EncodeUint16(eventIndex))
	return ret
}

func (k EventLookupKey) BlockIndex() uint32 {
	return codec.MustDecodeUint32(k[:4])
}

func (k EventLookupKey) RequestIndex() uint16 {
	return codec.MustDecodeUint16(k[4:6])
}

func (k EventLookupKey) RequestEventIndex() uint16 {
	return codec.MustDecodeUint16(k[6:8])
}

func (k EventLookupKey) Bytes() []byte {
	return k[:]
}

func (k *EventLookupKey) Write(w io.Writer) error {
	_, err := w.Write(k[:])
	return err
}

func EventLookupKeyFromBytes(r io.Reader) (*EventLookupKey, error) {
	k := EventLookupKey{}
	n, err := r.Read(k[:])
	if err != nil || n != 8 {
		return nil, io.EOF
	}
	return &k, nil
}
