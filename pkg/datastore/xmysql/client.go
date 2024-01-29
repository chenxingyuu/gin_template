package xmysql

import (
	"database/sql"
)

type XDB struct {
	*sql.DB
}
