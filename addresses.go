package nettools

import (
	"fmt"
	"strconv"
	"strings"
)

func BinaryToDottedPort(port string) string {
	return fmt.Sprintf("%d.%d.%d.%d:%d", port[0], port[1], port[2], port[3],
		(uint16(port[4])<<8)|uint16(port[5]))
}

// 97.98.99.100:25958 becames "abcdef" or an empty string if the input is invalid.
func DottedPortToBinary(b string) string {
	a := make([]byte, 6, 6)

	son := [4]string{".", ".", ".", ":"}
	endpos := len(b)
	beginPos := 0

	// IP.
	for i := 0; i < len(son); i++ {
		p1 := strings.Index(b[beginPos:endpos], son[i])
		if p1 == -1 {
			return ""
		}
		aa, _ := strconv.ParseUint(b[beginPos:(beginPos+p1)], 10, 8)
		a[i] = byte(aa)
		beginPos = beginPos + p1 + 1
	}

	// Port.
	aa, _ := strconv.ParseUint(b[beginPos:], 10, 16)
	c := uint16(aa)
	a[4] = byte(c >> 8)
	a[5] = byte(c)

	return string(a)
}
