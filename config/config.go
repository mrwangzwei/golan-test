package config

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

var Conf *ServerConf

type ServerConf struct {
	ConfigPath string
	ServerName string `json:"server_name"`
	ServerEnv  string `json:"server_env"`
	WebName    string `json:"web_name"`
	WebListen  string `json:"web_listen"`

	MysqlHost   string `json:"mysql_host"`
	MysqlPort   int `json:"mysql_port"`
	MysqlUser   string `json:"mysql_user"`
	MysqlPwd    string `json:"mysql_pwd"`
	MysqlDbname string `json:"mysql_dbname"`
	MysqlPoolOpen int `json:"mysql_pool_open"`
	MysqlPoolIdle int `json:"mysql_pool_idle"`

	MongoHost   string `json:"mongo_host"`
	MongoPort   int `json:"mongo_port"`
	MongoUser   string `json:"mongo_user"`
	MongoPwd    string `json:"mongo_pwd"`
	MongoDbname string `json:"mongo_dbname"`
}

func InitConfigPath(cmd *cobra.Command) *ServerConf {
	c := &ServerConf{}
	cmd.Flags().StringVar(&c.ConfigPath, "config", "./config/config.json", "path to config file")
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
	if err := json.Unmarshal(data, c); err != nil {
		return err
	}
	Conf = c
	return nil
}
