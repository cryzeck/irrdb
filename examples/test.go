package main

import "github.com/cryzeck/irrdb"

func main() {
  irrdb.Queryv6("rr.ntt.net:43", "AS-BAKE")
  irrdb.Queryv6("rr.ntt.net:43", "49271")
}
