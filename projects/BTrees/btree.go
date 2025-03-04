package btree

import (
	"fmt"
	"sort"
)

type tree interface {

	add(T value) boolean
	remove(T value) boolean
	clear()
	contains(T value) boolean
	size() int
	validate() boolean

}

type node[K comparable] struct {
	//slice of keys
	keys []K
	//number of current keys in use and number of current children
	keysSize, childrenSize int
	//slice of pointers to children nodes
	children [] *node[K]
	parent *node[K]
}

type btree[K comparable, V any] struct {
	minSize, minChildSize, maxSize, maxChildSize int
	root node[K, V]
}


func NewBTree(order int) {
	return btree {
		minSize: order,
		minChildSize: order + 1,
		maxSize: 2 * order,
		maxChildSize: (2 * order) + 1,
		root: nil
	}
}

func newNode[K comparable](par *node[K], maxKeySize, maxChildSize int) *node[K] {
	return &node[K] {
		parent: par,
		keys: make([]K, maxKeySize + 1),
		keysSize: 0,
		children: make([]*node[K], maxChildSize + 1),
		childrenSize: 0 
	}
}

func (n node[K]) getKey(index int) {
	return node.keys[index]
}

func (n node[K]) indexOf(T value) {
	for i := 0; i < n.keysSize; i++ {
		if (n.keys[i] == value) {
			return i
		}
	}

	return -1
}

func (n *node[K]) compareKeys(a, b K) int {
    // Return -1 if a < b, 0 if a == b, 1 if a > b
    if a < b {
        return -1
    } else if a > b {
        return 1
    }
    return 0
}

func (n *node[K]) addKey(T value) {
	n.keys[n.keysSize + 1] = value
	n.keysSize++

	sort.slice(n.keys[:n.keysSize], func(i, j int) bool {
		return n.compareKeys(n.keys[i], n.keys[j]) < 0
	})

}