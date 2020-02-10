package stripe

import (
	"encoding/json"

	"github.com/stripe/stripe-go/form"
)

// InvoiceLineType is the list of allowed values for the invoice line's type.
type InvoiceLineType string

// List of values that InvoiceLineType can take.
const (
	InvoiceLineTypeInvoiceItem  InvoiceLineType = "invoiceitem"
	InvoiceLineTypeSubscription InvoiceLineType = "subscription"
)

// InvoiceBillingReason is the reason why a given invoice was created
type InvoiceBillingReason string

// List of values that InvoiceBillingReason can take.
const (
	InvoiceBillingReasonManual                InvoiceBillingReason = "manual"
	InvoiceBillingReasonSubscription          InvoiceBillingReason = "subscription"
	InvoiceBillingReasonSubscriptionCreate    InvoiceBillingReason = "subscription_create"
	InvoiceBillingReasonSubscriptionCycle     InvoiceBillingReason = "subscription_cycle"
	InvoiceBillingReasonSubscriptionThreshold InvoiceBillingReason = "subscription_threshold"
	InvoiceBillingReasonSubscriptionUpdate    InvoiceBillingReason = "subscription_update"
	InvoiceBillingReasonUpcoming              InvoiceBillingReason = "upcoming"
)

// InvoiceStatus is the reason why a given invoice was created
type InvoiceStatus string

// List of values that InvoiceStatus can take.
const (
	InvoiceStatusDraft         InvoiceStatus = "draft"
	InvoiceStatusOpen          InvoiceStatus = "open"
	InvoiceStatusPaid          InvoiceStatus = "paid"
	InvoiceStatusUncollectible InvoiceStatus = "uncollectible"
	InvoiceStatusVoid          InvoiceStatus = "void"
)

// InvoiceCollectionMethod is the type of collection method for this invoice.
type InvoiceCollectionMethod string

// List of values that InvoiceCollectionMethod can take.
const (
	InvoiceCollectionMethodChargeAutomatically InvoiceCollectionMethod = "charge_automatically"
	InvoiceCollectionMethodSendInvoice         InvoiceCollectionMethod = "send_invoice"
)

// InvoiceUpcomingInvoiceItemPeriodParams represents the period associated with that invoice item
type InvoiceUpcomingInvoiceItemPeriodParams struct {
	End   *int64 `form:"end" json:"end"`
	Start *int64 `form:"start" json:"start"`
}

// InvoiceUpcomingInvoiceItemParams is the set of parameters that can be used when adding or modifying
// invoice items on an upcoming invoice.
// For more details see https://stripe.com/docs/api#upcoming_invoice-invoice_items.
type InvoiceUpcomingInvoiceItemParams struct {
	Amount            *int64                                  `form:"amount" json:"amount"`
	Currency          *string                                 `form:"currency" json:"currency"`
	Description       *string                                 `form:"description" json:"description"`
	Discountable      *bool                                   `form:"discountable" json:"discountable"`
	InvoiceItem       *string                                 `form:"invoiceitem" json:"invoiceitem"`
	Period            *InvoiceUpcomingInvoiceItemPeriodParams `form:"period" json:"period"`
	Quantity          *int64                                  `form:"quantity" json:"quantity"`
	Schedule          *string                                 `form:"schedule" json:"schedule"`
	TaxRates          []*string                               `form:"tax_rates" json:"tax_rates"`
	UnitAmount        *int64                                  `form:"unit_amount" json:"unit_amount"`
	UnitAmountDecimal *float64                                `form:"unit_amount_decimal,high_precision" json:"unit_amount_decimal,high_precision"`
}

// InvoiceCustomFieldParams represents the parameters associated with one custom field on an invoice.
type InvoiceCustomFieldParams struct {
	Name  *string `form:"name" json:"name"`
	Value *string `form:"value" json:"value"`
}

// InvoiceTransferDataParams is the set of parameters allowed for the transfer_data hash.
type InvoiceTransferDataParams struct {
	Destination *string `form:"destination" json:"destination"`
}

