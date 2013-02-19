package main

import "net"
import "fmt"
import "log"

type Address struct {
	i int
	j int
	k int
	l int
}

func get_domain(cur_address Address) string {
	address_string := stringlize(cur_address)
	names, err := net.LookupAddr(address_string)
	
	if err != nil {
		log.Print(err)
		return ""
	}
	
	return names[0]
}

func incr(cur_address Address) Address {
	if cur_address.l != 255 {
		cur_address.l += 1
		return cur_address
	}
	
	cur_address.l = 0
	
	if cur_address.k != 255 {
		cur_address.k += 1
		return cur_address
	}
	
	cur_address.k = 0
	
	if cur_address.j != 255 {
		cur_address.j += 1
		return cur_address
	}
	
	cur_address.j = 0
	
	if cur_address.i != 255 {
		cur_address.i += 1
		return cur_address
	}
	
	cur_address.i = 0
	return cur_address
}

func stringlize(cur_address Address) string {
	return fmt.Sprintf("%d.%d.%d.%d", cur_address.i, cur_address.j, cur_address.k, cur_address.l)
}

func print(cur_address Address, domain string) {
	fmt.Printf("%s\t%s\n", stringlize(cur_address), domain)
}

func main() {
	cur_address, zero := Address{}, Address{}
	cur_address = incr(cur_address)
	
	for cur_address != zero {
		domain := get_domain(cur_address)
		cur_address = incr(cur_address)
		print(cur_address, domain)
	}
	
}
