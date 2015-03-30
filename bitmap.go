package iso8583

import (
	"errors"
)

type BitMap struct {
	a uint64
	b uint64
}

func (bm *BitMap) SetBit(n int) error {
	if n <= 1 || n > 128 {
		return errors.New("iso8583.SetBit: field <= 1 or field > 128")
	}
	if n <= 64 {
		bm.a |= (uint64(1) << uint(64-n))
	} else {
		bm.b |= (uint64(1) << uint(128-n))
		bm.a |= 0x8000000000000000
	}
	return nil
}

func (bm *BitMap) DelBit(n int) error {
	if n <= 1 || n > 128 {
		return errors.New("iso8583.SetBit: field <= 1 or field > 128")
	}
	if n <= 64 {
		bm.a &= ^(uint64(1) << uint(64-n))
	} else {
		bm.b &= ^(uint64(1) << uint(128-n))
		if bm.b == 0 {
			bm.a &= 0x7FFFFFFFFFFFFFFF
		}
	}
	return nil
}

func (bm *BitMap) Clear() {
	bm.a = 0
	bm.b = 0
}

func (bm BitMap) ToBytes() []byte {
	var b [16]byte
	for i := uint(0); i < 8; i++ {
		b[7-i] = byte(bm.a >> (i * 8))
	}
	if bm.b == 0 {
		return b[:8]
	} else {
		for j := uint(0); j < 8; j++ {
			b[15-j] = byte(bm.b >> (j * 8))
		}
		return b[:]
	}
}
