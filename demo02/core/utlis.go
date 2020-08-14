package core

import "encoding/binary"

func IntToHex(i int64) []byte {
	buf := make([]byte, 8, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}
