package types

type PointOfInterest struct {
	Id         int `json:"id"`
	NameId     int `json:"nameId"`
	CategoryId int `json:"categoryId"`
}

type PointOfInterestCategory struct {
	Id            int `json:"id"`
	ActionLabelId int `json:"actionLabelId"`
}

type LegendaryTreasureHunt struct {
	Id        int     `json:"id"`
	NameId    int     `json:"nameId"`
	Level     int     `json:"level"`
	ChestId   uint    `json:"chestId"`
	MonsterId uint    `json:"monsterId"`
	MapItemId uint    `json:"mapItemId"`
	XpRatio   float64 `json:"xpRatio"`
}
