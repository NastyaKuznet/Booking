package model

import "time"

const (
	Salt       = "your-salt-here"
	SigningKey = "your-signing-key-here"
	TokenTTL   = 12 * time.Hour
)
