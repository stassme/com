package legalentities

type LegalEntities struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	BIK      string `json:"bik"`
	Bank     string `json:"bankName"`
	Address  string `json:"address"`
	KBill    string `json:"k_Bill"`
	RBill    string `json:"r_Bill"`
	Currency string `json:"currency"`
	Comment  string `json:"comment"`
}
