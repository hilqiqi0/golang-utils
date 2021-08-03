package config

type RedisClusterData struct {
	Addrs       []string `yaml:"Addrs"`
	Master_host []string `yaml:"Master_host"`
	Master_port []string `yaml:"Master_port"`
	Slave_host  []string `yaml:"Slave_host"`
	Slave_port  []string `yaml:"Slave_port"`
	Password    string   `yaml:"Password"`
	Nodes       int      `yaml:"Nodes"`
	Data_time   int      `yaml:"Data_time"`
	Pool_size   int      `yaml:"Pool_size"`
}

func (that *ConfigEngine) GetRedisClusterDataFromConf(name string) *RedisClusterData {
	login := new(RedisClusterData)
	redisLogin := that.GetStruct(name, login)
	return redisLogin.(*RedisClusterData)
}
