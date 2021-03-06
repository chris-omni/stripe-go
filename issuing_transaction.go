package stripe

import "encoding/json"

// IssuingTransactionType is the type of an issuing transaction.
type IssuingTransactionType string

// List of values that IssuingTransactionType can take.
const (
	IssuingTransactionTypeCapture        IssuingTransactionType = "capture"
	IssuingTransactionTypeCashWithdrawal IssuingTransactionType = "cash_withdrawal"
	IssuingTransactionTypeRefund         IssuingTransactionType = "refund"
	IssuingTransactionTypeRefundReversal IssuingTransactionType = "refund_reversal"
)

// IssuingTransactionParams is the set of parameters that can be used when creating or updating an issuing transaction.
type IssuingTransactionParams struct {
	Params `form:"*" json:"*"`
}

// IssuingTransactionListParams is the set of parameters that can be used when listing issuing transactions.
type IssuingTransactionListParams struct {
	ListParams   `form:"*" json:"*"`
	Card         *string           `form:"card" json:"card"`
	Cardholder   *string           `form:"cardholder" json:"cardholder"`
	Created      *int64            `form:"created" json:"created"`
	CreatedRange *RangeQueryParams `form:"created" json:"created"`
	Dispute      *string           `form:"dispute" json:"dispute"`
}

// IssuingTransaction is the resource representing a Stripe issuing transaction.
type IssuingTransaction struct {
	Amount             int64                  `json:"amount"`
	Authorization      *IssuingAuthorization  `json:"authorization"`
	BalanceTransaction *BalanceTransaction    `json:"balance_transaction"`
	Card               *IssuingCard           `json:"card"`
	Cardholder         *IssuingCardholder     `json:"cardholder"`
	Created            int64                  `json:"created"`
	Currency           Currency               `json:"currency"`
	Dispute            *IssuingDispute        `json:"dispute"`
	ID                 string                 `json:"id"`
	Livemode           bool                   `json:"livemode"`
	MerchantData       *IssuingMerchantData   `json:"merchant_data"`
	MerchantAmount     int64                  `json:"merchant_amount"`
	MerchantCurrency   Currency               `json:"merchant_currency"`
	Metadata           map[string]string      `json:"metadata"`
	Object             string                 `json:"object"`
	Type               IssuingTransactionType `json:"type"`
}

// IssuingTransactionList is a list of issuing transactions as retrieved from a list endpoint.
type IssuingTransactionList struct {
	ListMeta
	Data []*IssuingTransaction `json:"data"`
}

// UnmarshalJSON handles deserialization of an IssuingTransaction.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingTransaction) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type issuingTransaction IssuingTransaction
	var v issuingTransaction
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingTransaction(v)
	return nil
}
