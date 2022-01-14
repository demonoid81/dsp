package app

const DefaultEnvFile string = ".env"
const DefaultCrypto string = "5368316e676520742879732370617371"
const DefaultClickURL string = "https://api.labyrinthads.com"
const DefaultMediaURL string = "http://i.labyrinthads.com"
const DefaultRedirectURL string = "https://labyrinthads.com"
const DefaultPGHost string = "162.55.244.120"
const DefaultPGPort int = 5432
const DefaultPGUser string = "postgres"
const DefaultPGPassword string = "9mcP7ZxEuMNcFU"
const DefaultPGDatabase string = "dsp"
const DefaultRevshare float64 = 0.00
const DefaultMongoURL string = "mongodb://dspadmin:2XqUnHQbjGhw2pM68TkgGKe@162.55.244.120:27017"
const DefaultMongoDatabase string = "dsp"
const DefaultRedisHost string = "127.0.0.1"
const DefaultRedisPort int = 6379

type PG struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Config struct {
	Crypto             string
	ClickURL           string
	MediaURL           string
	RedirectURL        string
	Revshare           float64

	PG                 *PG

	MongoURL           string
	MongoDatabase      string

	RedisHost string
	RedisPort int

	CorsAllowedMethods []string
	CorsAllowedOrigins []string
	CorsAllowedHeaders []string
}
