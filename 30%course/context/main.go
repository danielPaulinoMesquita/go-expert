package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*6)
	defer cancel()
	bookHotel(ctx)

}

func bookHotel(ctx context.Context) {
	// case is working as listener

	select {
	case <-ctx.Done():
		fmt.Println("HOTEL Booking cancelled. TimeOut")

	case <-time.After(5 * time.Second):
		fmt.Println("HOTEL BOOKED")
	}
}
