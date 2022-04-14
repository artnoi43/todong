package store

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/artnoi43/todong/config"
	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/domain/usecase"
	"github.com/artnoi43/todong/domain/usecase/redisclient"
	"github.com/artnoi43/todong/lib/enums"
	"github.com/artnoi43/todong/lib/postgres"
)

func Init(conf *config.Config) usecase.DataGateway {
	switch enums.StoreType(strings.ToUpper(string(conf.Store))) {
	case enums.Gorm:
		pg, err := postgres.New(&conf.Postgres)
		if err != nil {
			log.Fatalln("error: failed to open database:", err.Error())
		}
		if err := pg.AutoMigrate(&model.Todo{}, &model.User{}); err != nil {
			panic(fmt.Sprintf("error: failed to auto-migrate *gorm.DB: %v", err.Error()))
		}
		gormStore := NewGormStore(pg)
		return NewGormDataGateway(gormStore)
	case enums.Redis:
		ctx := context.Background()
		rd, err := redisclient.New(ctx, &conf.Redis)
		if err != nil {
			log.Fatalln("error: failed to open database:", err.Error())
		}
		return NewRedisDataGateway(rd)
	}
	panic(fmt.Sprintf("invalid store config: %s", conf.Store))
}
