package gen

import (
	"testing"

	"github.com/fzf-labs/fpkg/db"
	"github.com/stretchr/testify/assert"
)

func TestGenerationPostgres(t *testing.T) {
	client, err := db.NewGormPostgresClient(&db.GormPostgresClientConfig{
		DataSourceName:  "host=0.0.0.0 port=5432 user=postgres password=123456 dbname=user sslmode=disable TimeZone=Asia/Shanghai",
		MaxIdleConn:     0,
		MaxOpenConn:     0,
		ConnMaxLifeTime: 0,
		ShowLog:         false,
		Tracing:         false,
	})
	if err != nil {
		return
	}
	Generation(client, DefaultMySQLDataMap, "./example/postgres/")
	assert.Equal(t, nil, err)
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
	Generation(client, DefaultPostgresDataMap, "./example/postgres/")
	assert.Equal(t, nil, err)
}
