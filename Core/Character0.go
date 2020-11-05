package Core

type Character0 struct {
	Character
	CharacterName string
	Faction       string
	Sex           string
	CurrentLife   int
}

func (char *Character0) GetName() string {
	return char.OnGlance()
}

func (char *Character0) OnGlance() string {
	if char.CharacterName == "" {
		return char.Name
	} else {
		return char.CharacterName
	}
}

func (char *Character0) OnInspect() string {
	description := "这是一个玩家，"
	switch char.Sex {
	case "男":
		description = description + "他"
	case "女":
		description = description + "她"
	default:
		description = description + "祂"
	}
	switch char.Faction {
	case "":
		description = description + "无门无派。"
	default:
		description = description + "属于" + char.Faction + "。"
	}
	switch char.Sex {
	case "男":
		description = description + "他"
	case "女":
		description = description + "她"
	default:
		description = description + "祂"
	}
	if float32(char.CurrentLife)/float32(char.Life) > 0.75 {
		description = description + "看上去气色不错。"
	} else if float32(char.CurrentLife)/float32(char.Life) > 0.5 {
		description = description + "看上去状态一般。"
	} else if float32(char.CurrentLife)/float32(char.Life) > 0.25 {
		description = description + "看上去状态很差。"
	} else {
		description = description + "看上去奄奄一息。"
	}
	return description
}

func (char *Character0) OnProvoke(actor *Actor) string {
	description := (*actor).GetName() + "企图袭击" + char.GetName() + "。"
	return description
}

func (char *Character0) OnProvoked(actor *Actor, action string) string {
	switch action {
	case "袭击":
		return (*actor).GetName() + "要袭击你。"
	default:
		return "你不知道发生了什么事。"
	}
}
