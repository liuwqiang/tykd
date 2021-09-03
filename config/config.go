package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var (
	global   atomic.Value
	globalMu sync.Mutex
	Default  = Config{
		Server: Server{
			"0.0.0.0:8080",
		},
		Jdbc: Jdbc{
			"root",
			"wftest@231",
			"122.115.55.3:32453",
			"gateway",
			60,
			10,
			60,
		},
	}
)

type Config struct {
	Server Server `json:"server"`
	Jdbc   Jdbc   `json:"jdbc"`
}

type Jdbc struct {
	Username        string        `json:"username"`
	Password        string        `json:"password"`
	Address         string        `json:"address"`
	DBName          string        `json:"db_name"`
	MaxOpenConn     int           `json:"max_open_conn"`
	MaxIdleConn     int           `json:"max_idle_conn"`
	ConnMaxLifeTime time.Duration `json:"conn_max_life_time"`
}

type Server struct {
	Addr string `json:"addr"`
}

func init() {
	SetGlobal(Config{})
}

func Global() Config {
	return global.Load().(Config)
}

func SetGlobal(conf Config) {
	globalMu.Lock()
	defer globalMu.Unlock()
	global.Store(conf)
}

func Load(paths []string, conf *Config) error {
	var r io.Reader
	for _, path := range paths {
		f, err := os.Open(path)
		if err == nil {
			r = f
			break
		}
		if os.IsNotExist(err) {
			continue
		}
		return err
	}
	if r == nil {
		return fmt.Errorf("read conf err,conf not found")
	}
	if err := json.NewDecoder(r).Decode(&conf); err != nil {
		return fmt.Errorf("couldn't unmarshal config: %v", err)
	}
	return nil
}
