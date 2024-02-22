package types

type User struct {
	// omitempty: don't show publicly
	ID        string `bson:"_id" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson: "lastName" json:"lastName"`
}