// InvoiceParams is the set of parameters that can be used when creating or updating an invoice.
// For more details see https://stripe.com/docs/api#create_invoice, https://stripe.com/docs/api#update_invoice.
type InvoiceParams struct {
	Params               `form:"*" json:"*"`
	AutoAdvance          *bool                       `form:"auto_advance" json:"auto_advance"`
	ApplicationFeeAmount *int64                      `form:"application_fee_amount" json:"application_fee_amount"`
	CollectionMethod     *string                     `form:"collection_method" json:"collection_method"`
	CustomFields         []*InvoiceCustomFieldParams `form:"custom_fields" json:"custom_fields"`
	Customer             *string                     `form:"customer" json:"customer"`
	DaysUntilDue         *int64                      `form:"days_until_due" json:"days_until_due"`
	DefaultPaymentMethod *string                     `form:"default_payment_method" json:"default_payment_method"`
	DefaultSource        *string                     `form:"default_source" json:"default_source"`
	DefaultTaxRates      []*string                   `form:"default_tax_rates" json:"default_tax_rates"`
	Description          *string                     `form:"description" json:"description"`
	DueDate              *int64                      `form:"due_date" json:"due_date"`
	Footer               *string                     `form:"footer" json:"footer"`
	Paid                 *bool                       `form:"paid" json:"paid"`
	StatementDescriptor  *string                     `form:"statement_descriptor" json:"statement_descriptor"`
	Subscription         *string                     `form:"subscription" json:"subscription"`
	TransferData         *InvoiceTransferDataParams  `form:"transfer_data" json:"transfer_data"`

	// These are all for exclusive use by GetNext.

	Coupon                                  *string                             `form:"coupon" json:"coupon"`
	InvoiceItems                            []*InvoiceUpcomingInvoiceItemParams `form:"invoice_items" json:"invoice_items"`
	SubscriptionBillingCycleAnchor          *int64                              `form:"subscription_billing_cycle_anchor" json:"subscription_billing_cycle_anchor"`
	SubscriptionBillingCycleAnchorNow       *bool                               `form:"-" json:"-"` // See custom AppendTo
	SubscriptionBillingCycleAnchorUnchanged *bool                               `form:"-" json:"-"` // See custom AppendTo
	SubscriptionCancelAt                    *int64                              `form:"subscription_cancel_at" json:"subscription_cancel_at"`
	SubscriptionCancelAtPeriodEnd           *bool                               `form:"subscription_cancel_at_period_end" json:"subscription_cancel_at_period_end"`
	SubscriptionCancelNow                   *bool                               `form:"subscription_cancel_now" json:"subscription_cancel_now"`
	SubscriptionDefaultTaxRates             []*string                           `form:"subscription_default_tax_rates" json:"subscription_default_tax_rates"`
	SubscriptionItems                       []*SubscriptionItemsParams          `form:"subscription_items" json:"subscription_items"`
	SubscriptionPlan                        *string                             `form:"subscription_plan" json:"subscription_plan"`
	SubscriptionProrate                     *bool                               `form:"subscription_prorate" json:"subscription_prorate"`
	SubscriptionProrationBehavior           *string                             `form:"subscription_proration_behavior" json:"subscription_proration_behavior"`
	SubscriptionProrationDate               *int64                              `form:"subscription_proration_date" json:"subscription_proration_date"`
	SubscriptionQuantity                    *int64                              `form:"subscription_quantity" json:"subscription_quantity"`
	SubscriptionTrialEnd                    *int64                              `form:"subscription_trial_end" json:"subscription_trial_end"`
	SubscriptionTrialFromPlan               *bool                               `form:"subscription_trial_from_plan" json:"subscription_trial_from_plan"`

	// This parameter is deprecated and we recommend that you use DefaultTaxRates instead.
	TaxPercent *float64 `form:"tax_percent" json:"tax_percent"`

	// This parameter is deprecated and we recommend that you use SubscriptionDefaultTaxRates instead.
	SubscriptionTaxPercent *float64 `form:"subscription_tax_percent" json:"subscription_tax_percent"`
}

// AppendTo implements custom encoding logic for InvoiceParams so that the special
// "now" value for subscription_billing_cycle_anchor can be implemented
// (they're otherwise timestamps rather than strings).
func (p *InvoiceParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.SubscriptionBillingCycleAnchorNow) {
		body.Add(form.FormatKey(append(keyParts, "subscription_billing_cycle_anchor")), "now")
	}

	if BoolValue(p.SubscriptionBillingCycleAnchorUnchanged) {
		body.Add(form.FormatKey(append(keyParts, "subscription_billing_cycle_anchor")), "unchanged")
	}
}

// InvoiceListParams is the set of parameters that can be used when listing invoices.
// For more details see https://stripe.com/docs/api#list_customer_invoices.
type InvoiceListParams struct {
	ListParams       `form:"*" json:"*"`
	CollectionMethod *string           `form:"collection_method" json:"collection_method"`
	Customer         *string           `form:"customer" json:"customer"`
	Created          *int64            `form:"created" json:"created"`
	CreatedRange     *RangeQueryParams `form:"created" json:"created"`
	DueDate          *int64            `form:"due_date" json:"due_date"`
	DueDateRange     *RangeQueryParams `form:"due_date" json:"due_date"`
	Status           *string           `form:"status" json:"status"`
	Subscription     *string           `form:"subscription" json:"subscription"`
}

