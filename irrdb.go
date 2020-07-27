package irrdb

import (
	"bufio"
	"fmt"
	"log"
  "regexp"
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

func getquery(res string) string{
/* if AS-SET  : !a6+res
   if AUT-NUM : !6as+res */

  var autnum = regexp.MustCompile(`^\d{1,10}$`)
  var asset = regexp.MustCompile(`AS-[A-Za-z0-9]+`)

  if (autnum.MatchString(res)) {
    return fmt.Sprintf("!6as%s", res)
  }

  if (asset.MatchString(res)) {
    return fmt.Sprintf("!a6%s", res)
  }
  return "0"
}

func Queryv6(s, res string) {
  doquery(s, getquery(res))
}

/* if AS-SET  : !a4+res
   if AUT-NUM : !gas+res */

func Queryv4(s, res string) {
	doquery(s, "!a6"+res)
}
