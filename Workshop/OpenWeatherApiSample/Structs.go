package main

type OpenWeatherMap struct {
	Coord      CoordType     `json:"coord"`
	Weather    []WeatherType `json:"weather"`
	Base       string        `json:"base"`
	Main       MainType      `json:"main"`
	Visibility int           `json:"visibility"`
	Wind       WindType      `json:"wind"`
	Cloud      CloudType     `json:"clouds"`
	Rain       RainType      `json:"rain"`
	Dt         int           `json:"dt"`
	Sys        SysType       `json:"sys"`
	Id         int           `json:"id"`
	Name       string        `json:"name"`
	Cod        int           `json:"cod"`
}

type CoordType struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

type WeatherType struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type MainType struct {
	Temp        float32 `json:"temp"`
	Pressure    float32 `json:"pressure"`
	Humidity    float32 `json:"humidity"`
	MinTemp     float32 `json:"temp_min"`
	MaxTemp     float32 `json:"temp_max"`
	SeaLevel    float32 `json:"sea_level"`
	GroundLevel float32 `json:"grnd_level"`
}

type WindType struct {
	Speed float32 `json:"speed"`
	Deg   float32 `json:"deg"`
}

type CloudType struct {
	All float32 `json:"all"`
}

type RainType struct {
	T3h float32 `json:"3h"`
}

type SnowType struct {
	T3h float32 `json:"3h"`
}

type CalDateType struct {
	CDate string `json:"dt"`
}

type SysType struct {
	Type    int     `json:"type"`
	Id      int     `json:"id"`
	Message float32 `json:"message"`
	Country string  `json:"country"`
	Sunrise int     `json:"sunrise"`
	SunSet  int     `json:"sunset"`
}
