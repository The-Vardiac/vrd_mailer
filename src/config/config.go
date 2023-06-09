package config

type Config struct{}

func (cfg *Config) InitRabbitmq() {
	// make connection
	var rabbitmqConf RabbitmqConf
	rabbitmqConf.RabbitmqMakeConn()
}