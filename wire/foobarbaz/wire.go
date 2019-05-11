//+build wireinject

package foobarbaz


import (
	"context"
	"github.com/google/wire"
)


func InitializeBaz(ctx context.Context) (LBaz, error) {
	wire.Build(SuperSet)
	return LBaz{}, nil
}

func InitializeBaz1(ctx context.Context) (Baz, error) {

	wire.Build(SuperSet)
	return Baz{}, nil
}
