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
	fmt.Println(spec)
}