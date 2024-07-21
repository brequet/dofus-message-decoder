package types

type AlignmentRank struct {
	Id               int   `json:"id"`
	OrderId          int   `json:"orderId"`
	NameId           int   `json:"nameId"`
	DescriptionId    int   `json:"descriptionId"`
	MinimumAlignment int   `json:"minimumAlignment"`
	Gifts            []int `json:"gifts"`
}

type AlignmentRankJntGift struct {
	Id     int   `json:"id"`
	Gifts  []int `json:"gifts"`
	Levels []int `json:"levels"`
}

type AlignmentSide struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
}

type AlignmentTitle struct {
	SideId   int   `json:"sideId"`
	NamesId  []int `json:"namesId"`
	ShortsId []int `json:"shortsId"`
}

type AlignmentGift struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
}

type AlignmentOrder struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
	SideId int `json:"sideId"`
}
