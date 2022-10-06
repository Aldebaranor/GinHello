package entity

func (Node) TableName() string {
	return "node"
}

type Node struct {
	NId      string   `json:"n_id" gorm:"column:n_id"`
	Layer    int      `json:"layer" gorm:"layer"`
	Tcam     int      `json:"tcam" gorm:"tcam"`
	Hash     int      `json:"hash" gorm:"hash"`
	Alu      int      `json:"alu" gorm:"alu"`
	Qualify  int      `json:"qualify" gorm:"qualify"`
	Write    []string `json:"write" gorm:"write"`
	Read     []string `json:"read" gorm:"read"`
	Children []*Node  `json:"children" gorm:"children"`
}
