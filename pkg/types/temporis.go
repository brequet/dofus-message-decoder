package types

type AchievementProgress struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	SeasonId int    `json:"seasonId"`
}

type AchievementProgressStep struct {
	Id            int  `json:"id"`
	ProgressId    int  `json:"progressId"`
	Score         int  `json:"score"`
	IsCosmetic    bool `json:"isCosmetic"`
	AchievementId int  `json:"achievementId"`
	IsBuyable     bool `json:"isBuyable"`
}
