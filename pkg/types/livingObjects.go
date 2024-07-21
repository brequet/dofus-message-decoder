package types

type LivingObjectSkinJntMood struct {
	SkinId int `json:"skinId"`
	// moods vector subtype not implemented
}

type SpeakingItemText struct {
	TextId          int     `json:"textId"`
	TextProba       float64 `json:"textProba"`
	TextStringId    int     `json:"textStringId"`
	TextLevel       int     `json:"textLevel"`
	TextSound       int     `json:"textSound"`
	TextRestriction string  `json:"textRestriction"`
}

type SpeakingItemsTrigger struct {
	TriggersId int   `json:"triggersId"`
	TextIds    []int `json:"textIds"`
	States     []int `json:"states"`
}
