package types

type IncarnationLevel struct {
	Id            int `json:"id"`
	IncarnationId int `json:"incarnationId"`
	Level         int `json:"level"`
	RequiredXp    int `json:"requiredXp"`
}

type ItemSet struct {
	Id            int    `json:"id"`
	Items         []uint `json:"items"`
	NameId        int    `json:"nameId"`
	BonusIsSecret bool   `json:"bonusIsSecret"`
	// effects vector subtype not implemented
}

type ItemSuperType struct {
	Id                int   `json:"id"`
	PossiblePositions []int `json:"possiblePositions"`
}

type Item struct {
	Id                     int     `json:"id"`
	NameId                 int     `json:"nameId"`
	TypeId                 int     `json:"typeId"`
	DescriptionId          int     `json:"descriptionId"`
	IconId                 int     `json:"iconId"`
	Level                  int     `json:"level"`
	RealWeight             int     `json:"realWeight"`
	Cursed                 bool    `json:"cursed"`
	UseAnimationId         int     `json:"useAnimationId"`
	Usable                 bool    `json:"usable"`
	Targetable             bool    `json:"targetable"`
	Exchangeable           bool    `json:"exchangeable"`
	Price                  float64 `json:"price"`
	TwoHanded              bool    `json:"twoHanded"`
	Etheral                bool    `json:"etheral"`
	ItemSetId              int     `json:"itemSetId"`
	Criteria               string  `json:"criteria"`
	CriteriaTarget         string  `json:"criteriaTarget"`
	HideEffects            bool    `json:"hideEffects"`
	Enhanceable            bool    `json:"enhanceable"`
	NonUsableOnAnother     bool    `json:"nonUsableOnAnother"`
	AppearanceId           int     `json:"appearanceId"`
	SecretRecipe           bool    `json:"secretRecipe"`
	RecipeSlots            int     `json:"recipeSlots"`
	RecipeIds              []uint  `json:"recipeIds"`
	DropMonsterIds         []uint  `json:"dropMonsterIds"`
	DropTemporisMonsterIds []uint  `json:"dropTemporisMonsterIds"`
	ObjectIsDisplayOnWeb   bool    `json:"objectIsDisplayOnWeb"`
	BonusIsSecret          bool    `json:"bonusIsSecret"`
	// possibleEffects vector custom subtype not implemented (1)
	EvolutiveEffectIds        []uint  `json:"evolutiveEffectIds"`
	FavoriteSubAreas          []uint  `json:"favoriteSubAreas"`
	FavoriteSubAreasBonus     int     `json:"favoriteSubAreasBonus"`
	CraftXpRatio              int     `json:"craftXpRatio"`
	NeedUseConfirm            bool    `json:"needUseConfirm"`
	IsDestructible            bool    `json:"isDestructible"`
	IsSaleable                bool    `json:"isSaleable"`
	IsColorable               bool    `json:"isColorable"`
	IsLegendary               bool    `json:"isLegendary"`
	CraftVisible              string  `json:"craftVisible"`
	CraftConditional          string  `json:"craftConditional"`
	CraftFeasible             string  `json:"craftFeasible"`
	Visibility                string  `json:"visibility"`
	RecyclingNuggets          float64 `json:"recyclingNuggets"`
	FavoriteRecyclingSubareas []uint  `json:"favoriteRecyclingSubareas"`
	ContainerIds              []uint  `json:"containerIds"`
	// resourcesBySubarea vector subtype not implemented
	ImportantNoticeId     int     `json:"importantNoticeId"`
	ChangeVersion         string  `json:"changeVersion"`
	TooltipExpirationDate float64 `json:"tooltipExpirationDate"`
}

type LegendaryPowerCategory struct {
	Id                  int    `json:"id"`
	CategoryName        string `json:"categoryName"`
	CategoryOverridable bool   `json:"categoryOverridable"`
	CategorySpells      []int  `json:"categorySpells"`
}

type RandomDropItem struct {
	Id          int     `json:"id"`
	ItemId      int     `json:"itemId"`
	Probability float64 `json:"probability"`
	MinQuantity int     `json:"minQuantity"`
	MaxQuantity int     `json:"maxQuantity"`
}

type RandomDropGroup struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// randomDropItems vector custom subtype not implemented (1)
	DisplayContent bool `json:"displayContent"`
	DisplayChances bool `json:"displayChances"`
}

type EvolutiveItemType struct {
	Id                int     `json:"id"`
	MaxLevel          int     `json:"maxLevel"`
	ExperienceBoost   float64 `json:"experienceBoost"`
	ExperienceByLevel []int   `json:"experienceByLevel"`
}

type VeteranReward struct {
	Id              int  `json:"id"`
	RequiredSubDays uint `json:"requiredSubDays"`
	ItemGID         uint `json:"itemGID"`
	ItemQuantity    uint `json:"itemQuantity"`
}

