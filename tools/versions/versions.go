package versions

import (
	"github.com/hilqiqi0/golang-utils/config"
	"github.com/hilqiqi0/golang-utils/tools/types/strs"
)

type VersionMethod interface {
	GetAlgoVersion(c *config.ConfigEngine, name string)
}

type Version struct {
	Algo    string `yaml:"Algorithm" json:"algo"`
	Version string `yaml:"Versions" json:"version"`
}

func (this *Version) GetAlgoVersion() string {
	return strs.JoinStrs("-", this.Algo, this.Version)
}

func (this *Version) getVersion(v *config.Version) {
	this.Version = v.Version
	this.Algo = v.Algo
}

func (this *Version) GeVersionFromConf(c *config.ConfigEngine, name string) {
	this.getVersion(c.GetVersionFromConf(name))
}
