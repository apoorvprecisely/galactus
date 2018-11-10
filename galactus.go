package galactus

import "github.com/jinzhu/configor"

type Galactus struct {
	Config Config
}
type Config struct {
	APPName string `default:"galactus"`
	Reader  struct {
		Uri   string `default:"localhost"`
		Topic string `default:"galactus"`
	}
}

func New() (Galactus, error) {
	var defaultC = Config{}
	configor.Load(&defaultC, "config.yml")
	return Galactus{Config: defaultC}, nil
}
