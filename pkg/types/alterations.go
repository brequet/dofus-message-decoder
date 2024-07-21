package types

type AlterationCategory struct {
	Id       int `json:"id"`
	NameId   int `json:"nameId"`
	ParentId int `json:"parentId"`
}

type Alteration struct {
	Id            int    `json:"id"`
	NameId        int    `json:"nameId"`
	DescriptionId int    `json:"descriptionId"`
	CategoryId    int    `json:"categoryId"`
	IconId        int    `json:"iconId"`
	IsVisible     bool   `json:"isVisible"`
	Criteria      string `json:"criteria"`
	IsWebDisplay  bool   `json:"isWebDisplay"`
	// possibleEffects vector custom subtype not implemented (1)
}
