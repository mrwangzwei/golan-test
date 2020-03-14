package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var Conf *ServerConf

type ServerConf struct {
	ConfigPath  string
	ServerName  string       `yaml:"server_name"`
	ServerEnv   string       `yaml:"server_env"`
	WebName     string       `yaml:"web_name"`
	WebListen   string       `yaml:"web_listen"`
	MysqlConfig *mysqlConfig `yaml:"mysql"`
	MongoConfig *mongoConfig `yaml:"mongo"`
}

type mysqlConfig struct {
	MysqlHost     string `yaml:"mysql_host"`
	MysqlPort     int    `yaml:"mysql_port"`
	MysqlUser     string `yaml:"mysql_user"`
	MysqlPwd      string `yaml:"mysql_pwd"`
	MysqlDbname   string `yaml:"mysql_dbname"`
	MysqlPoolOpen int    `yaml:"mysql_pool_open"`
	MysqlPoolIdle int    `yaml:"mysql_pool_idle"`
}

type mongoConfig struct {
	MongoHost   string `yaml:"mongo_host"`
	MongoPort   int    `yaml:"mongo_port"`
	MongoUser   string `yaml:"mongo_user"`
	MongoPwd    string `yaml:"mongo_pwd"`
	MongoDbname string `yaml:"mongo_dbname"`
}

func NewDefaultConfig() *ServerConf {
	c := &ServerConf{}
	return c
}

func (c *ServerConf) LoadConfigFile() error {
	f, err := os.Open(c.ConfigPath)
	defer f.Close()
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(data, c); err != nil {
		return err
	}
	Conf = c
	return nil
}
