package types

type User struct {
	// omitempty: don't show publicly
	// both in json and in mongo
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson: "lastName" json:"lastName"`
}
