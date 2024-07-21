package types

type BreachPrize struct {
	Id             int    `json:"id"`
	NameId         int    `json:"nameId"`
	CategoryId     int    `json:"categoryId"`
	Currency       int    `json:"currency"`
	TooltipKey     string `json:"tooltipKey"`
	DescriptionKey int    `json:"descriptionKey"`
	ItemId         int    `json:"itemId"`
}

type LuaFormula struct {
	Id          int    `json:"id"`
	FormulaName string `json:"formulaName"`
	LuaFormula  string `json:"luaFormula"`
}

type Month struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
}

type Pack struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	HasSubAreas bool   `json:"hasSubAreas"`
}

type Tips struct {
	Id     int `json:"id"`
	DescId int `json:"descId"`
}

type BreachBoss struct {
	Id                  int    `json:"id"`
	MonsterId           int    `json:"monsterId"`
	Category            int    `json:"category"`
	ApparitionCriterion string `json:"apparitionCriterion"`
	AccessCriterion     string `json:"accessCriterion"`
	IncompatibleBosses  []int  `json:"incompatibleBosses"`
	RewardId            uint   `json:"rewardId"`
}

type CensoredContent struct {
	Lang     string `json:"lang"`
	Type     int    `json:"type"`
	OldValue int    `json:"oldValue"`
	NewValue int    `json:"newValue"`
}

type CharacterXPMapping struct {
	Level            int     `json:"level"`
	ExperiencePoints float64 `json:"experiencePoints"`
}

type Subhint struct {
	Hint_id                    int     `json:"hint_id"`
	Hint_parent_uid            string  `json:"hint_parent_uid"`
	Hint_anchored_element      string  `json:"hint_anchored_element"`
	Hint_anchor                int     `json:"hint_anchor"`
	Hint_position_x            int     `json:"hint_position_x"`
	Hint_position_y            int     `json:"hint_position_y"`
	Hint_width                 int     `json:"hint_width"`
	Hint_height                int     `json:"hint_height"`
	Hint_highlighted_element   string  `json:"hint_highlighted_element"`
	Hint_order                 int     `json:"hint_order"`
	Hint_tooltip_text          int     `json:"hint_tooltip_text"`
	Hint_tooltip_position_enum int     `json:"hint_tooltip_position_enum"`
	Hint_tooltip_url           string  `json:"hint_tooltip_url"`
	Hint_tooltip_offset_x      int     `json:"hint_tooltip_offset_x"`
	Hint_tooltip_offset_y      int     `json:"hint_tooltip_offset_y"`
	Hint_tooltip_width         int     `json:"hint_tooltip_width"`
	Hint_creation_date         float64 `json:"hint_creation_date"`
}
