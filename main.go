package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("Mouse mover started...")

	for {
		// Rastgele bir süre bekle (2-5 saniye arası)
		delay := time.Duration(rand.Intn(3)+2) * time.Second
		time.Sleep(delay)

		// Mevcut fare konumunu al
		x, y := robotgo.GetMousePos()

		// Yeni rastgele bir konum belirle (küçük hareket)
		newX := x + rand.Intn(6) - 3 // -3 ile +3 arasında rastgele hareket
		newY := y + rand.Intn(6) - 3

		// Fareyi yeni konuma taşı
		robotgo.MoveMouse(newX, newY)
		fmt.Printf("Mouse moved to: (%d, %d)\n", newX, newY)
	}
}
