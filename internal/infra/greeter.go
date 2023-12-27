package infra

import (
	"context"

	"blog/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type GreeterRepo struct {
	db  *PostgresDB
	log *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(db *PostgresDB, logger log.Logger) *GreeterRepo {
	return &GreeterRepo{
		db:  db,
		log: log.NewHelper(logger),
	}
}

func (r *GreeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *GreeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *GreeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *GreeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *GreeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
