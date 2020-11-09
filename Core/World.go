package Core

import (
	"FlatEarth/SharedLib"
	"errors"
	"gopkg.in/yaml.v2"
	"log"
	"math/rand"
	"time"
)

type WorldInstance struct {
	World      *World
	MessageBus *chan WorldEvent
}

type World struct {
	Xsize       int        `yaml:"Xsize"`
	Ysize       int        `yaml:"Ysize"`
	BlockList   [][]*Block `yaml:"BlockList"`
	Season      Season     `yaml:"Season"`
	CurrentTime string     `yaml:"CurrentTime,omitempty"`
}

type WorldEvent struct {
	EventType   WorldEventType
	LocationX   int
	LocationY   int
	Description string
	Timestamp   int64
	Source      *Actor
	Target      []*Actor
}

type Timer struct {
	Name       string
	LastFired  int64
	Period     time.Duration
	MessageBus *chan WorldEvent
}

func EventSink(evtBus *chan WorldEvent) {
	for {
		event := <-*evtBus
		switch event.EventType {
		case WeatherChange:
			continue
		case SeasonChange:
			log.Println(event.Description)
		default:
			SharedLib.PanicOnError(errors.New("UnknownEvent"), SharedLib.WARNING)
		}
	}
}

func WeatherCheck(world *World, evtBus *chan WorldEvent) {
	for x := 0; x < world.Xsize; x++ {
		for y := 0; y < world.Ysize; y++ {
			previousWeather := world.BlockList[x][y].Weather
			world.BlockList[x][y].Weather = CalculateWeather(WeatherChangeMatrix, world.BlockList[x][y].Weather)
			if previousWeather != world.BlockList[x][y].Weather {
				var event WorldEvent
				event.LocationX = x
				event.LocationY = y
				event.EventType = WeatherChange
				event.Timestamp = time.Now().UnixNano()
				switch world.BlockList[x][y].Weather {
				case Sunny:
					if previousWeather == Cloudy {
						event.Description = "云散了"
					} else {
						event.Description = "天放晴了"
					}
				case Rainy:
					if previousWeather == Stormy {
						event.Description = "雨小了"
					} else {
						event.Description = "现在下起雨来"
					}
				case Cloudy:
					if previousWeather == Sunny {
						event.Description = "天上飘过几朵云"
					} else {
						event.Description = "雨停了，天上只剩几朵云"
					}
				case Stormy:
					event.Description = "天上下起暴雨来"
				}
				*evtBus <- event
			}
		}
	}
}

func SeasonCheck(world *World, t time.Time, evtBus *chan WorldEvent) {
	switch t.Month() {
	case time.January, time.February, time.March:
		if world.Season != Winter {
			var event WorldEvent
			event.EventType = SeasonChange
			event.Timestamp = time.Now().UnixNano()
			event.Description = "冬天到了"
			world.Season = Winter
			*evtBus <- event
		}
	case time.April, time.May, time.June:
		if world.Season != Spring {
			var event WorldEvent
			event.EventType = SeasonChange
			event.Timestamp = time.Now().UnixNano()
			event.Description = "春天到了"
			world.Season = Spring
			*evtBus <- event
		}
	case time.July, time.August, time.September:
		if world.Season != Summer {
			var event WorldEvent
			event.EventType = SeasonChange
			event.Timestamp = time.Now().UnixNano()
			event.Description = "夏天到了"
			world.Season = Summer
			*evtBus <- event
		}
	case time.October, time.November, time.December:
		if world.Season != Fall {
			var event WorldEvent
			event.EventType = SeasonChange
			event.Timestamp = time.Now().UnixNano()
			event.Description = "秋天到了"
			world.Season = Fall
			*evtBus <- event
		}
	}
}

func GlobalTime(interval int, world *World, evtBus *chan WorldEvent) {
	if world.CurrentTime == "" {
		world.CurrentTime = "2000-04-01T00:00:00.000Z"
	}
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, world.CurrentTime)
	SharedLib.PanicOnError(err, SharedLib.WARNING)
	sleepTime := time.Duration(interval) * time.Second
	for {
		time.Sleep(sleepTime)
		t = t.Add(1 * time.Hour)
		world.CurrentTime = t.Format(layout)
		go SeasonCheck(world, t, evtBus)
		go WeatherCheck(world, evtBus)
	}
}

func (ins *WorldInstance) InitWorldInstance(system, filename string, x, y int) {

	ins.World = new(World)
	if filename == "" {
		ins.World.InitWorld(x, y)
	} else {
		ins.World.Load(filename)
	}
	dat := SharedLib.ReadFile(system)

	var sys System
	err := yaml.Unmarshal(dat, &sys)
	SharedLib.PanicOnError(err, SharedLib.FATAL)
	eventBus := make(chan WorldEvent, MaxEvent)
	go GlobalTime(sys.Hour, ins.World, &eventBus)
	go EventSink(&eventBus)
	go ins.World.AutoSave(filename, sys.AutoSave)
	for {
		time.Sleep(60 * time.Hour)
	}
}

func (world *World) InitWorld(x, y int) {
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

func (world *World) AutoSave(filename string, interval int) {
	for {
		time.Sleep(time.Duration(interval) * time.Second)
		log.Println("Saving")
		world.Save(filename)
		log.Println("Done")
	}
}

func (world *World) Print() {
	var x, y int
	for x = 0; x < world.Xsize; x++ {
		for y = 0; y < world.Ysize; y++ {
			log.Print(world.BlockList[x][y])
		}
	}
	log.Print(world.Xsize, world.Xsize, world.Season)
}

func (world *World) RandomSet(mountainShare, hillShare, plainShare, lakeShare, swampShare, grassShare int, filename string) {
	totalLand := mountainShare + hillShare + plainShare + lakeShare + swampShare + grassShare
	rand.Seed(time.Now().UTC().UnixNano())
	var x, y int
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
		}
	}
	world.Save(filename)
}