type ItemType struct {
	Id               int    `json:"id"`
	NameId           int    `json:"nameId"`
	SuperTypeId      int    `json:"superTypeId"`
	CategoryId       int    `json:"categoryId"`
	IsInEncyclopedia bool   `json:"isInEncyclopedia"`
	Plural           bool   `json:"plural"`
	Gender           int    `json:"gender"`
	RawZone          string `json:"rawZone"`
	Mimickable       bool   `json:"mimickable"`
	CraftXpRatio     int    `json:"craftXpRatio"`
	EvolutiveTypeId  int    `json:"evolutiveTypeId"`
}

type Weapon struct {
	Etheral                    bool   `json:"etheral"`
	CraftConditional           string `json:"craftConditional"`
	Cursed                     bool   `json:"cursed"`
	AppearanceId               int    `json:"appearanceId"`
	Criteria                   string `json:"criteria"`
	CriticalFailureProbability int    `json:"criticalFailureProbability"`
	CraftVisible               string `json:"craftVisible"`
	NonUsableOnAnother         bool   `json:"nonUsableOnAnother"`
	Enhanceable                bool   `json:"enhanceable"`
	// possibleEffects vector custom subtype not implemented (1)
	CriticalHitBonus          int     `json:"criticalHitBonus"`
	ImportantNoticeId         int     `json:"importantNoticeId"`
	ChangeVersion             string  `json:"changeVersion"`
	RecipeIds                 []uint  `json:"recipeIds"`
	Price                     float64 `json:"price"`
	IsLegendary               bool    `json:"isLegendary"`
	Id                        int     `json:"id"`
	RecyclingNuggets          float64 `json:"recyclingNuggets"`
	DropTemporisMonsterIds    []uint  `json:"dropTemporisMonsterIds"`
	IconId                    int     `json:"iconId"`
	Visibility                string  `json:"visibility"`
	Level                     int     `json:"level"`
	NeedUseConfirm            bool    `json:"needUseConfirm"`
	FavoriteSubAreasBonus     int     `json:"favoriteSubAreasBonus"`
	CraftXpRatio              int     `json:"craftXpRatio"`
	FavoriteRecyclingSubareas []uint  `json:"favoriteRecyclingSubareas"`
	UseAnimationId            int     `json:"useAnimationId"`
	NameId                    int     `json:"nameId"`
	TypeId                    int     `json:"typeId"`
	RecipeSlots               int     `json:"recipeSlots"`
	MinRange                  int     `json:"minRange"`
	Exchangeable              bool    `json:"exchangeable"`
	CriticalHitProbability    int     `json:"criticalHitProbability"`
	IsSaleable                bool    `json:"isSaleable"`
	Range                     int     `json:"range"`
	CastInLine                bool    `json:"castInLine"`
	ApCost                    int     `json:"apCost"`
	CastInDiagonal            bool    `json:"castInDiagonal"`
	Usable                    bool    `json:"usable"`
	ContainerIds              []uint  `json:"containerIds"`
	ObjectIsDisplayOnWeb      bool    `json:"objectIsDisplayOnWeb"`
	DescriptionId             int     `json:"descriptionId"`
	TwoHanded                 bool    `json:"twoHanded"`
	EvolutiveEffectIds        []uint  `json:"evolutiveEffectIds"`
	CriteriaTarget            string  `json:"criteriaTarget"`
	FavoriteSubAreas          []uint  `json:"favoriteSubAreas"`
	SecretRecipe              bool    `json:"secretRecipe"`
	HideEffects               bool    `json:"hideEffects"`
	CraftFeasible             string  `json:"craftFeasible"`
	CastTestLos               bool    `json:"castTestLos"`
	IsColorable               bool    `json:"isColorable"`
	// resourcesBySubarea vector subtype not implemented
	RealWeight            int     `json:"realWeight"`
	Targetable            bool    `json:"targetable"`
	IsDestructible        bool    `json:"isDestructible"`
	ItemSetId             int     `json:"itemSetId"`
	DropMonsterIds        []uint  `json:"dropMonsterIds"`
	TooltipExpirationDate float64 `json:"tooltipExpirationDate"`
	BonusIsSecret         bool    `json:"bonusIsSecret"`
	MaxCastPerTurn        int     `json:"maxCastPerTurn"`
}

type PresetIcon struct {
	Id    int `json:"id"`
	Order int `json:"order"`
}

type Incarnation struct {
	Id           int    `json:"id"`
	LookMale     string `json:"lookMale"`
	LookFemale   string `json:"lookFemale"`
	MaleBoneId   int    `json:"maleBoneId"`
	FemaleBoneId int    `json:"femaleBoneId"`
}
