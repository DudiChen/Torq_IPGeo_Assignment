package csv

import (
	"Torq_IPGeo_Assignment/models"
	"context"
	"strings"
	"testing"
)

// Table Driven Tests
func TestFindLocation(t *testing.T) {
	var flagtests = []struct {
		in  string
		out models.Location
		err error
	}{
		{"1.1.1.1", models.Location{Country: "United-States", City: "Yavne"}, nil},
		{"0.0.0.0", models.Location{}, models.ErrNotFound},
	}

	csvData := "1.1.1.1,Yavne,United-States\n2.2.2.2,Lod,Germany"
	var store models.Store
	r := strings.NewReader(csvData)
	store, err := NewStore(r)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range flagtests {
		t.Run(tt.in, func(t *testing.T) {
			location, err := store.FindLocation(context.Background(), tt.in)
			if err != tt.err {
				t.Errorf("got %v, want %v", err, tt.err)
			}
			if location != nil && *location != tt.out {
				t.Errorf("got %v, want %v", location, tt.out)
			}
		})
	}
}
