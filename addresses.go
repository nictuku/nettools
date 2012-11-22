package nettools

import (
	"fmt"
	"strings"
)

func BinaryToDottedPort(port string) string {
	return fmt.Sprintf("%d.%d.%d.%d:%d", port[0], port[1], port[2], port[3],
		(uint16(port[4])<<8)|uint16(port[5]))
}

var (
	m = make(map[string]byte, 256)
	n = make(map[string]uint16, 65536)
)

func init() {
	var (
		i byte
		j uint16
	)

	for i = 0; i < 255; i++ {
		s := fmt.Sprintf("%d", i)
		m[s] = i
	}
	i++
	s := fmt.Sprintf("%d", i)
	m[s] = i
	for j = 0; j < 65535; j++ {
		s := fmt.Sprintf("%d", j)
		n[s] = j
	}
	j++
	s = fmt.Sprintf("%d", j)
	n[s] = j

}

// 97.98.99.100:25958 becames "abcdef".
func DottedPortToBinary(b string) string {
	a := make([]byte, 6, 6)

	son := [4]string{".", ".", ".", ":"}
	endpos := len(b)
	beginpos := 0
	for i := 0; i < 4; i++ {
		p1 := strings.Index(b[beginpos:endpos], son[i])
		a[i] = m[b[beginpos:(beginpos+p1)]]
		beginpos = beginpos + p1 + 1
	}
	c := n[b[beginpos:]]
	a[4] = uint8(c >> 8)
	a[5] = uint8(c)

	return string(a)
}
