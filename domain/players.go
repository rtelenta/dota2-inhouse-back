package domain

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	ID           *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Alias        string              `json:"alias,omitempty" bson:"alias,omitempty"`
	SteamId      string              `json:"steamid" bson:"steam_id,omitempty"`
	PersonaName  string              `json:"personaname" bson:"persona_name,omitempty"`
	ProfileUrl   string              `json:"profileurl" bson:"profile_url,omitempty"`
	Avatar       string              `json:"avatar" bson:"avatar,omitempty"`
	AvatarMedium string              `json:"avatarmedium" bson:"avatar_medium,omitempty"`
	AvatarFull   string              `json:"avatarfull" bson:"avatar_full,omitempty"`
	Active       bool                `json:"-" bson:"active,omitempty"`
}

type Players []Player

func (c Player) Clone() (clone Player) {
	data, _ := json.Marshal(c)
	json.Unmarshal(data, &clone)
	return
}
