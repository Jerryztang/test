package main

import (
	"fmt"
	"hash/adler32"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("请提供两个参数，第一个参数为机器ip，第二个参数为consul端口")
		os.Exit(1)
	}

	CvmIp := os.Args[1]
	Port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("第二个参数必须为整数")
		os.Exit(1)
	}

	PrintID(CvmIp, Port)
}

func PrintID(ip string, port int) {

	fmt.Println("================================")
	OldCode(ip, port)

	NewCode(ip, port)
}

func OldCode(addr string, port int) {
	ipaddr := inet_aton(addr)
	index := uint64(uint64(port)<<32) + uint64(ipaddr)
	fmt.Println("old id:", addr, port, index)
}

func NewCode(addr string, port int) {
	ipaddr := adler32.Checksum([]byte(addr))
	index := uint64(uint64(port)<<32) + uint64(ipaddr)
	fmt.Println("new id:", addr, port, index)
}

func inet_aton(ipaddr string) int64 {
	bits := strings.Split(ipaddr, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}
