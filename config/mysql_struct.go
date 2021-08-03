package config

type MysqlDbData struct {
	User       string            `yaml:"user"`
	Password   string            `yaml:"password"`
	Host       string            `yaml:"host"`
	Port       string            `yaml:"port"`
	Db_name    string            `yaml:"db_name"`
	Table_name map[string]string `yaml:"table_name"`
	Max_conns  int               `yaml:"max_conns"`
	Time_out   int               `yaml:"time_out"`
}

func (that *ConfigEngine) GetMySqlFromConf(name string) *MysqlDbData {
	login := new(MysqlDbData)
	mysqlLogin := that.GetStruct(name, login)
	return mysqlLogin.(*MysqlDbData)
}
