package config

type Config struct {
	ConnectionString string
	AcceptDuplicates bool

	FurniLocation string
	IconLocation  string

	WebserverPort int
}
