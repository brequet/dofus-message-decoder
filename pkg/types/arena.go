package types

type ArenaLeagueReward struct {
	Id               int    `json:"id"`
	SeasonId         int    `json:"seasonId"`
	LeagueId         int    `json:"leagueId"`
	TitlesRewards    []uint `json:"titlesRewards"`
	EndSeasonRewards bool   `json:"endSeasonRewards"`
}

type ArenaLeague struct {
	Id              int    `json:"id"`
	NameId          int    `json:"nameId"`
	OrnamentId      int    `json:"ornamentId"`
	Icon            string `json:"icon"`
	Illus           string `json:"illus"`
	IsLastLeague    bool   `json:"isLastLeague"`
	LowRatingBound  int    `json:"lowRatingBound"`
	HighRatingBound int    `json:"highRatingBound"`
}
