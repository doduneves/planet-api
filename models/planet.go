package models

import (
	"errors"
	"fmt"

	. "github.com/doduneves/planet-api/config"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Planet struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Nome      string        `bson:"nome" json:"nome" validate:"required"`
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

func (p Planet) GetAll(nome string, page int) ([]Planet, error) {
	db = config.Connect()
	var planets []Planet
	var err error

	skip := (page - 1) * 10

	if nome != "" {
		err = db.C(COLLECTION).Find(bson.M{"nome": nome}).Skip(skip).Limit(10).All(&planets)

	} else {
		err = db.C(COLLECTION).Find(bson.M{}).Skip(skip).Limit(10).All(&planets)
	}
	return planets, err
}

func (p Planet) GetByID(id string) (Planet, error) {
	db = config.Connect()

	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&p)
	return p, err
}

func (p Planet) Create(planet Planet) error {
	db = config.Connect()

	fmt.Println(planet)

	v := validator.New()
	err := v.Struct(planet)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e)
		}
	}

	count, errCount := db.C(COLLECTION).Find(bson.M{"nome": planet.Nome}).Count()

	if errCount != nil {
		return errCount
	} else {
		if count == 0 {
			if err == nil {
				err = db.C(COLLECTION).Insert(&planet)
			}
		} else {
			err = errors.New("The planet " + planet.Nome + " already exists")
		}
	}
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
