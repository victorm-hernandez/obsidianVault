package weather

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestLocationInfoLog_Log(t *testing.T) {
	logLen := 3
	log := NewLocationInfoLog(logLen, time.Minute*5)

	_, ok := log.Read()

	if ok {
		t.Error("Was able to read value on empty log")
	}

	// Validates it is sequential and rotates the writes and reads
	for i := 0; i < logLen*2; i++ {
		addedInfo := LocationInfo{Name: fmt.Sprintf("LocInfo%v", i), Created: time.Now()}
		log.Log(addedInfo)

		readInfo, ok := log.Read()

		if !ok {
			t.Error("Could not read value from log")
		}

		if readInfo.Name != addedInfo.Name {
			t.Error("Didnt read the expected element on the log")
		}
	}

	// Verify we can continue reading records in sequence since they are still current
	res, ok := log.Read()

	if !ok {
		t.Error("Wasnt able to read data after a full rotation.")
	}

	// Verify the scenario where the writer catches, overwrites data and catches up to the reader multiple times
	for i := 0; i < logLen*3; i++ {
		addedInfo := LocationInfo{Name: fmt.Sprintf("LocInfo%v", i), Created: time.Now()}
		log.Log(addedInfo)
	}

	for i := logLen * 2; i < logLen*3; i++ {
		readInfo, ok := log.Read()

		if !ok {
			t.Error("Could not read value from log")
		}

		if !strings.HasSuffix(readInfo.Name, strconv.Itoa(i)) {
			t.Error("Didnt read the expected element on the log")
		}
	}

	// Verify stale records are never returned
	for i := 0; i < 6; i++ {
		addedInfo := LocationInfo{Name: fmt.Sprintf("LocInfo%v", i), Created: time.Now().Add(time.Hour * (-1))}
		log.Log(addedInfo)
	}

	_, ok = log.Read()

	if ok {
		t.Error("Was able to read values on empty log")
	}

	// Mix of stale and current records
	log.Log(LocationInfo{Name: "LocInfoStale1", Created: time.Now().Add(time.Minute * (-10))})
	log.Log(LocationInfo{Name: "LocInfoStale2", Created: time.Now().Add(time.Minute * (-7))})
	log.Log(LocationInfo{Name: "LocInfoCurrent", Created: time.Now().Add(time.Minute * (-2))})

	res, ok = log.Read()

	if !ok {
		t.Error("Wasn't able to read records")
	}

	if res.Name != "LocInfoCurrent" {
		t.Error("Unexpected record retrieved")
	}
}

func TestLocationInfoLog_Log_Concurrency(t *testing.T) {
	log := NewLocationInfoLog(1000, time.Minute*5)

	// We will simulate populating the log then consuming it
	wg := sync.WaitGroup{}

	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func(workerId int) {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				log.Log(LocationInfo{Name: fmt.Sprintf("Loc%v_%v", workerId, j), Created: time.Now()})
			}
		}(i)
	}

	wg.Wait()

	// Validate that all of the expected values are found, and only found once by concurrent routines
	seenValues := make(map[string]bool)

	for i := 0; i < 1000; i++ {
		res, ok := log.Read()

		if !ok {
			t.Error("Could not read value from the log")
		}

		_, seen := seenValues[res.Name]

		if seen {
			t.Errorf("Duplicated log! Name:%v", res.Name)
		}

		seenValues[res.Name] = true
	}

	// Validate that we have seen all the value we expect
	for i := 0; i < 20; i++ {
		for j := 0; j < 50; j++ {

			_, seen := seenValues[fmt.Sprintf("Loc%v_%v", i, j)]

			if !seen {
				t.Error("Expected value not found")
			}
		}
	}

	// TODO: Validate that we can read/write concurrently without problems
}
