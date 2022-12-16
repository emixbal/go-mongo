package models

const jamaahs = "jamaahs"

type Jamaah struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
}

func JamaahKey() string {
	return jamaahs
}
