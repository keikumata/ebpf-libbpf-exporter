package main

import (
	"fmt"

	"github.com/cilium/ebpf"
)

func main() {
	file := "test.elf"
	spec, err := ebpf.LoadCollectionSpec(file)
	if err != nil {
		panic(err)
	}
	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		panic(err)
	}
	defer coll.Close()
	fmt.Println(coll.Programs)
}