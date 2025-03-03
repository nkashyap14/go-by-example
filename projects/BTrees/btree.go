package btree

import (
	"fmt"
)

type tree interface {

	add(T value) boolean
	remove(T value) boolean
	clear()
	contains(T value) boolean
	size() int
	validate() boolean

}

type btree struct {
	minSize, minChildSize, maxSize, maxChildSize int  
}


func NewBTree(order int) {
	return btree {
		minSize: order,
		minChildSize: order + 1,
		maxSize: 2 * order,
		maxChildSize: (2 * order) + 1
	}
}