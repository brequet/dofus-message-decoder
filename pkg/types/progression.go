package types

type ActivitySuggestion struct {
	Id            int     `json:"id"`
	NameId        int     `json:"nameId"`
	DescriptionId int     `json:"descriptionId"`
	CategoryId    int     `json:"categoryId"`
	Level         uint    `json:"level"`
	MapId         int     `json:"mapId"`
	IsLarge       bool    `json:"isLarge"`
	StartDate     float64 `json:"startDate"`
	EndDate       float64 `json:"endDate"`
	Icon          string  `json:"icon"`
}

type ActivitySuggestionsCategory struct {
	Id       int `json:"id"`
	NameId   int `json:"nameId"`
	ParentId int `json:"parentId"`
}

type FeatureDescription struct {
	Id            int    `json:"id"`
	NameId        int    `json:"nameId"`
	DescriptionId int    `json:"descriptionId"`
	Priority      int    `json:"priority"`
	ParentId      int    `json:"parentId"`
	Children      []int  `json:"children"`
	Criterion     string `json:"criterion"`
}
