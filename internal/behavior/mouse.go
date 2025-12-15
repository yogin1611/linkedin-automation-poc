package behavior

import (
	"fmt"
	"math/rand"
)

type MousePoint struct {
	X int
	Y int
}

func SimulateMouseMovement(from, to MousePoint) {
	steps := rand.Intn(10) + 15
	fmt.Println("[MOUSE] Simulating human-like mouse movement")

	for i := 0; i < steps; i++ {
		fmt.Printf("[MOUSE] Move step %d\n", i+1)
	}
}
