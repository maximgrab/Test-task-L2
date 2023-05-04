package main

import "flag"

func getFlags()

var (
	ia    []int
	kFlag = flag.Var(&ia, "f", 0, "fields")
	nFlag = flag.String("d", "\t", "delimiter")
	rFlag = flag.Bool("s", true, "separated")
)

func main() {

}
