package Core

type Block struct {
	Season    Season     `yaml:"Season"`
	Weather   Weather    `yaml:"Weather"`
	Land      Land       `yaml:"Land"`
	ItemList  []ItemPack `yaml:"ItemList,omitempty"`
	ActorList []Actor    `yaml:"ActorList,omitempty"`
}

func (block *Block) InitToDefault() {
	block.Land = Plain
	block.Weather = Sunny
	block.Season = Spring
	block.ItemList = make([]ItemPack, 0)
	block.ActorList = make([]Actor, 0)
}
