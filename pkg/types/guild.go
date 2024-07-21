package types

type GuildRankNameSuggestion struct {
	UiKey string `json:"uiKey"`
}

type GuildRank struct {
	Id           int  `json:"id"`
	NameId       int  `json:"nameId"`
	Order        int  `json:"order"`
	IsModifiable bool `json:"isModifiable"`
	GfxId        int  `json:"gfxId"`
}

type GuildRightGroup struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
	Order  int `json:"order"`
	// rights vector custom subtype not implemented (1)
}

type GuildRight struct {
	Id      int `json:"id"`
	NameId  int `json:"nameId"`
	Order   int `json:"order"`
	GroupId int `json:"groupId"`
}

type GuildTag struct {
	Id     int `json:"id"`
	TypeId int `json:"typeId"`
	NameId int `json:"nameId"`
	Order  int `json:"order"`
}

type GuildTagsType struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
}

type GuildChestTab struct {
	TabId      int `json:"tabId"`
	NameId     int `json:"nameId"`
	Index      int `json:"index"`
	GfxId      int `json:"gfxId"`
	ServerType int `json:"serverType"`
	Cost       int `json:"cost"`
	Seniority  int `json:"seniority"`
	OpenRight  int `json:"openRight"`
	DropRight  int `json:"dropRight"`
	TakeRight  int `json:"takeRight"`
}
