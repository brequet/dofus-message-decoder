package types

type SmileyCategory struct {
	Id     int    `json:"id"`
	Order  int    `json:"order"`
	GfxId  string `json:"gfxId"`
	IsFake bool   `json:"isFake"`
}

type SmileyPack struct {
	Id      int    `json:"id"`
	NameId  int    `json:"nameId"`
	Order   int    `json:"order"`
	Smileys []uint `json:"smileys"`
}

type Smiley struct {
	Id          int      `json:"id"`
	Order       int      `json:"order"`
	GfxId       string   `json:"gfxId"`
	ForPlayers  bool     `json:"forPlayers"`
	Triggers    []string `json:"triggers"`
	ReferenceId int      `json:"referenceId"`
	CategoryId  int      `json:"categoryId"`
}

type CensoredWord struct {
	Id          int    `json:"id"`
	ListId      int    `json:"listId"`
	Language    string `json:"language"`
	Word        string `json:"word"`
	DeepLooking bool   `json:"deepLooking"`
}

type ChatChannel struct {
	Id            int    `json:"id"`
	NameId        int    `json:"nameId"`
	DescriptionId int    `json:"descriptionId"`
	Shortcut      string `json:"shortcut"`
	ShortcutKey   string `json:"shortcutKey"`
	IsPrivate     bool   `json:"isPrivate"`
	AllowObjects  bool   `json:"allowObjects"`
}

type Emoticon struct {
	Id               int    `json:"id"`
	NameId           int    `json:"nameId"`
	ShortcutId       int    `json:"shortcutId"`
	Order            uint   `json:"order"`
	AnimName         string `json:"animName"`
	Persistancy      bool   `json:"persistancy"`
	Eight_directions bool   `json:"eight_directions"`
	Aura             bool   `json:"aura"`
	Cooldown         uint   `json:"cooldown"`
	Duration         uint   `json:"duration"`
	Weight           uint   `json:"weight"`
	SpellLevelId     uint   `json:"spellLevelId"`
}

type InfoMessage struct {
	TypeId    int `json:"typeId"`
	MessageId int `json:"messageId"`
	TextId    int `json:"textId"`
}

type NamingRule struct {
	Id        int    `json:"id"`
	MinLength int    `json:"minLength"`
	MaxLength int    `json:"maxLength"`
	Regexp    string `json:"regexp"`
}
