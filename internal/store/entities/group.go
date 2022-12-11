package entities

type Group struct {
	Id          int
	Title       string
	Label       string
	Icon        string
	Jumpbtntext string
	HintType    int
}

type OpGroup struct {
	Group
	IsShow int
	Num    int
}
