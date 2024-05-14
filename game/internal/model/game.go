package model

type Score struct {
	Id    int
	Score int
}

type MultipUpdate struct {
	Id       int
	NameType string
}

type Game struct {
	OwnerId    int `db:"ownerid" json:"ownerId"`
	Score      int `db:"score" json:"score"`
	GasStorage int `db:"gasstorage" json:"gasStorage"`
	GasMining  int `db:"gasmining" json:"gasMining"`
	Protection int `db:"protection" json:"protection"`
}
