package Core

import (
	"FlatEarth/SharedLib"
	"gopkg.in/yaml.v2"
	"log"
	"math/rand"
	"time"
)

type World struct {
	Xsize     uint       `yaml:"Xsize"`
	Ysize     uint       `yaml:"Ysize"`
	BlockList [][]*Block `yaml:"BlockList"`
	Season    Season     `yaml:"Season"`
}

func (world *World) InitWorld(x, y uint) {
	world.Xsize = x
	world.Ysize = y
	world.Season = Spring
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

func (world *World) Load(filename string) {
	dat := SharedLib.ReadFile(filename)
	err := yaml.Unmarshal(dat, world)
	SharedLib.PanicOnError(err, SharedLib.WARNING)
}

func (world *World) Print() {
	var x, y uint
	for x = 0; x < world.Xsize; x++ {
		for y = 0; y < world.Ysize; y++ {
			log.Print(world.BlockList[x][y])
		}
	}
	log.Print(world.Xsize, world.Xsize, world.Season)
}

func (world *World) RandomSet(mountainShare, hillShare, plainShare, lakeShare, swampShare, grassShare int, sunny, rainy, cloudy, stormy int, filename string) {
	totalLand := mountainShare + hillShare + plainShare + lakeShare + swampShare + grassShare
	totalWeather := sunny + rainy + cloudy + stormy
	rand.Seed(time.Now().UTC().UnixNano())
	var x, y uint
	for x = 0; x < world.Xsize; x++ {
		for y = 0; y < world.Ysize; y++ {
			score := rand.Intn(totalLand)
			if score > (mountainShare + hillShare + plainShare + lakeShare + swampShare) {
				world.BlockList[x][y].Land = Grassland
			} else if score > (mountainShare + hillShare + plainShare + lakeShare) {
				world.BlockList[x][y].Land = Swamp
			} else if score > (mountainShare + hillShare + plainShare) {
				world.BlockList[x][y].Land = Lake
			} else if score > (mountainShare + hillShare) {
				world.BlockList[x][y].Land = Plain
			} else if score > mountainShare {
				world.BlockList[x][y].Land = Hill
			} else {
				world.BlockList[x][y].Land = Mountain
			}
			score = rand.Intn(totalWeather)
			if score > (sunny + rainy + cloudy) {
				world.BlockList[x][y].Weather = Stormy
			} else if score > (sunny + rainy) {
				world.BlockList[x][y].Weather = Cloudy
			} else if score > sunny {
				world.BlockList[x][y].Weather = Rainy
			} else {
				world.BlockList[x][y].Weather = Sunny
			}
		}
	}
	world.Save(filename)
}
