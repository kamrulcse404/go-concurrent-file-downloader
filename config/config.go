package config

import "time"

const (
	WorkerCount = 5
	DataDir     = "downloads"
	TimeFormat  = "20060102_150405"
)

const HTTPTimeout = 30 * time.Second