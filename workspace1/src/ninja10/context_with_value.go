package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "language", "Go")
	foo2(ctx, "language")

	ctx = context.WithValue(ctx, "dog", "Gaston")
	foo2(ctx, "dog")

	foo2(ctx, "color")
}

func foo2(ctx context.Context, s string) {
	if v := ctx.Value(s); v != nil {
		fmt.Println("found value:", v)
		return
	}
	fmt.Println("key not found:", s)
}
