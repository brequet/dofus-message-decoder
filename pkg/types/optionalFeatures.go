package types

type CustomModeBreedSpell struct {
	Id             int  `json:"id"`
	PairId         int  `json:"pairId"`
	BreedId        int  `json:"breedId"`
	IsInitialSpell bool `json:"isInitialSpell"`
	IsHidden       bool `json:"isHidden"`
}

type ForgettableSpell struct {
	Id     int `json:"id"`
	PairId int `json:"pairId"`
	ItemId int `json:"itemId"`
}

type Modster struct {
	Id                        int   `json:"id"`
	ItemId                    int   `json:"itemId"`
	SpellIdKroma              int   `json:"spellIdKroma"`
	ItemIdKroma               int   `json:"itemIdKroma"`
	ModsterId                 int   `json:"modsterId"`
	ParentsModsterId          []int `json:"parentsModsterId"`
	ModsterActiveSpells       []int `json:"modsterActiveSpells"`
	ModsterPassiveSpells      []int `json:"modsterPassiveSpells"`
	ModsterHiddenAchievements []int `json:"modsterHiddenAchievements"`
	ModsterAchievements       []int `json:"modsterAchievements"`
	KromaHiddenAchievements   []int `json:"kromaHiddenAchievements"`
	KromaAchievements         []int `json:"kromaAchievements"`
	Order                     int   `json:"order"`
}
