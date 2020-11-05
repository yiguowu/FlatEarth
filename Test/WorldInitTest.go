package main

import "FlatEarth/Core"

const worldFile = "../Config/World.yaml"

func main() {
	world := new(Core.World)
	world.InitWorld(200, 200)
	world.Load(worldFile)
	world.Print()
}
