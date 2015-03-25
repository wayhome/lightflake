// Package lightflake is an ID generater, based
// on Timestamp + WorkerId + RandomNumber
package lightflake

import (
	"math/rand"
	"time"
)

// EPOCH  for lightflake timestamps, starts at the year 2010
var EPOCH = time.Unix(1262275200, 0)

const (
	// Field length 64 =  TimestampLength + WorkerLength + RandomLength

	// TimestampLength We'll use 41 bits for the timestamp
	TimestampLength = 41
	// WorkerLength  We'll use 8 bits for the WorkerId
	WorkerLength = 8
	// RandomLength leave 15 bits for the random data
	RandomLength = 15
	// RandomMAX max value we can store in random data
	RandomMAX = 0x7fff

	// TimestampShift left shift amounts for timestamp
	TimestampShift = 64 - TimestampLength
	// WorkerShift left shift amounts for the random data
	WorkerShift = RandomLength
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func extraBits(data uint64, shift, length uint) uint64 {
	bitmask := uint64((1<<length - 1) << shift)
	return (data & bitmask) >> shift

}

// Generate Generate a 64 bit, roughly-ordered, globally-unique ID.
func Generate(workerID uint64) uint64 {
	milliseconds := uint64(time.Since(EPOCH).Nanoseconds() / 1000000)
	flake := milliseconds<<TimestampShift +
		workerID<<WorkerShift +
		uint64(rand.Intn(RandomMAX))
	return flake
}

// ParseFlake Parses a lightflake and return a timestamp by milliseconds
func ParseFlake(flake uint64) (timestamp uint64, workerid uint64) {
	timestamp = uint64(EPOCH.Unix()*1000) + extraBits(flake, TimestampShift, TimestampLength)
	workerid = extraBits(flake, WorkerShift, WorkerLength)
	return
}
