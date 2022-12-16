package models

type Student struct {
	ID    string `bson:"_id,omitempty"`
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}
