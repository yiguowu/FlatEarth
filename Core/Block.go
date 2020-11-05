package Core

type Block struct {
	Season    Season     `yaml:"Season"`
	Weather   Weather    `yaml:"Weather"`
	Land      Land       `yaml:"Land"`
	ItemList  []ItemPack `yaml:"ItemList, omitempty"`
	ActorList []Actor    `yaml:"ActorList, omitempty"`
}
