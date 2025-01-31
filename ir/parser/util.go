package parser

import (
	"io"
)

func readUint16(r io.Reader) (uint16, error) {
	var bts [2]byte
	if _, err := io.ReadFull(r, bts[:]); err != nil {
		return 0, err
	}
	return ((uint16)(bts[0]) << 8) | (uint16)(bts[1]), nil
}

func readInt16(r io.Reader) (int16, error) {
	var bts [2]byte
	if _, err := io.ReadFull(r, bts[:]); err != nil {
		return 0, err
	}
	return ((int16)(bts[0]) << 8) | (int16)(bts[1]), nil
}

func readInt32(r io.Reader) (int32, error) {
	var bts [4]byte
	if _, err := io.ReadFull(r, bts[:]); err != nil {
		return 0, err
	}
	return ((int32)(bts[0]) << 24) | ((int32)(bts[1]) << 16) | ((int32)(bts[2]) << 8) | (int32)(bts[3]), nil
}
