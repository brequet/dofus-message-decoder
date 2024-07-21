package types

type MonsterSuperRace struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
}

type AnimFunMonsterData struct {
	AnimId     int    `json:"animId"`
	EntityId   int    `json:"entityId"`
	AnimName   string `json:"animName"`
	AnimWeight int    `json:"animWeight"`
}

type MonsterDropCoefficient struct {
	MonsterId       int     `json:"monsterId"`
	MonsterGrade    int     `json:"monsterGrade"`
	DropCoefficient float64 `json:"dropCoefficient"`
	Criteria        string  `json:"criteria"`
}

type CompanionCharacteristic struct {
	Id          int `json:"id"`
	CaracId     int `json:"caracId"`
	CompanionId int `json:"companionId"`
	Order       int `json:"order"`
	// statPerLevelRange vector subtype not implemented
}

type CompanionSpell struct {
	Id           int    `json:"id"`
	SpellId      int    `json:"spellId"`
	CompanionId  int    `json:"companionId"`
	GradeByLevel string `json:"gradeByLevel"`
}

type Companion struct {
	Id                   int    `json:"id"`
	NameId               int    `json:"nameId"`
	Look                 string `json:"look"`
	WebDisplay           bool   `json:"webDisplay"`
	DescriptionId        int    `json:"descriptionId"`
	StartingSpellLevelId int    `json:"startingSpellLevelId"`
	AssetId              int    `json:"assetId"`
	Characteristics      []uint `json:"characteristics"`
	Spells               []uint `json:"spells"`
	CreatureBoneId       int    `json:"creatureBoneId"`
	Visibility           string `json:"visibility"`
}

type MonsterRace struct {
	Id                          int    `json:"id"`
	SuperRaceId                 int    `json:"superRaceId"`
	NameId                      int    `json:"nameId"`
	AggressiveZoneSize          int    `json:"aggressiveZoneSize"`
	AggressiveLevelDiff         int    `json:"aggressiveLevelDiff"`
	AggressiveImmunityCriterion string `json:"aggressiveImmunityCriterion"`
	AggressiveAttackDelay       int    `json:"aggressiveAttackDelay"`
	Monsters                    []uint `json:"monsters"`
}

type MonsterGrade struct {
	Grade             int  `json:"grade"`
	MonsterId         int  `json:"monsterId"`
	Level             int  `json:"level"`
	LifePoints        int  `json:"lifePoints"`
	ActionPoints      int  `json:"actionPoints"`
	MovementPoints    int  `json:"movementPoints"`
	Vitality          int  `json:"vitality"`
	PaDodge           int  `json:"paDodge"`
	PmDodge           int  `json:"pmDodge"`
	Wisdom            int  `json:"wisdom"`
	EarthResistance   int  `json:"earthResistance"`
	AirResistance     int  `json:"airResistance"`
	FireResistance    int  `json:"fireResistance"`
	WaterResistance   int  `json:"waterResistance"`
	NeutralResistance int  `json:"neutralResistance"`
	GradeXp           int  `json:"gradeXp"`
	DamageReflect     int  `json:"damageReflect"`
	HiddenLevel       uint `json:"hiddenLevel"`
	Strength          int  `json:"strength"`
	Intelligence      int  `json:"intelligence"`
	Chance            int  `json:"chance"`
	Agility           int  `json:"agility"`
	BonusRange        int  `json:"bonusRange"`
	StartingSpellId   int  `json:"startingSpellId"`
	// bonusCharacteristics custom type not implemented (2)
}

type MonsterMiniBoss struct {
	Id                 int `json:"id"`
	MonsterReplacingId int `json:"monsterReplacingId"`
}

type Monster struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
	Race   int `json:"race"`
	// grades vector custom subtype not implemented (3)
	Look          string `json:"look"`
	UseSummonSlot bool   `json:"useSummonSlot"`
	UseBombSlot   bool   `json:"useBombSlot"`
	// animFunList vector custom subtype not implemented (1)
	IsBoss bool `json:"isBoss"`
	// drops vector custom subtype not implemented (5)
	// temporisDrops vector custom subtype not implemented (5)
	Subareas                    []uint   `json:"subareas"`
	Spells                      []uint   `json:"spells"`
	SpellGrades                 []string `json:"spellGrades"`
	FavoriteSubareaId           int      `json:"favoriteSubareaId"`
	IsMiniBoss                  bool     `json:"isMiniBoss"`
	IsQuestMonster              bool     `json:"isQuestMonster"`
	CorrespondingMiniBossId     int      `json:"correspondingMiniBossId"`
	SpeedAdjust                 int      `json:"speedAdjust"`
	CreatureBoneId              int      `json:"creatureBoneId"`
	FastAnimsFun                bool     `json:"fastAnimsFun"`
	CanPlay                     bool     `json:"canPlay"`
	CanTackle                   bool     `json:"canTackle"`
	CanBePushed                 bool     `json:"canBePushed"`
	CanSwitchPos                bool     `json:"canSwitchPos"`
	CanSwitchPosOnTarget        bool     `json:"canSwitchPosOnTarget"`
	CanBeCarried                bool     `json:"canBeCarried"`
	CanUsePortal                bool     `json:"canUsePortal"`
	IncompatibleChallenges      []uint   `json:"incompatibleChallenges"`
	UseRaceValues               bool     `json:"useRaceValues"`
	AggressiveZoneSize          int      `json:"aggressiveZoneSize"`
	AggressiveLevelDiff         int      `json:"aggressiveLevelDiff"`
	AggressiveImmunityCriterion string   `json:"aggressiveImmunityCriterion"`
	AggressiveAttackDelay       int      `json:"aggressiveAttackDelay"`
	ScaleGradeRef               int      `json:"scaleGradeRef"`
	// characRatios vector subtype not implemented
}

type MonsterBonusCharacteristics struct {
	LifePoints        int `json:"lifePoints"`
	Strength          int `json:"strength"`
	Wisdom            int `json:"wisdom"`
	Chance            int `json:"chance"`
	Agility           int `json:"agility"`
	Intelligence      int `json:"intelligence"`
	EarthResistance   int `json:"earthResistance"`
	FireResistance    int `json:"fireResistance"`
	WaterResistance   int `json:"waterResistance"`
	AirResistance     int `json:"airResistance"`
	NeutralResistance int `json:"neutralResistance"`
	TackleEvade       int `json:"tackleEvade"`
	TackleBlock       int `json:"tackleBlock"`
	BonusEarthDamage  int `json:"bonusEarthDamage"`
	BonusFireDamage   int `json:"bonusFireDamage"`
	BonusWaterDamage  int `json:"bonusWaterDamage"`
	BonusAirDamage    int `json:"bonusAirDamage"`
	APRemoval         int `json:"APRemoval"`
}

type MonsterDrop struct {
	DropId                  int     `json:"dropId"`
	MonsterId               int     `json:"monsterId"`
	ObjectId                int     `json:"objectId"`
	PercentDropForGrade1    float64 `json:"percentDropForGrade1"`
	PercentDropForGrade2    float64 `json:"percentDropForGrade2"`
	PercentDropForGrade3    float64 `json:"percentDropForGrade3"`
	PercentDropForGrade4    float64 `json:"percentDropForGrade4"`
	PercentDropForGrade5    float64 `json:"percentDropForGrade5"`
	Count                   int     `json:"count"`
	Criteria                string  `json:"criteria"`
	HasCriteria             bool    `json:"hasCriteria"`
	HiddenIfInvalidCriteria bool    `json:"hiddenIfInvalidCriteria"`
	// specificDropCoefficient vector custom subtype not implemented (4)
}
