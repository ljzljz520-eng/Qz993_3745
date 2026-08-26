package config

import "os"

type Config struct { DataPath string; Address string; ReadOnly bool; MaxResults int }
func Default() Config { return Config{DataPath:"artstock.db",Address:":8080",MaxResults:100} }
func FromEnv() Config { c:=Default(); if v:=os.Getenv("ARTSTOCK_DB"); v!="" { c.DataPath=v }; if v:=os.Getenv("ARTSTOCK_ADDR"); v!="" { c.Address=v }; return c }
func (c Config) Limit(n int) int { if n<1 { return 1 }; if n>c.MaxResults { return c.MaxResults }; return n }
func EnsurePath(c Config) error { if c.DataPath=="" { return os.ErrInvalid }; return nil }
