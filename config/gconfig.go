package config

var (
	GConfig = new(GlobalConfig)
)

type GlobalConfig struct {
	ServiceMapping ServiceMapping
	GateWayConfig  GateWayConfig
	SystemConfig   SystemConfig
	RedisConfig    RedisConfig
	NsqdConfig     NsqdConfig
}

type ServiceMapping struct {
	HomeApi string
}

type GateWayConfig struct {
	LimitH int
	LimitM int
}
type SystemConfig struct {
	Serverip string
	Port     int
}

type RedisConfig struct {
	IP       string
	Port     int
	Network  string
	Password string
	DB       int
}

type NsqdConfig struct {
	Host     string
	Port     int
	Topic    string
	Channel  string
	Channel2 string
}
