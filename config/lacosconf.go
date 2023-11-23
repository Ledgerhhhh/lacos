package config

var (
	RCofnig = new(RemoteConfig)
)

type RemoteConfig struct {
	Discovery Discovery
}

type Discovery struct {
	Host       string
	Port       int
	ConfigName string
}
