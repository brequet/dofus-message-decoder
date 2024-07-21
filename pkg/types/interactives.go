package types

type Sign struct {
	Id      int    `json:"id"`
	Params  string `json:"params"`
	SkillId int    `json:"skillId"`
	TextKey int    `json:"textKey"`
}

type SkillName struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
}

type StealthBones struct {
	Id int `json:"id"`
}

type Interactive struct {
	Id             int  `json:"id"`
	NameId         int  `json:"nameId"`
	ActionId       int  `json:"actionId"`
	DisplayTooltip bool `json:"displayTooltip"`
}
