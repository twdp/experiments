package main

import (
	"context"
	"experiments/wire/foobarbaz"
	"fmt"
)

func main()  {
	f, err := foobarbaz.InitializeBaz(context.Background())
	fmt.Println(f, err)
}
