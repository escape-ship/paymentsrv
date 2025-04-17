package mysql

import (
	"database/sql"
	"time"
)

const (
	_defaultConnAttempts = 3
	_defaultConnTimeout  = time.Second
)

type DBConnString string

type mysql struct {
	connAttempts int
	connTimeout  time.Duration

	db *sql.DB
}

var _ DBEngine = (*mysql)(nil)

func New(url DBConnString) (DBEngine, error) {
	mysql := &mysql{
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	var err error
	for mysql.connAttempts > 0 {
		mysql.db, err = sql.Open("postgres", string(url))
		if err == nil {
			break
		}
		time.Sleep(mysql.connTimeout)
		mysql.connAttempts--
	}

	return mysql, nil
}

func (m *mysql) Configure(opts ...Option) DBEngine {
	for _, opt := range opts {
		opt(m)
	}
	return m
}

func (m *mysql) GetDB() *sql.DB {
	return m.db
}

func (m *mysql) Close() {
	if m.db != nil {
		m.db.Close()
	}
}
