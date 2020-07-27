package irrdb

import (
        "bufio"
        "fmt"
        "net"
        "log"
)

/*
DEBUG: bgpq_expander.c:892 bgpq_expand Acquired sendbuf of 4608 bytes
DEBUG: bgpq_expander.c:914 bgpq_expand Sending '!!' to server to request for theconnection to remain open
DEBUG: bgpq_expander.c:923 bgpq_expand b->identify: Sending '!n bgpq4 0.0.6' to server.
DEBUG: bgpq_expander.c:935 bgpq_expand Got answer C
DEBUG: bgpq_expander.c:391 bgpq_pipeline expander: sending !6as49271
2001:678:dcc::/48n2001:678:dd0::/48n2a0e:8f02:2020::/44n2a0e:8f02:2020::/48n
*/

func doquery(connect, args string) {
        version := "go-irrdb 0.1"

        CONNECT :=connect
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
        fmt.Println("For loop")
        fmt.Fprintf(c, args+"\n")
        req, _ := bufio.NewReader(c).ReadString('C')
        fmt.Println(messagg)
        return
}

/* Need to check if res = AS-SET or AUT-NUM.
 * AS-SET   : AS-%s
 * AUT-NUM  : AS%d
 */

func Queryv6(s,res string) {
  /* if AS-SET  : !a6+res
     if AUT-NUM : !6as+res */
  doquery(s, "!a6"+res)
}

func Queryv4() {
  /* if AS-SET  : !a4+res
     if AUT-NUM : !gas+res */

	fmt.Println("getv6!")
}
