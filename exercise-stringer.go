package main

import "fmt"

type IPAddr [4]byte

func(x IPAddr) String() string{
	ans:=fmt.Sprintf("%v", x[0])
	for i:=1; i<len(x); i++{
		ans+=fmt.Sprintf(".%v", x[i])
	}
	return ans
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	
}
