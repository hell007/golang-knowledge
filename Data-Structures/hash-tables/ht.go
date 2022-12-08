package hashTables

import (
	"errors"
	list "golang-knowledge/Data-Structures/linked-List"
	"math"
)

type HashTable struct {
	Table    map[int]*list.List
	Size     int
	Capacity int
}

type item struct {
	key   string
	value interface{}
}

func New(cap int) *HashTable {
	table := make(map[int]*list.List, cap)
	return &HashTable{Table: table, Size: 0, Capacity: cap}
}

// Horner's Method to hash string of length L (O(L))
func hashCode(str string) int {
	hash := int32(0)
	for i := 0; i < len(str); i++ {
		hash = int32(hash<<5-hash) + int32(str[i])
		hash &= hash
	}
	return int(math.Abs(float64(hash)))
}

// 散列函数，又称为哈希（Hash函数）
func (ht *HashTable) position(str string) int {
	return hashCode(str) % ht.Capacity
}

func (ht *HashTable) find(index int, key string) (*item, error) {
	l := ht.Table[index]
	var val *item

	l.Each(func(node list.Node) {
		if node.Value.(*item).key == key {
			val = node.Value.(*item)
		}
	})

	if val == nil {
		return nil, errors.New("Not Found")
	}

	return val, nil
}

// 获取
func (ht *HashTable) Get(key string) (interface{}, error) {
	index := ht.position(key)
	item, err := ht.find(index, key)

	if item == nil {
		return "", errors.New("Not Found")
	}

	return item.value, err
}

// 添加
func (ht *HashTable) Append(key, value string) {
	index := ht.position(key)

	if ht.Table[index] == nil {
		ht.Table[index] = list.NewList()
	}

	item := &item{key: key, value: value}

	a, err := ht.find(index, key)
	if err != nil {
		// The key doesn't exist in HashTable
		ht.Table[index].Append(item)
		ht.Size++
	} else {
		// The key exists so we overwrite its value
		a.value = value
	}
}

// 移除
func (ht *HashTable) Remove(key string) error {
	index := ht.position(key)
	l := ht.Table[index]
	var val *item

	l.Each(func(node list.Node) {
		if node.Value.(*item).key == key {
			val = node.Value.(*item)
		}
	})

	if val == nil {
		return nil
	}

	ht.Size--
	return l.Remove(val)
}

// 遍历
func (ht *HashTable) Each(f func(*item)) {
	for k := range ht.Table {
		if ht.Table[k] != nil {
			ht.Table[k].Each(func(node list.Node) {
				f(node.Value.(*item))
			})
		}
	}
}
