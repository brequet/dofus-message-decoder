package types

type Challenge struct {
	Id                     int    `json:"id"`
	NameId                 int    `json:"nameId"`
	DescriptionId          int    `json:"descriptionId"`
	IncompatibleChallenges []uint `json:"incompatibleChallenges"`
	CategoryId             int    `json:"categoryId"`
	IconId                 uint   `json:"iconId"`
	CompletionCriterion    string `json:"completionCriterion"`
	ActivationCriterion    string `json:"activationCriterion"`
	TargetMonsterId        int    `json:"targetMonsterId"`
}
