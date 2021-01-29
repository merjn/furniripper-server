package config

type Config struct {
	FurniLocation string
	IconLocation  string

	DbName string
	DbUser string
	DbPass string
	DbPort int

	WebserverPort int
}
