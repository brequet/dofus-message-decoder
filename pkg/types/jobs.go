package types

type Job struct {
	Id                int  `json:"id"`
	NameId            int  `json:"nameId"`
	IconId            int  `json:"iconId"`
	HasLegendaryCraft bool `json:"hasLegendaryCraft"`
}

type Recipe struct {
	ResultId              int     `json:"resultId"`
	ResultNameId          int     `json:"resultNameId"`
	ResultTypeId          uint    `json:"resultTypeId"`
	ResultLevel           int     `json:"resultLevel"`
	IngredientIds         []int   `json:"ingredientIds"`
	Quantities            []uint  `json:"quantities"`
	JobId                 int     `json:"jobId"`
	SkillId               int     `json:"skillId"`
	ChangeVersion         string  `json:"changeVersion"`
	TooltipExpirationDate float64 `json:"tooltipExpirationDate"`
}

type Skill struct {
	Id                    int    `json:"id"`
	NameId                int    `json:"nameId"`
	ParentJobId           int    `json:"parentJobId"`
	IsForgemagus          bool   `json:"isForgemagus"`
	ModifiableItemTypeIds []int  `json:"modifiableItemTypeIds"`
	GatheredRessourceItem int    `json:"gatheredRessourceItem"`
	CraftableItemIds      []int  `json:"craftableItemIds"`
	InteractiveId         int    `json:"interactiveId"`
	Range                 int    `json:"range"`
	UseRangeInClient      bool   `json:"useRangeInClient"`
	UseAnimation          string `json:"useAnimation"`
	Cursor                int    `json:"cursor"`
	ElementActionId       int    `json:"elementActionId"`
	AvailableInHouse      bool   `json:"availableInHouse"`
	ClientDisplay         bool   `json:"clientDisplay"`
	LevelMin              int    `json:"levelMin"`
}
