package byter

import (
	"encoding/binary"
	"reflect"
)

func OfUint8(val uint8) []byte {
	return []byte{val}
}

func ToUint8(data []byte) uint8 {
	return data[0]
}

func EatUint8(data []byte) ([]byte, uint8) {
	val := data[0]
	data = data[1:]
	return data, val
}

func TryEatUint8(data []byte) ([]byte, uint8, error) {
	if len(data) < 1 {
		return data, 0, ErrNotEnough
	}
	var val uint8
	data, val = EatUint8(data)
	return data, val, nil
}

func OfUint16(val uint16) []byte {
	r := make([]byte, 2)
	binary.BigEndian.PutUint16(r, val)
	return r
}

func ToUint16(data []byte) uint16 {
	return binary.BigEndian.Uint16(data[:2])
}

func EatUint16(data []byte) ([]byte, uint16) {
	val := binary.BigEndian.Uint16(data[:2])
	data = data[2:]
	return data, val
}

func TryEatUint16(data []byte) ([]byte, uint16, error) {
	if len(data) < 2 {
		return data, 0, ErrNotEnough
	}
	var val uint16
	data, val = EatUint16(data)
	return data, val, nil
}

func OfUint32(val uint32) []byte {
	r := make([]byte, 4)
	binary.BigEndian.PutUint32(r, val)
	return r
}

func ToUint32(data []byte) uint32 {
	return binary.BigEndian.Uint32(data[:4])
}

func EatUint32(data []byte) ([]byte, uint32) {
	val := binary.BigEndian.Uint32(data[:4])
	data = data[4:]
	return data, val
}

func TryEatUint32(data []byte) ([]byte, uint32, error) {
	if len(data) < 4 {
		return data, 0, ErrNotEnough
	}
	var val uint32
	data, val = EatUint32(data)
	return data, val, nil
}

func OfUint64(val uint64) []byte {
	r := make([]byte, 8)
	binary.BigEndian.PutUint64(r, val)
	return r
}

func ToUint64(data []byte) uint64 {
	return binary.BigEndian.Uint64(data[:8])
}

func EatUint64(data []byte) ([]byte, uint64) {
	val := binary.BigEndian.Uint64(data[:8])
	data = data[8:]
	return data, val
}

func TryEatUint64(data []byte) ([]byte, uint64, error) {
	if len(data) < 8 {
		return data, 0, ErrNotEnough
	}
	var val uint64
	data, val = EatUint64(data)
	return data, val, nil
}

func EatBytes(data []byte, leng int) ([]byte, []byte) {
	val := data[:leng]
	data = data[leng:]
	return data, val
}

func TryEatBytes(data []byte, leng int) ([]byte, []byte, error) {
	if len(data) < leng {
		return data, nil, ErrNotEnough
	}
	var val []byte
	data, val = EatBytes(data, leng)
	return data, val, nil
}

type bytes interface {
	Bytes() []byte
}

func Appends(val []byte, i interface{}) []byte {
	isByteArr := false
	switch i.(type) {
	case *uint8: // *uint8
		val = append(val, OfUint8(*i.(*uint8))...)
	case uint8: // uint8, byte
		val = append(val, OfUint8(i.(uint8))...)
	case *uint16: // *uint16
		val = append(val, OfUint16(*i.(*uint16))...)
	case uint16: // uint16
		val = append(val, OfUint16(i.(uint16))...)
	case *uint32: // *uint32
		val = append(val, OfUint32(*i.(*uint32))...)
	case uint32: // uint32
		val = append(val, OfUint32(i.(uint32))...)
	case *uint64: // *uint64
		val = append(val, OfUint64(*i.(*uint64))...)
	case uint64: // uint64
		val = append(val, OfUint64(i.(uint64))...)
	case []byte: // []byte, []uint8
		isByteArr = true
		val = append(val, i.([]byte)...)
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		if isByteArr {
			break
		}
		for x := 0; x < v.Len(); x++ {
			_v := v.Index(x)
			// handle []bytes firstly
			if _v.Kind() == reflect.Ptr || _v.Kind() == reflect.Interface {
				if _b, ok := _v.Elem().Interface().(bytes); ok { // []bytes
					val = append(val, _b.Bytes()...)
					continue
				}
			}
			switch _v.Kind() {
			case reflect.Uint8: // []uint8, []byte
				val = append(val, OfUint8(uint8(_v.Uint()))...)
			case reflect.Uint16: // []uint16
				val = append(val, OfUint16(uint16(_v.Uint()))...)
			case reflect.Uint32: // []uint32
				val = append(val, OfUint32(uint32(_v.Uint()))...)
			case reflect.Uint64: // []uint64
				val = append(val, OfUint64(uint64(_v.Uint()))...)
			case reflect.Ptr, reflect.Interface:
				b := _v.Elem().Interface()
				switch b.(type) {
				case uint8: // []*uint8
					val = append(val, OfUint8(b.(uint8))...)
				case uint16: // []*uint16
					val = append(val, OfUint16(b.(uint16))...)
				case uint32: // []*uint32
					val = append(val, OfUint32(b.(uint32))...)
				case uint64: // []*uint64
					val = append(val, OfUint64(b.(uint64))...)
				}
			}
		}
	case reflect.Ptr, reflect.Interface:
		if b, ok := v.Elem().Interface().(bytes); ok { // bytes
			val = append(val, b.Bytes()...)
		}
	}
	return val
}
