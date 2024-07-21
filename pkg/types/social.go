package types

type EmblemBackground struct {
	Id    int `json:"id"`
	Order int `json:"order"`
}

type EmblemSymbolCategory struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
}

type EmblemSymbol struct {
	Id          int  `json:"id"`
	SkinId      int  `json:"skinId"`
	IconId      int  `json:"iconId"`
	Order       int  `json:"order"`
	CategoryId  int  `json:"categoryId"`
	Colorizable bool `json:"colorizable"`
}

type SocialRight struct {
	Id      int `json:"id"`
	NameId  int `json:"nameId"`
	Order   int `json:"order"`
	GroupId int `json:"groupId"`
}
