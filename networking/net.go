package networking

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

var (
	unicast = Atoi("255.255.255.255")
)

func IsIPv4(ip string) bool {
	if ip == "" {
		return false
	}
	_ip := net.ParseIP(ip)
	return _ip != nil && _ip.To4() != nil
}

func GetVeryFirstIPv4(ip, mask string) string {
	return Itoa(Atoi(ip) & Atoi(mask))
}

func GetFirstIPv4(ip, mask string) string {
	return Itoa(Atoi(ip)&Atoi(mask) + 1)
}

func isPowerOf2(i uint32) bool {
	m := math.Log2(float64(i))
	return m == float64(int(m))
}

func IsIPv4Netmask(mask string) bool {
	if mask == "" {
		return false
	}
	m := unicast ^ Atoi(mask)
	return m == 0 || isPowerOf2(m+1)
}

func GetIPv4NetmaskBits(mask string) int {
	m := unicast ^ Atoi(mask)
	return 32 - int(math.Log2(float64(m+1)))
}

func ParseBroadcast(ip, netmask string) string {
	return Itoa(unicast ^ Atoi(netmask) | Atoi(ip))
}

func GetBroadcast(ip string, maskBits uint) string {
	bits := 32 - maskBits
	return Itoa(((unicast >> bits) << bits) ^ unicast | Atoi(ip))
}

// Atoi converts string ip to uint32
func Atoi(ip string) uint32 {
	bits := strings.Split(ip, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	sum := uint32(b0<<24 + b1<<16 + b2<<8 + b3)

	return sum
}

func AtoBytes(ip string) []byte {
	return ItoBytes(Atoi(ip))
}

func AatoBytes(ips []string) []byte {
	rs := make([]byte, len(ips)*4)
	for i := range ips {
		copy(rs[i*4:i*4+4], AtoBytes(ips[i]))
	}
	return rs
}

// ItoBytes converts uint32 ip to byte array
func ItoBytes(ip uint32) []byte {
	bits := []byte{0, 0, 0, 0}
	bits[0] = byte(ip >> 24)
	bits[1] = byte(ip >> 16)
	bits[2] = byte(ip >> 8)
	bits[3] = byte(ip)
	return bits
}

func Itoa(ip uint32) string {
	bits := ItoBytes(ip)
	return fmt.Sprintf("%d.%d.%d.%d", bits[0], bits[1], bits[2], bits[3])
}

func BytesToA(ip []byte) string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func IsTCPAddrOccupied(addr string) bool {
	laddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return true
	}
	lis, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return true
	}
	lis.Close()
	return false
}

func IsSameSubnet(ip0, ip1, mask string) bool {
	_mask := Atoi(mask)
	return Atoi(ip0)&_mask == Atoi(ip1)&_mask
}

func IsIPGreaterEqual(ip0, ip1 string) bool {
	return Atoi(ip0) >= Atoi(ip1)
}
