package main

import (
	"net"
	"fmt"
)

func main(){
	address := "MorningSheep-PC"
	addr, err := net.LookupIP(address)
	if err != nil {
		fmt.Println(err)
	}
	if len(addr) < 1 {
		fmt.Printf("failed to parse IP from address '%v'\n", address)
	}
	for _,ip := range addr{
		fmt.Printf("addr:\t%s\n",ip)
	}
}