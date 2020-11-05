package Core

import (
	"FlatEarth/SharedLib"
	"gopkg.in/yaml.v2"
)

type World struct {
	Xsize     uint       `yaml:"Xsize"`
	Ysize     uint       `yaml:"Ysize"`
	BlockList [][]*Block `yaml:"BlockList"`
}

func (world *World) InitWorld(x, y uint) {
	world.Xsize = x
	world.Ysize = y
	world.BlockList = make([][]*Block, world.Xsize)
	for x = 0; x < world.Xsize; x++ {
		world.BlockList[x] = make([]*Block, world.Ysize)
		for y = 0; y < world.Ysize; y++ {
			world.BlockList[x][y] = new(Block)
			world.BlockList[x][y].InitToDefault()
		}
	}
}

func (world *World) Save(filename string) {
	dat, err := yaml.Marshal(world)
	SharedLib.PanicOnError(err, SharedLib.WARNING)
	if err == nil {
		SharedLib.WriteFile(dat, filename)
	}
}
