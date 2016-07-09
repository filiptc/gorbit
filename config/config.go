package config

type fieldOfView struct {
	Width  int
	Height int
}

type cameraConfig struct {
	Width  uint32
	Height uint32
	Device string
}

type Config struct {
	FieldOfView fieldOfView
	AngleFactor float64
	Port        uint64
	Cam         cameraConfig
	Font        []byte
	Index       []byte
}

func NewConfig() *Config {
	return &Config{
		FieldOfView: fieldOfView{
			Width:  63,
			Height: 47,
		},
		AngleFactor: 1. / 64.,
		Port:        8001,
	}
}
