package config

type FieldOfView struct {
	Width  int
	Height int
}

type Config struct {
	FieldOfView FieldOfView
	AngleFactor float64
	Port        int
}

func NewConfig() *Config {
	return &Config{
		FieldOfView: FieldOfView{
			Width:  63,
			Height: 47,
		},
		AngleFactor: 1. / 64.,
		Port:        8001,
	}
}
