package types

type BreachWorldMapSector struct {
	Id           int    `json:"id"`
	SectorNameId int    `json:"sectorNameId"`
	LegendId     int    `json:"legendId"`
	SectorIcon   string `json:"sectorIcon"`
	MinStage     int    `json:"minStage"`
	MaxStage     int    `json:"maxStage"`
}

type BreachDungeonModificator struct {
	Id                      int     `json:"id"`
	ModificatorId           int     `json:"modificatorId"`
	Criterion               string  `json:"criterion"`
	AdditionalRewardPercent float64 `json:"additionalRewardPercent"`
	Score                   float64 `json:"score"`
	IsPositiveForPlayers    bool    `json:"isPositiveForPlayers"`
	TooltipBaseline         string  `json:"tooltipBaseline"`
}

type BreachInfinityLevel struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
	Level  int `json:"level"`
}

type BreachWorldMapCoordinate struct {
	Id                int `json:"id"`
	MapStage          int `json:"mapStage"`
	MapCoordinateX    int `json:"mapCoordinateX"`
	MapCoordinateY    int `json:"mapCoordinateY"`
	UnexploredMapIcon int `json:"unexploredMapIcon"`
	ExploredMapIcon   int `json:"exploredMapIcon"`
}
