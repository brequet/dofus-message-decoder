package types

type CharacteristicCategory struct {
	Id                int    `json:"id"`
	NameId            int    `json:"nameId"`
	Order             int    `json:"order"`
	CharacteristicIds []uint `json:"characteristicIds"`
}

type Characteristic struct {
	Id             int    `json:"id"`
	Keyword        string `json:"keyword"`
	NameId         int    `json:"nameId"`
	Asset          string `json:"asset"`
	CategoryId     int    `json:"categoryId"`
	Visible        bool   `json:"visible"`
	Order          int    `json:"order"`
	ScaleFormulaId int    `json:"scaleFormulaId"`
	Upgradable     bool   `json:"upgradable"`
}
