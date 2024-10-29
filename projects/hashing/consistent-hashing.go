package hashing

import (
	"fmt"
	"math"
	"crypto/md5"
)

//defined signature for hashing functions
type HashFunc func(key string) uint32

//default hashing function using md5 hashing which produces a 128 bit hashing value
func DefaultHash(key string) uint32 {
	hasher := md5.New()
	hasher.Write([]byte(key))
	hash := hasher.Sum(nil)
	//md5 returns a 128 bit value. we only return the first four bytes
	return binary.BigEndian.Uint32(hash[:4])
}

//data record
type data struct {
	FirstName string
	LastName string
	Savings string
	Checking string
	Four01k string
}

type vNode struct {
	//serverId the vNode corresponds to
	ServerId int
	//vnode has a record of its hash position
	HashPosition uint32
}

type ring struct {
	//mapping of server id to vnodes
	mapping map[int][]vNode
	databases map[int]*database //tracking of the databases
	//max value 10
	replicationFactor int
	//represents the total size of our hash space
	hashSpace uint32
	hashfunc HashFunc
}

type database struct {
	serverid int
	records []data
}


//I'm assuming that the database is a banking app so for simplicity purposes will
//set it up with first name, last name, savings account balance, checkings balance,
//and 401k.
type data struct {
	first_name, last_name string
	savings, checkings, four01k float64
}

func NewRing(partitions, replicationFactor int, hashFunc HashFunc) *ring {
	if replicationFactor > 10 {
		replicationFactor = 10
	}

	if partitions > 100 {
		partitions = 100
	}

	if hashFunc == nil {
		hashFunc = DefaultHash
	}

	return &ring{
		mapping: make(map[int][]vNode),
		databases: make(map[int]*database),
		replicationFactor: replicationFactor,
		hashSpace: math.MaxUint32,
		hashFunc: hashFunc
	}
}

func (r ring) RemoveDB(db database) bool {
	//removes all vnodes of serverId from the ring
}

//server
func (r ring) AddDb(db database) {
	//takes the db gets a hash
	//mods the hash by perimeter of the ring, multiplies it from 1 -n where n is the replication factor and sets a vnode at each of those points
	// vnode is (n - 1) * hash % perimeter 

}

//internal rebalancing function that gets called by AddDB and RemoveDB to adjust the mappings of data to database
func (r ring) rebalance() {
}

func main() {
}