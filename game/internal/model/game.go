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
	OwnerId    int `db:"ownerid"`
	Score      int `db:"score"`
	GasStorage int `db:"gasstorage"`
	GasMining  int `db:"gasmining"`
	Protection int `db:"protection"`
}
