package types

type TransformData struct {
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	ScaleX       float64 `json:"scaleX"`
	ScaleY       float64 `json:"scaleY"`
	Rotation     int     `json:"rotation"`
	OriginalClip string  `json:"originalClip"`
	OverrideClip string  `json:"overrideClip"`
}
