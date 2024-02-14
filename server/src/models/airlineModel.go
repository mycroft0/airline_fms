package models

type Charge struct {
	Id        int
	Name      string
	Iata_code string
	Pima      string
	Pima_ccn  string
	Address   string
	Edi       bool
	Gross_wt  bool
	// Limit_amount    float32
	Agent string `gorm:"column:name"`
}
