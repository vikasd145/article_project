package article_admin

import (
	"fmt"

	"github.com/vikasd145/article_project/pkg/redis_cli"

	"github.com/vikasd145/article_project/pkg/Entity"

	"github.com/vikasd145/article_project/internal/config"
)

var (
	GlobalAdmin *AdminData
)

type AdminData struct {
	ArticleDb          *Entity.DB
	ArticleRedisClient *redis_cli.Rcli
}

func AdminInitialize(conf *config.Config) (*AdminData, error) {
	dbcli, err := Entity.InitDB(conf.Ormconfigs.MasterDSN)
	if err != nil {
		fmt.Errorf("Error in initializing db error:%v", err)
		return nil, err
	}
	rediscli, err := redis_cli.NewClient(conf.RedisHost)
	if err != nil {
		fmt.Errorf("Error in initializing redis error:%v", err)
		return nil, err
	}
	adminTemp := &AdminData{
		ArticleDb:          dbcli,
		ArticleRedisClient: rediscli,
	}
	GlobalAdmin = adminTemp
	return adminTemp, nil
}
