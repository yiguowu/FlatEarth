package Core

type Color string

const (
	Red    Color = "\033[31m"
	Green  Color = "\033[32m"
	Yellow Color = "\033[33m"
	Blue   Color = "\033[34m"
	Purple Color = "\033[35m"
	Cyan   Color = "\033[36m"
	White  Color = "\033[37m"
)

type Land string

const (
	Plain     Land = "平原"
	Hill      Land = "丘陵"
	Mountain  Land = "大山"
	Lake      Land = "湖泊"
	Swamp     Land = "沼泽"
	Grassland Land = "草地"
)

type Weather string

const (
	Sunny  Weather = "晴天"
	Rainy  Weather = "雨天"
	Cloudy Weather = "多云"
	Stormy Weather = "风暴"
)

var WeatherChangeMatrix = [4][4]float32{
	{0.7, 0.1, 0.2, 0.0},
	{0.3, 0.2, 0.4, 0.1},
	{0.5, 0.2, 0.3, 0.0},
	{0.3, 0.3, 0.3, 0.1},
}

type Item string

const (
	Tree Item = "树"
	Rock Item = "岩石"
)

type ItemPack struct {
	Item     Item `yaml:"Item"`
	Quantity int  `yaml:"Quantity"`
}

type Season string

const (
	Winter Season = "冬天"
	Spring Season = "春天"
	Summer Season = "夏天"
	Fall   Season = "秋天"
)

type WorldEventType string

const (
	WeatherChange WorldEventType = "天气变化"
	SeasonChange  WorldEventType = "季节变化"
)

type System struct {
	Hour     int     `yaml:"hour"`
	Weather  float64 `yaml:"weather"`
	AutoSave int     `yaml:"autosave"`
}

const MaxEvent = 65536
