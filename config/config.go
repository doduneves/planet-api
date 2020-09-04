package config

import (
	"log"

	"github.com/BurntSushi/toml"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

type Config struct {
	Server   string
	Database string
}

func (c *Config) Connect() *mgo.Database {
	c.Read()
	session, err := mgo.Dial(c.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(c.Database)
	return db
}

func (c *Config) Read() {
	if _, err := toml.DecodeFile("config/config.toml", &c); err != nil {
		log.Fatal(err)
	}
}
