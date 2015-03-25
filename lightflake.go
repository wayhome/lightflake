// Package lightflake is an ID generater, based
// on Timestamp + WorkerId + RandomNumber
package lightflake

import (
	"fmt"
	"math/rand"
	"time"
)

// EPOCH  for lightflake timestamps, starts at the year 2010
var EPOCH = time.Unix(1262275200, 0)

const (
	// Field length 64 =  TimestampLength + WorkerLength + RandomLength

	// WorkerIDBits We'll use 10 bits for the WorkerId
	WorkerIDBits = 10
	// RandomBits We'll use 12 bits for the Random Number
	RandomBits = 12

	// MaxWorkerID worker id mask
	MaxWorkerID = -1 ^ (-1 << 10)
	// MaxRandom ranom number mask
	MaxRandom = -1 ^ (-1 << 12)
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func extraBits(data uint64, shift, length uint) uint64 {
	bitmask := uint64((1<<length - 1) << shift)
	return (data & bitmask) >> shift

}

// Generate Generate a 64 bit, roughly-ordered, globally-unique ID.
func Generate(workerID uint64) (flake uint64, err error) {
	if workerID > MaxWorkerID {
		err = fmt.Errorf("Worker id %v is invalid", workerID)
		return
	}
	milliseconds := uint64(time.Since(EPOCH).Nanoseconds() / 1000000)
	flake = milliseconds<<(WorkerIDBits+RandomBits) +
		workerID<<RandomBits +
		uint64(rand.Intn(MaxRandom))
	return
}

// ParseFlake Parses a lightflake and return a timestamp by milliseconds
func ParseFlake(flake uint64) (timestamp uint64, workerid uint64) {
	timestamp = uint64(EPOCH.Unix()*1000) +
		extraBits(flake, WorkerIDBits+RandomBits, 64-WorkerIDBits-RandomBits)
	workerid = extraBits(flake, RandomBits, WorkerIDBits)
	return
}
