package hashing

import (
	"fmt"
	"math"
	"crypto/md5"
	"sort"
	"encoding/binary"
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

func generateVNodeKey(serverId, index int) string {
	return fmt.Sprintf("%d-%d", serverId, index)
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
	hashFunc HashFunc
}

type database struct {
	serverid int
	records []data
}


//I'm assuming that the database is a banking app so for simplicity purposes will
//set it up with first name, last name, savings account balance, checkings balance,
//and 401k.
type data struct {
	FirstName, LastName string
	savings, checkings, four01k float64
	hashPosition uint32
}

// Function to create a new record with pre-calculated hash
func NewRecord(firstName, lastName string, savings, checkings, four01k float64, hashFunc HashFunc) data {
    key := firstName + lastName
    return data{
        FirstName:    firstName,
        LastName:     lastName,
        savings:      savings,
        checkings:    checkings,
        four01k:      four01k,
        hashPosition: hashFunc(key),
    }
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
		hashFunc: hashFunc,
	}
}

func (r ring) RemoveDB(db *database) error {
	//removes all vnodes of serverId from the ring

	if _, exists := r.databases[db.serverid]; !exists {
		return fmt.Errorf("Database with Server ID %d does not exist", db.serverid)
	}

	delete(r.mapping, db.serverid)

	delete(r.databases, db.serverid)

	r.rebalance()
	
	//no error occured so we return nil
	return nil
}

//server
func (r *ring) AddDB(db *database) error {

	//when we look up a key in a map it returns two values, the value and a boolean indicating whether it is found or not
	//we are taking the booleanin this case and setting it to exists and initiating the block off the value of exists
	if _, exists := r.databases[db.serverid]; exists {
		return fmt.Errorf("Database with Server ID %d already exists", db.serverid)
	}

	//Store the database in the database map
	r.databases[db.serverid] = db

	//create the vnodes for the database
	vNodes := make([]vNode, 0, r.replicationFactor)

	for i := 0; i < r.replicationFactor; i++ {
		key := generateVNodeKey(db.serverid, i)
		hash := r.hashFunc(key)

		vNodes = append(vNodes, vNode{
			ServerId : db.serverid, 
			HashPosition: hash,
		})
	}

	r.mapping[db.serverid] = vNodes
	
	r.rebalance()

	//no error occurred so we return nil
	return nil

}

//internal rebalancing function that gets called by AddDB and RemoveDB to adjust the mappings of data to database
func (r *ring) rebalance() error {

	if len(r.databases) <= 1 {
		return nil
	}

	allVNodes := make([]vNode, 0)

	//collect alll the vnodes
	for _, nodes := range r.mapping {
		allVNodes = append(allVNodes, nodes...)
	}


	//sort the vnodes by their hash position on the ring in order to determine what records get moved
	sort.Slice(allVNodes, func(i, j int) bool {
		return allVNodes[i].HashPosition < allVNodes[j].HashPosition
	})

	for i, vnode := range allVNodes {
		prevIndex := (i - 1 + len(allVNodes)) % len(allVNodes)
		startPos := allVNodes[prevIndex].HashPosition
		endPos := vnode.HashPosition

		//get db responsible for this range
		responsibleDB := r.databases[vnode.ServerId]

		for dbID, db := range r.databases {
			if dbID == vnode.ServerId {
				continue //can skip responsiblie db
			}

			//get the records that have to be moved
			recordsToMove := make([]data, 0)
			records := make([]data, 0)

			for _, record := range db.records {
				hashPos := record.hashPosition

				belongsInRange := false

				if startPos < endPos {
					belongsInRange = hashPos > startPos && hashPos <= endPos
				} else { //in the case that the vnodes are wrapping around
					belongsInRange = hashPos > startPos || hashPos <= endPos
				}

				if belongsInRange {
					recordsToMove = append(recordsToMove, record)
				} else {
					records = append(records, record)
				}
			}

			if len(recordsToMove) > 0 {
				db.records = records
				responsibleDB.records = append(responsibleDB.records, recordsToMove...)

			}
		}
	}

	return nil
}

func main() {
}