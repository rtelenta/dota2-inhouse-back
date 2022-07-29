package domain

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Match struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name" binding:"required"`
}

type Matches []Match

func (c Match) Clone() (clone Match) {
	data, _ := json.Marshal(c)
	json.Unmarshal(data, &clone)
	return
}
