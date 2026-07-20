package main

import "time"

type AppConfig struct {
	Addr string
	DB   DBConfig
}

type DBConfig struct {
	URI            string
	Database       string
	ConnectTimeout time.Duration
	ServerTimeout  time.Duration
}
