package helpers

import "bytes"

type BitSet struct {
	entries []byte
}

func NewBitSet(bitCount uint) BitSet {
	return BitSet{make([]byte, (bitCount+7)/8)}
}

func (bs BitSet) HasBit(bit uint) bool {
	return (bs.entries[bit/8] & (1 << (bit & 7))) != 0
}

func (bs BitSet) SetBit(bit uint) {
	bs.entries[bit/8] |= 1 << (bit & 7)
}

func (bs BitSet) Subtract(other BitSet) {
	if len(bs.entries) != len(other.entries) {
		panic("subtract bitset must be same size")
	}

	for idx := range bs.entries {
		bs.entries[idx] = bs.entries[idx] & (^other.entries[idx])
	}
}

func (bs BitSet) Merge(other BitSet) {
	if len(bs.entries) != len(other.entries) {
		panic("subtract bitset must be same size")
	}

	for idx := range bs.entries {
		bs.entries[idx] = bs.entries[idx] | other.entries[idx]
	}
}

func (bs BitSet) Contains(other BitSet) bool {
	if len(bs.entries) != len(other.entries) {
		return false
	}

	for idx := range bs.entries {
		curBs := bs.entries[idx]
		curOther := other.entries[idx]
		if (curBs^curOther)&curOther != 0 {
			return false
		}
	}

	return true
}

func (bs BitSet) Equals(other BitSet) bool {
	return bytes.Equal(bs.entries, other.entries)
}

func (bs BitSet) String() string {
	return string(bs.entries)
}

func (bs BitSet) Clone() BitSet {
	res := bs
	res.entries = make([]byte, len(bs.entries))
	copy(res.entries, bs.entries)
	return res
}
