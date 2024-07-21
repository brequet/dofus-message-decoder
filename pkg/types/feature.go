package types

type OptionalFeature struct {
	Id                             int    `json:"id"`
	Keyword                        string `json:"keyword"`
	IsClient                       bool   `json:"isClient"`
	IsServer                       bool   `json:"isServer"`
	IsActivationOnLaunch           bool   `json:"isActivationOnLaunch"`
	IsActivationOnServerConnection bool   `json:"isActivationOnServerConnection"`
	ActivationCriterions           string `json:"activationCriterions"`
}
