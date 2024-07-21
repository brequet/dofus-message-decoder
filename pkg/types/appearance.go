package types

type Title struct {
	Id           int  `json:"id"`
	NameMaleId   int  `json:"nameMaleId"`
	NameFemaleId int  `json:"nameFemaleId"`
	Visible      bool `json:"visible"`
	CategoryId   int  `json:"categoryId"`
}

type Appearance struct {
	Id   int    `json:"id"`
	Type int    `json:"type"`
	Data string `json:"data"`
}

type CreatureBoneOverride struct {
	BoneId         int `json:"boneId"`
	CreatureBoneId int `json:"creatureBoneId"`
}

type CreatureBoneType struct {
	Id             int `json:"id"`
	CreatureBoneId int `json:"creatureBoneId"`
}

type Ornament struct {
	Id      int  `json:"id"`
	NameId  int  `json:"nameId"`
	Visible bool `json:"visible"`
	AssetId int  `json:"assetId"`
	IconId  int  `json:"iconId"`
	Order   int  `json:"order"`
}

type SkinMapping struct {
	Id       int `json:"id"`
	LowDefId int `json:"lowDefId"`
}

type SkinPosition struct {
	Id int `json:"id"`
	// transformation vector custom subtype not implemented (1)
	Clip []string `json:"clip"`
	Skin []uint   `json:"skin"`
}

type TitleCategory struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
}
