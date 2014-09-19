package nettools

import (
	"fmt"
	"net"
	"strconv"
)

func BinaryToDottedPort(port string) string {
	return fmt.Sprintf("%d.%d.%d.%d:%d", port[0], port[1], port[2], port[3],
		(uint16(port[4])<<8)|uint16(port[5]))
}

// 97.98.99.100:25958 becames "abcdef" or an empty string if the input is invalid.
func DottedPortToBinary(b string) string {
	var a []uint8
	host, port, _ := net.SplitHostPort(b)
	ip := net.ParseIP(host)
	if ip == nil {
		return ""
	}
	aa, _ := strconv.ParseUint(port, 10, 16)
	c := uint16(aa)
	if ip2 := net.IP.To4(ip); ip2 != nil {
		a = make([]byte, net.IPv4len+2, net.IPv4len+2)
		copy(a, ip2[0:net.IPv4len])
		a[4] = byte(c >> 8)
		a[5] = byte(c)
	} else {
		a = make([]byte, net.IPv6len+2, net.IPv6len+2)
		copy(a[0:], ip[0:net.IPv6len])
		a[16] = byte(c >> 8)
		a[17] = byte(c)
	}
	return (string(a))
}
