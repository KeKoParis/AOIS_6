package main

import (
	"fmt"
	ht "lab6_1/hash_table"
)

func main() {
	var table ht.HashTable

	table.Add("hedgehog", "indian")
	fmt.Println(table.Get("hedgehog"))
	table.Add("hedgehog", "somali")
	fmt.Println(table.Get("hedgehog"))

	table.Add("primate", "gibbons")
	table.Add("bear", "panda")
	fmt.Printf("Hash \"bear\": %d	index: %d \n", ht.Hash("bear"), ht.Hash("bear")%20)
	fmt.Printf("Hash \"primate\": %d	index: %d \n", ht.Hash("primate"), ht.Hash("bear")%20)
	fmt.Println(table.Get("bear"))
	fmt.Println(table.Get("primate"))

	table.Delete("primate")
	_, err := table.Get("primate")
	if err != nil {
		panic(err)
	}
}
