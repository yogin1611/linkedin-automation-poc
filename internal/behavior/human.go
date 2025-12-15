package behavior

import (
	"math/rand"
	"time"
)

type HumanBehavior struct{}

func NewHumanBehavior() Behavior {
	rand.Seed(time.Now().UnixNano())
	return &HumanBehavior{}
}

func (h *HumanBehavior) Think() {
	sleepRandom(800, 2000)
}

func (h *HumanBehavior) ShortPause() {
	sleepRandom(300, 700)
}

func (h *HumanBehavior) LongPause() {
	sleepRandom(2000, 4000)
}

func sleepRandom(minMs, maxMs int) {
	delay := rand.Intn(maxMs-minMs+1) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
