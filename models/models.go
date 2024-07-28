package models

import "context"

type Store interface {
	FindLocation(ctx context.Context, ip string) (*Location, error)
}

type Location struct {
	Country string `json:"country"`
	City    string `json:"city"`
}
