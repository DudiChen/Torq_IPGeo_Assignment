package csv

import (
	"Torq_IPGeo_Assignment/models"
	"context"
	"encoding/csv"
	"io"
)

type Store struct {
	locationsByIp map[string]models.Location
}

func NewStore(r io.Reader) (*Store, error) {
	reader := csv.NewReader(r)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	locationsByIp := make(map[string]models.Location)
	for _, line := range lines {
		if len(line) != 3 {
			continue
		}
		locationsByIp[line[0]] = models.Location{Country: line[2], City: line[1]}
	}
	return &Store{locationsByIp: locationsByIp}, nil
}

// NOTE: In a real-world store implemntation the ctx will be used to support cancelations/timeouts
func (db *Store) FindLocation(ctx context.Context, ip string) (*models.Location, error) {
	location, exists := db.locationsByIp[ip]
	if !exists {
		//return nil, errors.New("location not found")
		return nil, models.ErrNotFound
	}
	return &location, nil
}
