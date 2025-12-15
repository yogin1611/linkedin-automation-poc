package behavior

import (
	"fmt"
	"math/rand"
	"time"
)

func SimulateTyping(text string) {
	fmt.Print("[TYPE] ")
	for _, ch := range text {
		fmt.Printf("%c", ch)
		time.Sleep(time.Duration(80+rand.Intn(120)) * time.Millisecond)
	}
	fmt.Println()
}
