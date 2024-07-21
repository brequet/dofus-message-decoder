package types

type EffectInstance struct {
	EffectUid         int     `json:"effectUid"`
	BaseEffectId      int     `json:"baseEffectId"`
	EffectId          int     `json:"effectId"`
	Order             int     `json:"order"`
	TargetId          int     `json:"targetId"`
	TargetMask        string  `json:"targetMask"`
	Duration          int     `json:"duration"`
	Random            float64 `json:"random"`
	Group             int     `json:"group"`
	VisibleInTooltip  bool    `json:"visibleInTooltip"`
	VisibleInBuffUi   bool    `json:"visibleInBuffUi"`
	VisibleInFightLog bool    `json:"visibleInFightLog"`
	VisibleOnTerrain  bool    `json:"visibleOnTerrain"`
	ForClientOnly     bool    `json:"forClientOnly"`
	Dispellable       int     `json:"dispellable"`
	RawZone           string  `json:"rawZone"`
	Delay             int     `json:"delay"`
	Triggers          string  `json:"triggers"`
	EffectElement     int     `json:"effectElement"`
	SpellId           int     `json:"spellId"`
}

type Effect struct {
	Id                       int     `json:"id"`
	DescriptionId            int     `json:"descriptionId"`
	IconId                   int     `json:"iconId"`
	Characteristic           int     `json:"characteristic"`
	Category                 int     `json:"category"`
	Operator                 string  `json:"operator"`
	ShowInTooltip            bool    `json:"showInTooltip"`
	UseDice                  bool    `json:"useDice"`
	ForceMinMax              bool    `json:"forceMinMax"`
	Boost                    bool    `json:"boost"`
	Active                   bool    `json:"active"`
	OppositeId               int     `json:"oppositeId"`
	TheoreticalDescriptionId int     `json:"theoreticalDescriptionId"`
	TheoreticalPattern       int     `json:"theoreticalPattern"`
	ShowInSet                bool    `json:"showInSet"`
	BonusType                int     `json:"bonusType"`
	UseInFight               bool    `json:"useInFight"`
	EffectPriority           int     `json:"effectPriority"`
	EffectPowerRate          float64 `json:"effectPowerRate"`
	ElementId                int     `json:"elementId"`
}

type EvolutiveEffect struct {
	Id       int `json:"id"`
	ActionId int `json:"actionId"`
	TargetId int `json:"targetId"`
	// progressionPerLevelRange vector subtype not implemented
}
