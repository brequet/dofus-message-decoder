package types

type Server struct {
	Id                    int      `json:"id"`
	NameId                int      `json:"nameId"`
	CommentId             int      `json:"commentId"`
	OpeningDate           float64  `json:"openingDate"`
	Language              string   `json:"language"`
	PopulationId          int      `json:"populationId"`
	GameTypeId            int      `json:"gameTypeId"`
	CommunityId           int      `json:"communityId"`
	RestrictedToLanguages []string `json:"restrictedToLanguages"`
	MonoAccount           bool     `json:"monoAccount"`
}

type ServerCommunity struct {
	Id                        int      `json:"id"`
	NameId                    int      `json:"nameId"`
	DefaultCountries          []string `json:"defaultCountries"`
	ShortId                   string   `json:"shortId"`
	SupportedLangIds          []int    `json:"supportedLangIds"`
	NamingRulePlayerNameId    int      `json:"namingRulePlayerNameId"`
	NamingRuleGuildNameId     int      `json:"namingRuleGuildNameId"`
	NamingRuleAllianceNameId  int      `json:"namingRuleAllianceNameId"`
	NamingRuleAllianceTagId   int      `json:"namingRuleAllianceTagId"`
	NamingRulePartyNameId     int      `json:"namingRulePartyNameId"`
	NamingRuleMountNameId     int      `json:"namingRuleMountNameId"`
	NamingRuleNameGeneratorId int      `json:"namingRuleNameGeneratorId"`
	NamingRuleAdminId         int      `json:"namingRuleAdminId"`
	NamingRuleModoId          int      `json:"namingRuleModoId"`
	NamingRulePresetNameId    int      `json:"namingRulePresetNameId"`
}

type ServerGameType struct {
	Id                 int  `json:"id"`
	SelectableByPlayer bool `json:"selectableByPlayer"`
	NameId             int  `json:"nameId"`
	RulesId            int  `json:"rulesId"`
	DescriptionId      int  `json:"descriptionId"`
}

type ServerLang struct {
	Id       int    `json:"id"`
	NameId   int    `json:"nameId"`
	LangCode string `json:"langCode"`
}

type ServerPopulation struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
	Weight int `json:"weight"`
}
