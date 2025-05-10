package config

import (
	"github.com/snansidansi/blog-aggregator/internal/database"
)

type State struct {
	Config *Config
	Db     *database.Queries
}
