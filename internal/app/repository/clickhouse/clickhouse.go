package clickhouse

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Repository struct {
	conn driver.Conn
}

func New(conn driver.Conn) *Repository {
	return &Repository{
		conn: conn,
	}
}