package main

import "FlatEarth/Core"

const charConfig = "../Config/CharacterFactoryState.yaml"
const charListConfig = "../Config/Character.yaml"

func main() {
	char := new(Core.CharacterFactory)
	char.Init(charConfig)
	char.Load(charListConfig)
	char.Print()
	char.Save(charConfig)
}
