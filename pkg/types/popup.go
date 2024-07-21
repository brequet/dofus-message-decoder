package types

type PopupButton struct {
	Id          int    `json:"id"`
	PopupId     int    `json:"popupId"`
	Type        int    `json:"type"`
	TextId      int    `json:"textId"`
	ActionType  int    `json:"actionType"`
	ActionValue string `json:"actionValue"`
}

type PopupInformation struct {
	Id            int    `json:"id"`
	ParentId      int    `json:"parentId"`
	TitleId       int    `json:"titleId"`
	DescriptionId int    `json:"descriptionId"`
	IlluName      string `json:"illuName"`
	// buttons vector custom subtype not implemented (1)
	Criterion   string `json:"criterion"`
	CacheType   int    `json:"cacheType"`
	AutoTrigger bool   `json:"autoTrigger"`
}
