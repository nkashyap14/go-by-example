package hashing

import (
	"testing"
	"fmt"
	"sync"
)

func TestDefaultHashIdempotency(t *testing.T) {
	key := "test"

	hash1 := DefaultHash(key)
	hash2 := DefaultHash(key)

	if hash1 != hash2 {
		t.Errorf("Hash function is not giving the same value for the same key. Got %d and %d for the same testing input", hash1, hash2)
	}
}


func TestNewRing(t *testing.T) {
	ring := NewRing(3, nil)

	if ring.replicationFactor != 3 {
		t.Errorf("Expected replication factor to be 3 got %d", ring.replicationFactor)
	}

	if ring.hashFunc == nil {
		t.Error("Default hash function was not appropriately set")
	}
}

func TestNewRecord(t *testing.T) {
	data := NewRecord("Nikhilesh", "Kashyap", 100000.99, 10000.00, 250000.00, nil)

	hash := DefaultHash("NikhileshKashyap")

	if data.hashPosition != hash {
		t.Errorf("Expected hash to be %d but got %d instead", hash, data.hashPosition)
	}

	if data.FirstName != "Nikhilesh" {
		t.Errorf("Expected first name to be %s but got %s", "Nikhilesh", data.FirstName)
	}

	if data.LastName != "Kashyap" {
		t.Errorf("Expected last name to be %s but got %s", "Kashyap", data.LastName)
	}

	if data.savings != 100000.99 {
		t.Errorf("Expected savings to be %f but got %f", 100000.99, data.savings)
	}

	if data.checkings != 10000.00 {
		t.Errorf("Expected checkings to be %f but got %f", 10000.00, data.checkings)
	}

	if data.four01k != 250000.00 {
		t.Errorf("Expected 401k balance to be %f but got %f", 250000.00, data.four01k)
	}
}

func TestAddDB(t *testing.T) {

	ring := NewRing(3, nil)

	db := &database{serverid: 1}

	err := ring.AddDB(db)
	if err != nil {
		t.Errorf("Failed to add database: %v", err)
	}

	err = ring.AddDB(db)
	if err == nil {
		t.Error("Expect an error while adding the same db again")
	}

	if nodes, exists := ring.mapping[db.serverid]; !exists {
		t.Error("VNodes not created for the db")
	} else if len(nodes) != ring.replicationFactor {
		t.Errorf("Expected %d vnodes but got %d", ring.replicationFactor, len(nodes))
	}
}

func TestRebalance(t *testing.T) {

	ring := NewRing(3, nil)
	dbs := make([]*database, 10)

	for i := range 10 {
		dbs[i] = &database{serverid: i + 1}
	}

	for i := range 10 {
		for j := range 1000 {
			record := NewRecord(fmt.Sprintf("%dNikhilesh", j + i), fmt.Sprintf("Kashyap%d", j * 10), 100000.99, 10000.00, 250000.00, nil)
			dbs[i].records = append(dbs[i].records, record)
		}
	}

	for i := range 10 {
		if len(dbs[i].records) != 1000 {
			t.Errorf("Expected each db to be initialized with records of length %d but this one has %d", 1000, len(dbs[i].records))
		}
		ring.AddDB(dbs[i])
	}


	for i := range 10 {
		if len(dbs[i].records) == 1000 {
			t.Errorf("Expected mappings of records to change but db %d has %d records", i, len(dbs[i].records))
		}
	}

}

func TestConcurrentRecordAddition(t *testing.T) {
    ring := NewRing(3, nil)
    db := &database{serverid: 1}

    var wg sync.WaitGroup
    
    err := ring.AddDB(db)
    if err != nil {
        t.Errorf("Failed to add database: %v", err)
    }

    for i := 0; i < 10000; i++ {
        wg.Add(1)
        record := NewRecord(
            fmt.Sprintf("%dNikhilesh", (i + 10 * (i + 100) - (i * 50))),
            fmt.Sprintf("Kashyap%d", i * 10),
            100000.99,
            10000.00,
            250000.00,
            nil,
        )
        
        go func() {
            defer wg.Done()
            err := ring.addRecord(record)  // Use ring directly from outer scope
            if err != nil {
                t.Errorf("Failed to add record: %v", err)
            }
        }()
    }

    wg.Wait()

    if len(db.records) != 10000 {
        t.Errorf("Expected db to have %d records added concurrently but we have %d", 1000, len(db.records))
    }
}