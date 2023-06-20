package topk

import "math/rand"

const LOOKUP_TABLE = 256

type bucket struct {
	fingerprint uint32
	count       uint32
}

type HeavyKeeper struct {
	k           uint32
	width       uint32
	depth       uint32
	decay       float64
	lookupTable []float64
	r           *rand.Rand
	buckets     [][]bucket
	minHeap
}
