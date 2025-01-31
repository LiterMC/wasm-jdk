package parser

func bytesToUint16(bts []byte) uint16 {
	return ((uint16)(bts[0]) << 8) | (uint16)(bts[1])
}

func bytesToInt16(bts []byte) int16 {
	return ((int16)(bts[0]) << 8) | (int16)(bts[1])
}

func bytesToInt32(bts []byte) int32 {
	return ((int32)(bts[0]) << 24) | ((int32)(bts[1]) << 16) | ((int32)(bts[2]) << 8) | (int32)(bts[3])
}
