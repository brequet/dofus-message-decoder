package types

type EffectInstanceInteger struct {
	TargetMask        string  `json:"targetMask"`
	VisibleInBuffUi   bool    `json:"visibleInBuffUi"`
	BaseEffectId      int     `json:"baseEffectId"`
	VisibleInFightLog bool    `json:"visibleInFightLog"`
	TargetId          int     `json:"targetId"`
	EffectElement     int     `json:"effectElement"`
	EffectUid         int     `json:"effectUid"`
	Dispellable       int     `json:"dispellable"`
	Triggers          string  `json:"triggers"`
	SpellId           int     `json:"spellId"`
	Duration          int     `json:"duration"`
	Random            float64 `json:"random"`
	EffectId          int     `json:"effectId"`
	Delay             int     `json:"delay"`
	VisibleOnTerrain  bool    `json:"visibleOnTerrain"`
	VisibleInTooltip  bool    `json:"visibleInTooltip"`
	RawZone           string  `json:"rawZone"`
	ForClientOnly     bool    `json:"forClientOnly"`
	Value             int     `json:"value"`
	Order             int     `json:"order"`
	Group             int     `json:"group"`
}

type EffectInstanceDice struct {
	TargetMask        string  `json:"targetMask"`
	DiceNum           int     `json:"diceNum"`
	VisibleInBuffUi   bool    `json:"visibleInBuffUi"`
	BaseEffectId      int     `json:"baseEffectId"`
	VisibleInFightLog bool    `json:"visibleInFightLog"`
	TargetId          int     `json:"targetId"`
	EffectElement     int     `json:"effectElement"`
	EffectUid         int     `json:"effectUid"`
	Dispellable       int     `json:"dispellable"`
	Triggers          string  `json:"triggers"`
	SpellId           int     `json:"spellId"`
	Duration          int     `json:"duration"`
	Random            float64 `json:"random"`
	EffectId          int     `json:"effectId"`
	Delay             int     `json:"delay"`
	DiceSide          int     `json:"diceSide"`
	VisibleOnTerrain  bool    `json:"visibleOnTerrain"`
	VisibleInTooltip  bool    `json:"visibleInTooltip"`
	RawZone           string  `json:"rawZone"`
	ForClientOnly     bool    `json:"forClientOnly"`
	Value             int     `json:"value"`
	Order             int     `json:"order"`
	Group             int     `json:"group"`
}