// InvoiceLineListParams is the set of parameters that can be used when listing invoice line items.
// For more details see https://stripe.com/docs/api#invoice_lines.
type InvoiceLineListParams struct {
	ListParams `form:"*" json:"*"`

	Customer *string `form:"customer" json:"customer"`

	// ID is the invoice ID to list invoice lines for.
	ID *string `form:"-" json:"-"` // Goes in the URL

	Subscription *string `form:"subscription" json:"subscription"`
}

// InvoiceFinalizeParams is the set of parameters that can be used when finalizing invoices.
type InvoiceFinalizeParams struct {
	Params      `form:"*" json:"*"`
	AutoAdvance *bool `form:"auto_advance" json:"auto_advance"`
}

// InvoiceMarkUncollectibleParams is the set of parameters that can be used when marking
// invoices as uncollectible.
type InvoiceMarkUncollectibleParams struct {
	Params `form:"*" json:"*"`
}

// InvoicePayParams is the set of parameters that can be used when
// paying invoices. For more details, see:
// https://stripe.com/docs/api#pay_invoice.
type InvoicePayParams struct {
	Params        `form:"*" json:"*"`
	Forgive       *bool   `form:"forgive" json:"forgive"`
	OffSession    *bool   `form:"off_session" json:"off_session"`
	PaidOutOfBand *bool   `form:"paid_out_of_band" json:"paid_out_of_band"`
	PaymentMethod *string `form:"payment_method" json:"payment_method"`
	Source        *string `form:"source" json:"source"`
}

// InvoiceSendParams is the set of parameters that can be used when sending invoices.
type InvoiceSendParams struct {
	Params `form:"*" json:"*"`
}

// InvoiceVoidParams is the set of parameters that can be used when voiding invoices.
type InvoiceVoidParams struct {
	Params `form:"*" json:"*"`
}

// Invoice is the resource representing a Stripe invoice.
// For more details see https://stripe.com/docs/api#invoice_object.
type Invoice struct {
	AccountCountry               string                   `json:"account_country"`
	AccountName                  string                   `json:"account_name"`
	AmountDue                    int64                    `json:"amount_due"`
	AmountPaid                   int64                    `json:"amount_paid"`
	AmountRemaining              int64                    `json:"amount_remaining"`
	ApplicationFeeAmount         int64                    `json:"application_fee_amount"`
	AttemptCount                 int64                    `json:"attempt_count"`
	Attempted                    bool                     `json:"attempted"`
	AutoAdvance                  bool                     `json:"auto_advance"`
	BillingReason                InvoiceBillingReason     `json:"billing_reason"`
	Charge                       *Charge                  `json:"charge"`
	CollectionMethod             *InvoiceCollectionMethod `json:"collection_method"`
	Created                      int64                    `json:"created"`
	Currency                     Currency                 `json:"currency"`
	CustomFields                 []*InvoiceCustomField    `json:"custom_fields"`
	Customer                     *Customer                `json:"customer"`
	CustomerAddress              *Address                 `json:"customer_address"`
	CustomerEmail                string                   `json:"customer_email"`
	CustomerName                 *string                  `json:"customer_name"`
	CustomerPhone                *string                  `json:"customer_phone"`
	CustomerShipping             *CustomerShippingDetails `json:"customer_shipping"`
	CustomerTaxExempt            CustomerTaxExempt        `json:"customer_tax_exempt"`
	CustomerTaxIDs               []*InvoiceCustomerTaxID  `json:"customer_tax_ids"`
	DefaultPaymentMethod         *PaymentMethod           `json:"default_payment_method"`
	DefaultSource                *PaymentSource           `json:"default_source"`
	DefaultTaxRates              []*TaxRate               `json:"default_tax_rates"`
	Description                  string                   `json:"description"`
	Discount                     *Discount                `json:"discount"`
	DueDate                      int64                    `json:"due_date"`
	EndingBalance                int64                    `json:"ending_balance"`
	Footer                       string                   `json:"footer"`
	HostedInvoiceURL             string                   `json:"hosted_invoice_url"`
	ID                           string                   `json:"id"`
	InvoicePDF                   string                   `json:"invoice_pdf"`
	Lines                        *InvoiceLineList         `json:"lines"`
	Livemode                     bool                     `json:"livemode"`
	Metadata                     map[string]string        `json:"metadata"`
	NextPaymentAttempt           int64                    `json:"next_payment_attempt"`
	Number                       string                   `json:"number"`
	Paid                         bool                     `json:"paid"`
	PaymentIntent                *PaymentIntent           `json:"payment_intent"`
	PeriodEnd                    int64                    `json:"period_end"`
	PeriodStart                  int64                    `json:"period_start"`
	PostPaymentCreditNotesAmount int64                    `json:"post_payment_credit_notes_amount"`
	PrePaymentCreditNotesAmount  int64                    `json:"pre_payment_credit_notes_amount"`
	ReceiptNumber                string                   `json:"receipt_number"`
	StartingBalance              int64                    `json:"starting_balance"`
	StatementDescriptor          string                   `json:"statement_descriptor"`
	Status                       InvoiceStatus            `json:"status"`
	StatusTransitions            InvoiceStatusTransitions `json:"status_transitions"`
	Subscription                 string                   `json:"subscription"`
	SubscriptionProrationDate    int64                    `json:"subscription_proration_date"`
	Subtotal                     int64                    `json:"subtotal"`
	Tax                          int64                    `json:"tax"`
	ThreasholdReason             *InvoiceThresholdReason  `json:"threshold_reason"`
	Total                        int64                    `json:"total"`
	TotalTaxAmounts              []*InvoiceTaxAmount      `json:"total_tax_amounts"`
	TransferData                 *InvoiceTransferData     `json:"transfer_data"`
	WebhooksDeliveredAt          int64                    `json:"webhooks_delivered_at"`

	// This field is deprecated and we recommend that you use TaxRates instead.
	TaxPercent float64 `json:"tax_percent"`
}

