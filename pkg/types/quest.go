package types

type AchievementCategory struct {
	Id                  int    `json:"id"`
	NameId              int    `json:"nameId"`
	ParentId            int    `json:"parentId"`
	Icon                string `json:"icon"`
	Order               int    `json:"order"`
	Color               string `json:"color"`
	AchievementIds      []uint `json:"achievementIds"`
	VisibilityCriterion string `json:"visibilityCriterion"`
}

type Achievement struct {
	Id            int   `json:"id"`
	NameId        int   `json:"nameId"`
	CategoryId    int   `json:"categoryId"`
	DescriptionId int   `json:"descriptionId"`
	IconId        int   `json:"iconId"`
	Points        int   `json:"points"`
	Level         int   `json:"level"`
	Order         int   `json:"order"`
	AccountLinked bool  `json:"accountLinked"`
	ObjectiveIds  []int `json:"objectiveIds"`
	RewardIds     []int `json:"rewardIds"`
}

type QuestCategory struct {
	Id       int    `json:"id"`
	NameId   int    `json:"nameId"`
	Order    int    `json:"order"`
	QuestIds []uint `json:"questIds"`
}

type QuestObjectiveType struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
}

type QuestStepRewards struct {
	Id                        int     `json:"id"`
	StepId                    int     `json:"stepId"`
	LevelMin                  int     `json:"levelMin"`
	LevelMax                  int     `json:"levelMax"`
	KamasRatio                float64 `json:"kamasRatio"`
	ExperienceRatio           float64 `json:"experienceRatio"`
	KamasScaleWithPlayerLevel bool    `json:"kamasScaleWithPlayerLevel"`
	// itemsReward vector subtype not implemented
	EmotesReward []uint `json:"emotesReward"`
	JobsReward   []uint `json:"jobsReward"`
	SpellsReward []uint `json:"spellsReward"`
	TitlesReward []uint `json:"titlesReward"`
}

type QuestStep struct {
	Id            int     `json:"id"`
	QuestId       int     `json:"questId"`
	NameId        int     `json:"nameId"`
	DescriptionId int     `json:"descriptionId"`
	DialogId      int     `json:"dialogId"`
	OptimalLevel  int     `json:"optimalLevel"`
	Duration      float64 `json:"duration"`
	ObjectiveIds  []uint  `json:"objectiveIds"`
	RewardsIds    []uint  `json:"rewardsIds"`
}

type Quest struct {
	Id             int    `json:"id"`
	NameId         int    `json:"nameId"`
	CategoryId     int    `json:"categoryId"`
	RepeatType     int    `json:"repeatType"`
	RepeatLimit    int    `json:"repeatLimit"`
	IsDungeonQuest bool   `json:"isDungeonQuest"`
	LevelMin       int    `json:"levelMin"`
	LevelMax       int    `json:"levelMax"`
	StepIds        []uint `json:"stepIds"`
	IsPartyQuest   bool   `json:"isPartyQuest"`
	StartCriterion string `json:"startCriterion"`
	Followable     bool   `json:"followable"`
}

type AchievementObjective struct {
	Id            int    `json:"id"`
	AchievementId int    `json:"achievementId"`
	Order         int    `json:"order"`
	NameId        int    `json:"nameId"`
	Criterion     string `json:"criterion"`
}

type AchievementReward struct {
	Id                        int     `json:"id"`
	AchievementId             int     `json:"achievementId"`
	Criteria                  string  `json:"criteria"`
	KamasRatio                float64 `json:"kamasRatio"`
	ExperienceRatio           float64 `json:"experienceRatio"`
	GuildPoints               uint    `json:"guildPoints"`
	KamasScaleWithPlayerLevel bool    `json:"kamasScaleWithPlayerLevel"`
	ItemsReward               []uint  `json:"itemsReward"`
	ItemsQuantityReward       []uint  `json:"itemsQuantityReward"`
	EmotesReward              []uint  `json:"emotesReward"`
	SpellsReward              []uint  `json:"spellsReward"`
	TitlesReward              []uint  `json:"titlesReward"`
	OrnamentsReward           []uint  `json:"ornamentsReward"`
}

type QuestObjective struct {
	Id       int `json:"id"`
	StepId   int `json:"stepId"`
	TypeId   int `json:"typeId"`
	DialogId int `json:"dialogId"`
	// parameters custom type not implemented (2)
	// coords custom type not implemented (1)
	MapId int `json:"mapId"`
}
