package gen

import (
	"github.com/fzf-labs/fpkg/db"
	"testing"
)

func TestGenerationPostgres(t *testing.T) {
	client, err := db.NewGormPostgresClient(&db.GormPostgresClientConfig{
		DataSourceName:  "",
		MaxIdleConn:     0,
		MaxOpenConn:     0,
		ConnMaxLifeTime: 0,
		ShowLog:         false,
		Tracing:         false,
	})
	if err != nil {
		return
	}
	Generation(client, defaultMySqlDataMap, "./", "./")
}

func TestGenerationMysql(t *testing.T) {
	client, err := db.NewGormMysqlClient(&db.GormMysqlClientConfig{
		DataSourceName:  "",
		MaxIdleConn:     0,
		MaxOpenConn:     0,
		ConnMaxLifeTime: 0,
		ShowLog:         false,
		Tracing:         false,
	})
	if err != nil {
		return
	}
	Generation(client, defaultPostgresDataMap, "./", "./")
}