package config

import gohbase "github.com/tsuna/gohbase"

type HbaseDbData struct {
	Thrift    string `yaml:"thrift"`
	Zkquorum  string `yaml:"zkquorum"`
	Option    string `yaml:"option"`
	Namespace string `yaml:"namespace"`
	Account   string `yaml:"account"`
}

func (that *ConfigEngine) GetHbaseFromConf(name string) *HbaseDbData {
	login := new(HbaseDbData)
	hbaseLogin := that.GetStruct(name, login)
	return hbaseLogin.(*HbaseDbData)
}

type HBaseDbInfo struct {
	Zkquorum  string
	Option    string
	Namespace string
	Client    gohbase.Client
}

type FunctionA interface {
	TransferA() map[string]string
}

type FunctionB interface {
	TransferB() map[string]string
}

type Sample struct {
	Name     string   `json:"name"`
	SectionA SectionA `json:"section_a"`
	SectionB SectionB `json:"section_b"`
}

type SectionA struct {
	FieldA string `json:"field_a"`
}

type SectionB struct {
	FieldB string `json:"field_b"`
}

func (s *Sample) TransferA() map[string]string {
	temp := make(map[string]string)
	temp[s.SectionA.FieldA] = s.SectionA.FieldA
	return temp
}

func (s *Sample) TransferB() map[string]string {
	temp := make(map[string]string)
	temp[s.SectionB.FieldB] = s.SectionB.FieldB
	return temp
}

func Transfer(s FunctionB) {
	if c, ok := s.(FunctionA); ok {
		c.TransferA()
	}
}
