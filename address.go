package stripe

// AddressParams describes the common parameters for an Address.
type AddressParams struct {
	City       *string `form:"city" json:"city"`
	Country    *string `form:"country" json:"country"`
	Line1      *string `form:"line1" json:"line1"`
	Line2      *string `form:"line2" json:"line2"`
	PostalCode *string `form:"postal_code" json:"postal_code"`
	State      *string `form:"state" json:"state"`
}

// Address describes common properties for an Address hash.
type Address struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
}
