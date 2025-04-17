package mysql

import "time"

type Option func(*mysql)

func WithConnAttempts(attempts int) Option {
	return func(m *mysql) {
		m.connAttempts = attempts
	}
}

func WithConnTimeout(timeout time.Duration) Option {
	return func(m *mysql) {
		m.connTimeout = timeout
	}
}
