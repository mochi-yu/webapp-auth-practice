package infrastructure

import (
	"context"

	"github.com/mochi-yu/webapp-auth-practice/config"
	"github.com/mochi-yu/webapp-auth-practice/ent"

	_ "github.com/lib/pq"
)

func NewDB(cfg *config.Config) (*ent.Client, error) {
	client, err := ent.Open("postgres", cfg.DBDataSource)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
