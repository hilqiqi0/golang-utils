package config

type RedisData struct {
	Addr      string `yaml:"Addr"`
	Password  string `yaml:"Password"`
	Pool_size int    `yaml:"Pool_size"`
	Db        int    `yaml:"Db"`
}

func (that *ConfigEngine) GetRedisDataFromConf(name string) *RedisData {
	login := new(RedisData)
	redisLogin := that.GetStruct(name, login)
	return redisLogin.(*RedisData)
}
