package types

type SpellVariant struct {
	Id       int    `json:"id"`
	BreedId  int    `json:"breedId"`
	SpellIds []uint `json:"spellIds"`
}

type SpellBomb struct {
	Id                   int `json:"id"`
	ChainReactionSpellId int `json:"chainReactionSpellId"`
	ExplodSpellId        int `json:"explodSpellId"`
	WallId               int `json:"wallId"`
	InstantSpellId       int `json:"instantSpellId"`
	ComboCoeff           int `json:"comboCoeff"`
}

type SpellType struct {
	Id          int `json:"id"`
	LongNameId  int `json:"longNameId"`
	ShortNameId int `json:"shortNameId"`
}

type EffectZone struct {
	Id                         int    `json:"id"`
	RawDisplayZone             string `json:"rawDisplayZone"`
	IsDefaultPreviewZoneHidden bool   `json:"isDefaultPreviewZoneHidden"`
	CasterMask                 string `json:"casterMask"`
	RawActivationZone          string `json:"rawActivationZone"`
	ActivationMask             string `json:"activationMask"`
}

type SpellLevel struct {
	Id                        int    `json:"id"`
	SpellId                   int    `json:"spellId"`
	Grade                     int    `json:"grade"`
	SpellBreed                int    `json:"spellBreed"`
	ApCost                    int    `json:"apCost"`
	MinRange                  int    `json:"minRange"`
	Range                     int    `json:"range"`
	CastInLine                bool   `json:"castInLine"`
	CastInDiagonal            bool   `json:"castInDiagonal"`
	CastTestLos               bool   `json:"castTestLos"`
	CriticalHitProbability    int    `json:"criticalHitProbability"`
	NeedFreeCell              bool   `json:"needFreeCell"`
	NeedTakenCell             bool   `json:"needTakenCell"`
	NeedVisibleEntity         bool   `json:"needVisibleEntity"`
	NeedCellWithoutPortal     bool   `json:"needCellWithoutPortal"`
	PortalProjectionForbidden bool   `json:"portalProjectionForbidden"`
	NeedFreeTrapCell          bool   `json:"needFreeTrapCell"`
	RangeCanBeBoosted         bool   `json:"rangeCanBeBoosted"`
	MaxStack                  int    `json:"maxStack"`
	MaxCastPerTurn            int    `json:"maxCastPerTurn"`
	MaxCastPerTarget          int    `json:"maxCastPerTarget"`
	MinCastInterval           int    `json:"minCastInterval"`
	InitialCooldown           int    `json:"initialCooldown"`
	GlobalCooldown            int    `json:"globalCooldown"`
	MinPlayerLevel            int    `json:"minPlayerLevel"`
	HideEffects               bool   `json:"hideEffects"`
	Hidden                    bool   `json:"hidden"`
	PlayAnimation             bool   `json:"playAnimation"`
	StatesCriterion           string `json:"statesCriterion"`
	// effects vector custom subtype not implemented (3)
	// criticalEffect vector custom subtype not implemented (3)
	// previewZones vector custom subtype not implemented (4)
}

type SpellPair struct {
	Id            int `json:"id"`
	NameId        int `json:"nameId"`
	DescriptionId int `json:"descriptionId"`
	IconId        int `json:"iconId"`
}

type SpellScript struct {
	Id        int    `json:"id"`
	TypeId    int    `json:"typeId"`
	RawParams string `json:"rawParams"`
}

type SpellState struct {
	Id                   int    `json:"id"`
	NameId               int    `json:"nameId"`
	PreventsSpellCast    bool   `json:"preventsSpellCast"`
	PreventsFight        bool   `json:"preventsFight"`
	IsSilent             bool   `json:"isSilent"`
	CantBeMoved          bool   `json:"cantBeMoved"`
	CantBePushed         bool   `json:"cantBePushed"`
	CantDealDamage       bool   `json:"cantDealDamage"`
	Invulnerable         bool   `json:"invulnerable"`
	CantSwitchPosition   bool   `json:"cantSwitchPosition"`
	Incurable            bool   `json:"incurable"`
	EffectsIds           []int  `json:"effectsIds"`
	Icon                 string `json:"icon"`
	IconVisibilityMask   int    `json:"iconVisibilityMask"`
	InvulnerableMelee    bool   `json:"invulnerableMelee"`
	InvulnerableRange    bool   `json:"invulnerableRange"`
	CantTackle           bool   `json:"cantTackle"`
	CantBeTackled        bool   `json:"cantBeTackled"`
	DisplayTurnRemaining bool   `json:"displayTurnRemaining"`
}

type BoundScriptUsageData struct {
	Id             int    `json:"id"`
	ScriptId       int    `json:"scriptId"`
	SpellLevels    []uint `json:"spellLevels"`
	Criterion      string `json:"criterion"`
	CasterMask     string `json:"casterMask"`
	TargetMask     string `json:"targetMask"`
	TargetZone     string `json:"targetZone"`
	ActivationMask string `json:"activationMask"`
	ActivationZone string `json:"activationZone"`
	Random         int    `json:"random"`
	RandomGroup    int    `json:"randomGroup"`
}

type FinishMove struct {
	Id         int  `json:"id"`
	Duration   int  `json:"duration"`
	Free       bool `json:"free"`
	NameId     int  `json:"nameId"`
	Category   int  `json:"category"`
	SpellLevel int  `json:"spellLevel"`
}

type SpellConversion struct {
	OldSpellId int `json:"oldSpellId"`
	NewSpellId int `json:"newSpellId"`
}

type Spell struct {
	Id            int `json:"id"`
	NameId        int `json:"nameId"`
	DescriptionId int `json:"descriptionId"`
	TypeId        int `json:"typeId"`
	Order         int `json:"order"`
	// boundScriptUsageData vector custom subtype not implemented (1)
	// criticalHitBoundScriptUsageData vector custom subtype not implemented (1)
	IconId                 int    `json:"iconId"`
	SpellLevels            []uint `json:"spellLevels"`
	VerboseCast            bool   `json:"verboseCast"`
	DefaultPreviewZone     string `json:"defaultPreviewZone"`
	BypassSummoningLimit   bool   `json:"bypassSummoningLimit"`
	CanAlwaysTriggerSpells bool   `json:"canAlwaysTriggerSpells"`
	AdminName              string `json:"adminName"`
	HideCastConditions     bool   `json:"hideCastConditions"`
}
