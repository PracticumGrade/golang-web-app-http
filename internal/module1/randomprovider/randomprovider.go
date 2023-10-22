package randomprovider

import (
	"math/rand"
	"time"
)

type RandomNumberProvider interface {
	GetRandomInt() int
}

type provideRandom struct {
	src rand.Source
}

func NewRandomNumberProvider() RandomNumberProvider {
	return &provideRandom{src: rand.NewSource(time.Now().UnixNano())}
}

func (pr *provideRandom) GetRandomInt() int {
	return int(pr.src.Int63())
}
