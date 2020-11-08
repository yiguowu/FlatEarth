package main

import "FlatEarth/Core"

const worldFile = "../Config/World.yaml"
const systemconfig = "../Config/System.yaml"

func main() {
	world := new(Core.WorldInstance)
	world.InitWorldInstance(systemconfig, worldFile, 0, 0)
}
