package Core

type World struct {
	Xsize     uint      `yaml:"Xsize"`
	Ysize     uint      `yaml:"Ysize"`
	BlockList [][]Block `yaml:"BlockList"`
}

func (world *World) InitWorld(x, y uint) {
	world.Xsize = x
	world.Ysize = y

}
