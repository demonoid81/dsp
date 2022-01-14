package app

type Env struct {
	Cfg   *Config
	Mongo *Mongo
	Redis *Redis
}
