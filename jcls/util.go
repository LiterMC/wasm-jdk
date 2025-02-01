package jcls

import (
	"io"
)

func readUint8(r io.Reader) (uint8, error) {
	var bts [1]byte
	if _, err := io.ReadFull(r, bts[:]); err != nil {
		return 0, err
	}
	return bts[0], nil
}

func readUint16(r io.Reader) (uint16, error) {
	var bts [2]byte
	if _, err := io.ReadFull(r, bts[:]); err != nil {
		return 0, err
	}
	return ((uint16)(bts[0]) << 8) | (uint16)(bts[1]), nil
}

func readUint32(r io.Reader) (uint32, error) {
	var bts [4]byte
	if _, err := io.ReadFull(r, bts[:]); err != nil {
		return 0, err
	}
	return ((uint32)(bts[0]) << 24) | ((uint32)(bts[1]) << 16) | ((uint32)(bts[2]) << 8) | (uint32)(bts[3]), nil
}

func readUint64(r io.Reader) (uint64, error) {
	var bts [8]byte
	if _, err := io.ReadFull(r, bts[:]); err != nil {
		return 0, err
	}
	return ((uint64)(bts[0]) << 56) | ((uint64)(bts[1]) << 48) |
		((uint64)(bts[2]) << 40) | ((uint64)(bts[3]) << 32) |
		((uint64)(bts[4]) << 24) | ((uint64)(bts[5]) << 16) |
		((uint64)(bts[6]) << 8) | (uint64)(bts[7]), nil
}