// InvoiceCustomField is a structure representing a custom field on an invoice.
type InvoiceCustomField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// InvoiceCustomerTaxID is a structure representing a customer tax id on an invoice.
type InvoiceCustomerTaxID struct {
	Type  TaxIDType `json:"type"`
	Value string    `json:"value"`
}

// InvoiceTaxAmount is a structure representing one of the tax amounts on an invoice.
type InvoiceTaxAmount struct {
	Amount    int64    `json:"amount"`
	Inclusive bool     `json:"inclusive"`
	TaxRate   *TaxRate `json:"tax_rate"`
}

// InvoiceThresholdReason is a structure representing a reason for a billing threshold.
type InvoiceThresholdReason struct {
	AmountGTE   int64                               `json:"amount_gte"`
	ItemReasons []*InvoiceThresholdReasonItemReason `json:"item_reasons"`
}

// InvoiceThresholdReasonItemReason is a structure representing the line items that
// triggered an invoice.
type InvoiceThresholdReasonItemReason struct {
	LineItemIDs []string `json:"line_item_ids"`
	UsageGTE    int64    `json:"usage_gte"`
}

// InvoiceList is a list of invoices as retrieved from a list endpoint.
type InvoiceList struct {
	ListMeta
	Data []*Invoice `json:"data"`
}

// InvoiceLine is the resource representing a Stripe invoice line item.
// For more details see https://stripe.com/docs/api#invoice_line_item_object.
type InvoiceLine struct {
	Amount           int64               `json:"amount"`
	Currency         Currency            `json:"currency"`
	Description      string              `json:"description"`
	Discountable     bool                `json:"discountable"`
	ID               string              `json:"id"`
	InvoiceItem      string              `json:"invoice_item"`
	Livemode         bool                `json:"livemode"`
	Metadata         map[string]string   `json:"metadata"`
	Period           *Period             `json:"period"`
	Plan             *Plan               `json:"plan"`
	Proration        bool                `json:"proration"`
	Quantity         int64               `json:"quantity"`
	Subscription     string              `json:"subscription"`
	SubscriptionItem string              `json:"subscription_item"`
	TaxAmounts       []*InvoiceTaxAmount `json:"tax_amounts"`
	TaxRates         []*TaxRate          `json:"tax_rates"`
	Type             InvoiceLineType     `json:"type"`
	UnifiedProration bool                `json:"unified_proration"`
}

// InvoiceTransferData represents the information for the transfer_data associated with an invoice.
type InvoiceTransferData struct {
	Destination *Account `json:"destination"`
}

// Period is a structure representing a start and end dates.
type Period struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
}

// InvoiceLineList is a list object for invoice line items.
type InvoiceLineList struct {
	ListMeta
	Data []*InvoiceLine `json:"data"`
}

// InvoiceStatusTransitions are the timestamps at which the invoice status was updated.
type InvoiceStatusTransitions struct {
	FinalizedAt           int64 `json:"finalized_at"`
	MarkedUncollectibleAt int64 `json:"marked_uncollectible_at"`
	PaidAt                int64 `json:"paid_at"`
	VoidedAt              int64 `json:"voided_at"`
}

// UnmarshalJSON handles deserialization of an Invoice.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *Invoice) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type invoice Invoice
	var v invoice
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = Invoice(v)
	return nil
}
