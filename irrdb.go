package irrdb

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func doquery(connect, args string) {
	version := "go-irrdb 0.1"

	CONNECT := connect
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(c, "!!\n")
	fmt.Fprintf(c, fmt.Sprintf("!n %s\n", version))
	message, _ := bufio.NewReader(c).ReadString('C')
	if message == "C" {
		fmt.Println("Got C from server.")
	} else {
		log.Fatal(fmt.Sprintf("Got: %s from server.", message))
	}
	fmt.Fprintf(c, args+"\n")
	req, _ := bufio.NewReader(c).ReadString('C')
	fmt.Println(req)
	return
}

/* Need to check if res = AS-SET or AUT-NUM.
 * AS-SET   : AS-%s
 * AUT-NUM  : AS%d
 */

/* if AS-SET  : !a6+res
   if AUT-NUM : !6as+res */
func Queryv6(s, res string) {
	doquery(s, "!a6"+res)
}

/* if AS-SET  : !a4+res
   if AUT-NUM : !gas+res */

func Queryv4() {
	doquery(s, "!a6"+res)
}
