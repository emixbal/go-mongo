package models

const JamaahCollectionName = "jamaahs"

type Jamaah struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	UUID string `json:"uuid" bson:"uuid"`
}
