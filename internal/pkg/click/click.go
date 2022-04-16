package click

import (
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/pkg/errors"

	"audit/internal/app/config"
)

func Setup(conf config.Clickhouse) (driver.Conn, error){
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: conf.Addr,
		Auth: clickhouse.Auth{
			Database: conf.Auth.DB,
			Username: conf.Auth.User,
			Password: conf.Auth.Password,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": conf.Settings.MaxExecutionTime,
		},
		DialTimeout: time.Duration(conf.DialSecondsTimeout) * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		Debug: conf.Debug,
	})
	return conn, errors.Wrap(err, "connect to clickhouse")
}
