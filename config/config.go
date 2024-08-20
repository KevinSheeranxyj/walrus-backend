package config

type config struct {
	System system `yaml:"system"`
}

type system struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var Config *config

func Init() {
	Config = &config{}
}
