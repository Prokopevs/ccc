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
	OwnerId    int    `db:"ownerId"`
	Score      string `db:"score"`
	GasStorage int    `db:"gasStorage"`
	GasMining  string `db:"gasMining"`
	Protection int    `db:"protection"`
}
