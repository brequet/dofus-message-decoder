package types

type AbuseReasons struct {
	AbuseReasonId int `json:"abuseReasonId"`
	Mask          int `json:"mask"`
	ReasonTextId  int `json:"reasonTextId"`
}
