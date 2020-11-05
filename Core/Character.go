package Core

import (
	"FlatEarth/SharedLib"
	"gopkg.in/yaml.v2"
)

type Character []struct {
	Type      string        `yaml:"Type"`
	ID        int           `yaml:"ID"`
	Mountable bool          `yaml:"Mountable"`
	Name      string        `yaml:"Name"`
	Rank      string        `yaml:"Rank"`
	Strength  int           `yaml:"Strength"`
	Life      int           `yaml:"Life"`
	Defense   int           `yaml:"Defense"`
	Speed     int           `yaml:"Speed"`
	Gold      int           `yaml:"Gold"`
	Backpack  []interface{} `yaml:"Backpack"`
	Wield     struct {
		Armor      string   `yaml:"Armor"`
		Weapon     string   `yaml:"Weapon"`
		Mount      string   `yaml:"Mount"`
		Additional []string `yaml:"Additional"`
	} `yaml:"Wield"`
	Skill []string `yaml:"Skill"`
}

type CharacterFactory struct {
	MaxID uint `yaml:"MaxID"`
}

func (charFac *CharacterFactory) Init(filename string) {
	var load CharacterFactory
	dat := SharedLib.ReadFile(filename)
	err := yaml.Unmarshal(dat, &load)
	SharedLib.PanicOnError(err, SharedLib.FATAL)
	charFac.MaxID = load.MaxID
}
