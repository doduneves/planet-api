package models

import (
	. "github.com/doduneves/planet-api/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Planet struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Nome      string        `bson:"nome" json:"nome"`
	Clima     string        `bson:"clima" json:"clima"`
	Terreno   string        `bson:"terreno" json:"terreno"`
	Aparicoes int           `bson:"aparicoes" json:"aparicoes"`
}

var db *mgo.Database
var config = Config{}

const (
	COLLECTION = "planets"
)

func (p *Planet) SetAppearance(aparicoes int) {
	p.Aparicoes = aparicoes
}

func (p Planet) GetAll() ([]Planet, error) {
	db = config.Connect()
	var planets []Planet

	err := db.C(COLLECTION).Find(bson.M{}).All(&planets)
	return planets, err
}

func (p Planet) GetByID(id string) (Planet, error) {
	db = config.Connect()

	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&p)
	return p, err
}

func (p Planet) Create(planet Planet) error {
	db = config.Connect()

	err := db.C(COLLECTION).Insert(&planet)
	return err
}

func (p Planet) Delete(id string) error {
	db = config.Connect()
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (p Planet) Update(id string) error {
	db = config.Connect()
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &p)
	return err
}
