package behavior

import (
	"fmt"
	"math/rand"
	"time"
)

func SimulateScrolling() {
	steps := rand.Intn(5) + 3
	for i := 0; i < steps; i++ {
		fmt.Println("[SCROLL] Scrolling viewport")
		time.Sleep(time.Duration(rand.Intn(400)+200) * time.Millisecond)
	}
}
