package types

type MapReference struct {
	Id     int `json:"id"`
	MapId  int `json:"mapId"`
	CellId int `json:"cellId"`
}

type MapScrollAction struct {
	Id           float64 `json:"id"`
	RightExists  bool    `json:"rightExists"`
	BottomExists bool    `json:"bottomExists"`
	LeftExists   bool    `json:"leftExists"`
	TopExists    bool    `json:"topExists"`
	RightMapId   float64 `json:"rightMapId"`
	BottomMapId  float64 `json:"bottomMapId"`
	LeftMapId    float64 `json:"leftMapId"`
	TopMapId     float64 `json:"topMapId"`
}

type Dungeon struct {
	Id                 int       `json:"id"`
	NameId             int       `json:"nameId"`
	OptimalPlayerLevel int       `json:"optimalPlayerLevel"`
	MapIds             []float64 `json:"mapIds"`
	EntranceMapId      float64   `json:"entranceMapId"`
	ExitMapId          float64   `json:"exitMapId"`
}

type HintCategory struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
}

type MapPosition struct {
	Id           float64 `json:"id"`
	PosX         int     `json:"posX"`
	PosY         int     `json:"posY"`
	Outdoor      bool    `json:"outdoor"`
	Capabilities int     `json:"capabilities"`
	NameId       int     `json:"nameId"`
	// playlists vector subtype not implemented
	SubAreaId              int  `json:"subAreaId"`
	WorldMap               int  `json:"worldMap"`
	HasPriorityOnWorldmap  bool `json:"hasPriorityOnWorldmap"`
	AllowPrism             bool `json:"allowPrism"`
	IsTransition           bool `json:"isTransition"`
	MapHasTemplate         bool `json:"mapHasTemplate"`
	TacticalModeTemplateId int  `json:"tacticalModeTemplateId"`
	HasPublicPaddock       bool `json:"hasPublicPaddock"`
}

type SubArea struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
	AreaId int `json:"areaId"`
	// playlists vector subtype not implemented
	MapIds               []float64 `json:"mapIds"`
	Shape                []int     `json:"shape"`
	WorldmapId           int       `json:"worldmapId"`
	CustomWorldMap       []uint    `json:"customWorldMap"`
	PackId               uint      `json:"packId"`
	Level                uint      `json:"level"`
	IsConquestVillage    bool      `json:"isConquestVillage"`
	BasicAccountAllowed  bool      `json:"basicAccountAllowed"`
	DisplayOnWorldMap    bool      `json:"displayOnWorldMap"`
	MountAutoTripAllowed bool      `json:"mountAutoTripAllowed"`
	PsiAllowed           bool      `json:"psiAllowed"`
	Monsters             []uint    `json:"monsters"`
	Capturable           bool      `json:"capturable"`
	// quests vector subtype not implemented
	// npcs vector subtype not implemented
	Harvestables        []int `json:"harvestables"`
	AssociatedZaapMapId int   `json:"associatedZaapMapId"`
	Neighbors           []int `json:"neighbors"`
}

type SuperArea struct {
	Id          int  `json:"id"`
	NameId      int  `json:"nameId"`
	WorldmapId  int  `json:"worldmapId"`
	HasWorldMap bool `json:"hasWorldMap"`
}

type Waypoint struct {
	Id        int  `json:"id"`
	MapId     int  `json:"mapId"`
	SubAreaId int  `json:"subAreaId"`
	Activated bool `json:"activated"`
}

type WorldMap struct {
	Id                 int      `json:"id"`
	NameId             int      `json:"nameId"`
	OrigineX           int      `json:"origineX"`
	OrigineY           int      `json:"origineY"`
	MapWidth           float64  `json:"mapWidth"`
	MapHeight          float64  `json:"mapHeight"`
	ViewableEverywhere bool     `json:"viewableEverywhere"`
	MinScale           float64  `json:"minScale"`
	MaxScale           float64  `json:"maxScale"`
	StartScale         float64  `json:"startScale"`
	TotalWidth         int      `json:"totalWidth"`
	TotalHeight        int      `json:"totalHeight"`
	Zoom               []string `json:"zoom"`
	VisibleOnMap       bool     `json:"visibleOnMap"`
}

type Area struct {
	Id              int  `json:"id"`
	NameId          int  `json:"nameId"`
	SuperAreaId     int  `json:"superAreaId"`
	ContainHouses   bool `json:"containHouses"`
	ContainPaddocks bool `json:"containPaddocks"`
	// bounds custom type not implemented (1)
	WorldmapId    int  `json:"worldmapId"`
	HasWorldMap   bool `json:"hasWorldMap"`
	HasSuggestion bool `json:"hasSuggestion"`
}

type Hint struct {
	Id         int  `json:"id"`
	CategoryId int  `json:"categoryId"`
	Gfx        int  `json:"gfx"`
	NameId     int  `json:"nameId"`
	MapId      int  `json:"mapId"`
	RealMapId  int  `json:"realMapId"`
	X          int  `json:"x"`
	Y          int  `json:"y"`
	Outdoor    bool `json:"outdoor"`
	SubareaId  int  `json:"subareaId"`
	WorldMapId int  `json:"worldMapId"`
	Level      uint `json:"level"`
}

type MapCoordinates struct {
	CompressedCoords int       `json:"compressedCoords"`
	MapIds           []float64 `json:"mapIds"`
}
