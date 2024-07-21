package types

type HavenbagFurniture struct {
	TypeId         int  `json:"typeId"`
	ThemeId        int  `json:"themeId"`
	ElementId      int  `json:"elementId"`
	Color          int  `json:"color"`
	LayerId        int  `json:"layerId"`
	BlocksMovement bool `json:"blocksMovement"`
	IsStackable    bool `json:"isStackable"`
	CellsWidth     uint `json:"cellsWidth"`
	CellsHeight    uint `json:"cellsHeight"`
}

type HavenbagTheme struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
	MapId  int `json:"mapId"`
}

type House struct {
	TypeId        int `json:"typeId"`
	DefaultPrice  int `json:"defaultPrice"`
	NameId        int `json:"nameId"`
	DescriptionId int `json:"descriptionId"`
	GfxId         int `json:"gfxId"`
}
