package settings

type HTTP struct {
	PORT string `env:"PORT" default:"8080"`
}
