package hash_table

import "fmt"

const sizeTable = 20

type HashTable struct {
	table [sizeTable][][2]string
}

func (ht *HashTable) Add(key string, value string) {
	var tableIndex, hashValue int32
	var currPair [2]string

	hashValue = Hash(key)
	tableIndex = hashValue%sizeTable - 1

	currPair[0], currPair[1] = key, value

	if len(ht.table[tableIndex]) == 0 {
		ht.table[tableIndex] = append(ht.table[tableIndex], currPair)
	} else {
		for i := range ht.table[tableIndex] {
			if Hash(ht.table[tableIndex][i][0]) == hashValue {

				ht.table[tableIndex][i] = currPair
				return
			}
		}
		ht.table[tableIndex] = append(ht.table[tableIndex], currPair)
	}
}

func (ht *HashTable) Get(key string) (string, error) {
	var tableIndex, hashValue int32

	hashValue = Hash(key)
	tableIndex = hashValue%sizeTable - 1

	if len(ht.table[tableIndex]) == 0 {
		return ht.table[tableIndex][0][1], nil
	} else {
		for i := range ht.table[tableIndex] {
			if Hash(ht.table[tableIndex][i][0]) == hashValue {
				return ht.table[tableIndex][i][1], nil
			}
		}
	}

	return "", fmt.Errorf("no such element")
}

func (ht *HashTable) Delete(key string) {
	var tableIndex, hashValue int32

	hashValue = Hash(key)
	tableIndex = hashValue%sizeTable - 1

	for i := range ht.table[tableIndex] {
		if ht.table[tableIndex][i][0] == key {
			if i == len(ht.table[tableIndex][i])-1 {
				ht.table[tableIndex] = ht.table[tableIndex][:i]
			} else {
				ht.table[tableIndex] = append(ht.table[tableIndex][:i], ht.table[tableIndex][i+1:]...)
			}
			return
		}
	}

	panic("No such element")
}
