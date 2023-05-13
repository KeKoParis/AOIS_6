package main

import (
	"fmt"
	ht "lab6_1/hash_table"
)

func main() {
	var table ht.HashTable
	fmt.Println("Get hedgehog")

	table.Add("hedgehog", "indian")
	fmt.Println(table.Get("hedgehog"))
	table.Add("hedgehog", "somali")
	fmt.Println(table.Get("hedgehog"))

	table.Add("primate", "gibbons")
	table.Add("bear", "panda")
	fmt.Printf("Hash \"bear\": %d \tindex: %d \n", ht.Hash("bear"), ht.Hash("bear")%20)
	fmt.Printf("Hash \"primate\": %d \tindex: %d \n", ht.Hash("primate"), ht.Hash("bear")%20)

	fmt.Println("\nGet bear, primate")
	fmt.Println(table.Get("bear"))
	fmt.Println(table.Get("primate"))

	table.Delete("primate")
	fmt.Println("get primate after removal:", table.Get("primate"), ".")

	fmt.Printf("\nHash \"beaver\": %d \tindex: %d \n", ht.Hash("beaver"), ht.Hash("beaver")%20)
	fmt.Printf("Hash \"cat\": %d \tindex: %d \n", ht.Hash("cat"), ht.Hash("cat")%20)

	table.Add("cat", "persian")
	table.Add("beaver", "north american")

	fmt.Println("\nGet beaver, cat")

	fmt.Println(table.Get("cat"))
	fmt.Println(table.Get("beaver"))
}
