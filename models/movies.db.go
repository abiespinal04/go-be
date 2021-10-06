package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Type for DBModel
type DBModel struct {
	DB *mongo.Client
}

// Get returns one movie and error, if any
func (m *DBModel) Get(id int) (*Movie, error) {

	return nil, nil
}

// Returns all movies and error, if any
func (m *DBModel) All(id int) ([]*Movie, error) {
	return nil, nil
}
