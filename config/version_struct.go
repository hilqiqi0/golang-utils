package config

/*
有关yaml配置文件的封装,读取算法版本号
*/
type Version struct {
	Algo    string `yaml:"Algorithm"`
	Version string `yaml:"Versions"`
}

func (this *ConfigEngine) GetVersionFromConf(name string) *Version {
	version := new(Version)
	login := this.GetStruct(name, version)
	return login.(*Version)
}
