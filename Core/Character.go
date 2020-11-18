package Core

import (
	"FlatEarth/SharedLib"
	"gopkg.in/yaml.v2"
	"log"
)

type Character struct {
	Type       string        `yaml:"Type"`
	ID         int           `yaml:"ID"`
	Status     Status        `yaml:"Status,omitempty"`
	Cooldown   int64         `yaml:"Cooldown,omitempty"`
	LastStatus int64         `yaml:"LastStatus,omitempty"`
	Mountable  bool          `yaml:"Mountable"`
	Name       string        `yaml:"Name"`
	Rank       string        `yaml:"Rank"`
	Strength   int           `yaml:"Strength"`
	Life       int           `yaml:"Life"`
	Defense    int           `yaml:"Defense"`
	Speed      int           `yaml:"Speed"`
	Gold       int           `yaml:"Gold"`
	Backpack   []interface{} `yaml:"Backpack"`
	Wield      struct {
		Armor      string   `yaml:"Armor"`
		Weapon     string   `yaml:"Weapon"`
		Mount      string   `yaml:"Mount"`
		Additional []string `yaml:"Additional"`
	} `yaml:"Wield"`
	Skill     []string `yaml:"Skill"`
	EnemyList []Actor  `yaml:"Enemy,omitempty"`
}

type CharacterFactory struct {
	MaxID         uint        `yaml:"MaxID"`
	CharacterList []Character `yaml:"CharacterList,omitempty"`
}

func (charFac *CharacterFactory) Init(filename string) {
	var load CharacterFactory
	dat := SharedLib.ReadFile(filename)
	err := yaml.Unmarshal(dat, &load)
	SharedLib.PanicOnError(err, SharedLib.FATAL)
	charFac.MaxID = load.MaxID
}

func (charFac *CharacterFactory) Save(filename string) {
	dat, err := yaml.Marshal(charFac)
	SharedLib.PanicOnError(err, SharedLib.WARNING)
	if err == nil {
		SharedLib.WriteFile(dat, filename)
	}
}

func (charFac *CharacterFactory) Print() {
	log.Print("MaxID is ", charFac.MaxID)
	log.Print("CharacterList ", charFac.CharacterList)
}

func (charFac *CharacterFactory) Load(filename string) {
	dat := SharedLib.ReadFile(filename)
	err := yaml.Unmarshal(dat, &(charFac.CharacterList))
	SharedLib.PanicOnError(err, SharedLib.WARNING)
}
