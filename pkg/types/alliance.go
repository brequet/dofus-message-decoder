package types

type AllianceRightGroup struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
	Order  int `json:"order"`
	// rights vector custom subtype not implemented (1)
}

type AllianceRight struct {
	Id      int `json:"id"`
	NameId  int `json:"nameId"`
	Order   int `json:"order"`
	GroupId int `json:"groupId"`
}

type AllianceTag struct {
	Id     int `json:"id"`
	TypeId int `json:"typeId"`
	NameId int `json:"nameId"`
	Order  int `json:"order"`
}

type AllianceTagsType struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
}

type KothRole struct {
	Id        int  `json:"id"`
	NameId    int  `json:"nameId"`
	IsDefault bool `json:"isDefault"`
}

type AllianceRankNameSuggestion struct {
	UiKey string `json:"uiKey"`
}

type AllianceRank struct {
	Id           int  `json:"id"`
	NameId       int  `json:"nameId"`
	Order        int  `json:"order"`
	IsModifiable bool `json:"isModifiable"`
	GfxId        int  `json:"gfxId"`
}
