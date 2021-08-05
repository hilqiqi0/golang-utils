package config

type HttpData struct {
	Url      string   `yaml:"url"`
	Para     []string `yaml:"para"`
	Time_out int      `yaml:"time_out"`
}

func (this *ConfigEngine) GetHttpFromConf(name string) *HttpData {
	login := new(HttpData)
	httpLogin := this.GetStruct(name, login)
	return httpLogin.(*HttpData)
}
