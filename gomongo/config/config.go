package config

type Conf struct {
	MongoConf *MongConf_
}

type MongConf_ struct {
	Hosts       []string //["192.168.34.105:27017"]
	MaxPoolSize uint64
	Username    string
	Password    string
}

var conf *Conf

func init() {
	conf = &Conf{MongoConf: &MongConf_{}}
	conf.MongoConf.Hosts = []string{"192.168.34.105:27017"}
	conf.MongoConf.MaxPoolSize = 10
	conf.MongoConf.Username = ""
	conf.MongoConf.Password = ""
}

func GetConf() *Conf {
	return conf
}
