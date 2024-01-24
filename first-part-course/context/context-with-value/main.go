package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	// Looks like a storage, localStorage, SessionStorage
	ctx = context.WithValue(ctx, "token", "senha 1283123")
	bookHotel(ctx, "token")
}

func bookHotel(ctx context.Context, name string) {
	token := ctx.Value(name)
	fmt.Println(token)
}
