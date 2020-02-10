package stripe

// TopupParams is the set of parameters that can be used when creating or updating a top-up.
// For more details see https://stripe.com/docs/api#create_topup and https://stripe.com/docs/api#update_topup.
type TopupParams struct {
	Params              `form:"*" json:"*"`
	Amount              *int64        `form:"amount" json:"amount"`
	Currency            *string       `form:"currency" json:"currency"`
	Description         *string       `form:"description" json:"description"`
	Source              *SourceParams `form:"*"` // SourceParams has custom encoding so brought to top level with "*" json:"*"` // SourceParams has custom encoding so brought to top level with "*"
	StatementDescriptor *string       `form:"statement_descriptor" json:"statement_descriptor"`
	TransferGroup       *string       `form:"transfer_group" json:"transfer_group"`
}

// SetSource adds valid sources to a TopupParams object,
// returning an error for unsupported sources.
func (p *TopupParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
}

// TopupListParams is the set of parameters that can be used when listing top-ups.
// For more details see https://stripe.com/docs/api#list_topups.
type TopupListParams struct {
	ListParams   `form:"*" json:"*"`
	Created      *int64            `form:"created" json:"created"`
	CreatedRange *RangeQueryParams `form:"created" json:"created"`
}

// TopupList is a list of top-ups as retrieved from a list endpoint.
type TopupList struct {
	ListMeta
	Data []*Topup `json:"data"`
}

// Topup is the resource representing a Stripe top-up.
// For more details see https://stripe.com/docs/api#topups.
type Topup struct {
	Amount                   int64               `json:"amount"`
	ArrivalDate              int64               `json:"arrival_date"`
	BalanceTransaction       *BalanceTransaction `json:"balance_transaction"`
	Created                  int64               `json:"created"`
	Currency                 Currency            `json:"currency"`
	Description              string              `json:"description"`
	ExpectedAvailabilityDate int64               `json:"expected_availability_date"`
	FailureCode              string              `json:"failure_code"`
	FailureMessage           string              `json:"failure_message"`
	ID                       string              `json:"id"`
	Livemode                 bool                `json:"livemode"`
	Source                   *PaymentSource      `json:"source"`
	StatementDescriptor      string              `json:"statement_descriptor"`
	Status                   string              `json:"status"`
	TransferGroup            string              `json:"transfer_group"`
}
