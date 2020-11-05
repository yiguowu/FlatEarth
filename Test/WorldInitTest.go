package main

import "FlatEarth/Core"

const worldFile = "../Config/World.yaml"

func main() {
	world := new(Core.World)
	world.InitWorld(200, 200)
	world.Load(worldFile)
	world.RandomSet(5, 10, 70, 5, 10, 20, 40, 10, 35, 5, worldFile)
	world.Print()
}
