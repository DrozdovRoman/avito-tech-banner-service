package infrastructure

import (
	"context"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/application/common/configuration"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/domain/banner"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/infrastructure/cache"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/infrastructure/db"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/infrastructure/db/postgres"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/infrastructure/db/repository"
	"github.com/DrozdovRoman/avito-tech-banner-service/internal/infrastructure/db/transaction"
	gocache "github.com/patrickmn/go-cache"
	"go.uber.org/fx"
	"log"
	"time"
)

var Module = fx.Options(
	fx.Provide(

		// cache
		func(cfg *configuration.Configuration) cache.Cache {
			return gocache.New(cfg.Cache.Expiration, cfg.Cache.Cleanup)
		},

		// transaction
		func(db db.Client) db.TxManager {
			return transaction.NewTransactionManager(db.DB())
		},

		// db
		func(cfg *configuration.Configuration) (db.Client, error) {
			return postgres.New(context.Background(), cfg)
		},

		// repository
		func(client db.Client) banner.Repository {
			return repository.NewBannerRepository(client)
		},
	),

	fx.Invoke(
		func(lifecycle fx.Lifecycle, client db.Client) {
			lifecycle.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					var err error
					for i := 0; i < 3; i++ {
						if err = client.DB().Ping(ctx); err == nil {
							return nil
						}
						log.Printf("Connection failed: %v. Retrying...", err)
						time.Sleep(time.Duration(i*5) * time.Second)
					}
					return err
				},
				OnStop: func(ctx context.Context) error {
					return client.Close()
				},
			})
		},
	),
)
