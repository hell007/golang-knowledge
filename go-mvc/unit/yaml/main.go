package main

import (
	"fmt"
	"github.com/kataras/golog"
	//"gopkg.in/yaml.v2"
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

// 方式一
var (
	DBConfig DB
)

type DB struct {
	Master DBConfigInfo
	Slave  DBConfigInfo
}

type DBConfigInfo struct {
	Dialect      string `yaml:"dialect"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	Database     string `yaml:"database"`
	Charset      string `yaml:"charset"`
	ShowSql      bool   `yaml:"showSql"`
	LogLevel     string `yaml:"logLevel"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
}

// 方式二
type conf struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	Dbname string `yaml:"dbname"`
}

func (c *conf) getConf() *conf {
	yamlFile, err := ioutil.ReadFile("db.yaml")

	if err != nil {
		fmt.Println(err.Error())
	}

	if err = yaml.Unmarshal(yamlFile, c); err != nil {
		fmt.Println(err.Error())
	}

	return c
}

func (d *DB) getDBConfig() *DB {
	yamlFile, err := ioutil.ReadFile("conf.yaml")

	if err != nil {
		golog.Fatalf("111 %s", err)
	}

	//err = yaml.Unmarshal(yamlFile, &DBConfig)

	if err = yaml.Unmarshal(yamlFile, d); err != nil {
		golog.Fatalf("222 Unmarshal db config error!! %s", err)
	}

	return d
}

func main() {

	// 方式一：
	var d DB
	dconf := d.getDBConfig()
	fmt.Println(dconf.Master)
	fmt.Println(dconf.Master.Host)

	//方式二：
	var c conf
	conf := c.getConf()
	fmt.Println(conf.Host)
}
