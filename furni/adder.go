package furni

// Adder is responsible for adding furniture to the database.
type Adder interface {
	Add(furni Furni) error
}
