package bitmap

func GetValue(from, offset, leng uint) uint {
	return (from >> offset) & getMask(leng)
}

func GetValueUint8(from uint8, offset, leng uint) uint8 {
	return uint8(GetValue(uint(from), offset, leng))
}

func GetValueUint16(from uint16, offset, leng uint) uint16 {
	return uint16(GetValue(uint(from), offset, leng))
}

func GetValueUint32(from uint32, offset, leng uint) uint32 {
	return uint32(GetValue(uint(from), offset, leng))
}

func SetValue(to, val, offset uint) uint {
	return to + (val << offset)
}

func SetValueUint8(to, val uint8, offset uint) uint8 {
	return uint8(SetValue(uint(to), uint(val), offset))
}

func SetValueUint16(to, val uint16, offset uint) uint16 {
	return uint16(SetValue(uint(to), uint(val), offset))
}

func SetValueUint32(to, val uint32, offset uint) uint32 {
	return uint32(SetValue(uint(to), uint(val), offset))
}

func getMask(leng uint) uint {
	return 2<<leng - 1
}
