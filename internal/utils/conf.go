package utils

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

// Conf Conf
type Conf struct {
	Cookie  string `yaml:"cookie"`
	Auth    string `yaml:"auth"`
	HTMLDir string `yaml:"htmlDir"`
}

// GetConf GetConf
func GetConf() *Conf {
	yamlFile, err := ioutil.ReadFile("../../cmd/lagoucoursepull/conf.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}
	c := &Conf{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatal(err.Error())
	}
	return c
}

// GetCookie GetCookie
func (c *Conf) GetCookie() string {
	return c.Cookie
}

// GetAuth GetAuth
func (c *Conf) GetAuth() string {
	return c.Auth
}

// GetHTMLDir GetHTMLDir
func (c *Conf) GetHTMLDir() string {
	return c.HTMLDir
}
