package config

type Config struct {
	// FurniLocation is the directory where all swf files are.
	FurniLocation string

	DbName string
	DbUser string
	DbPass string
	DbPort int

	WebserverPort int
}
