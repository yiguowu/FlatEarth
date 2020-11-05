package Core

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
