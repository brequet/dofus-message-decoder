package types

type MountBone struct {
	Id int `json:"id"`
}

type MountFamily struct {
	Id      int    `json:"id"`
	NameId  int    `json:"nameId"`
	HeadUri string `json:"headUri"`
}

type Mount struct {
	Id            int    `json:"id"`
	FamilyId      int    `json:"familyId"`
	NameId        int    `json:"nameId"`
	Look          string `json:"look"`
	CertificateId int    `json:"certificateId"`
	// effects vector custom subtype not implemented (1)
}

type RideFood struct {
	Gid      int `json:"gid"`
	TypeId   int `json:"typeId"`
	FamilyId int `json:"familyId"`
}

type MountBehavior struct {
	Id            int `json:"id"`
	NameId        int `json:"nameId"`
	DescriptionId int `json:"descriptionId"`
}
