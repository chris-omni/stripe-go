package stripe

import (
	"encoding/json"
)

// CheckoutSessionSubmitType is the list of allowed values for the `submit_type`
// of a Session.
type CheckoutSessionSubmitType string

// List of values that CheckoutSessionSubmitType can take.
const (
	CheckoutSessionSubmitTypeAuto   CheckoutSessionSubmitType = "auto"
	CheckoutSessionSubmitTypeBook   CheckoutSessionSubmitType = "book"
	CheckoutSessionSubmitTypeDonate CheckoutSessionSubmitType = "donate"
	CheckoutSessionSubmitTypePay    CheckoutSessionSubmitType = "pay"
)

// CheckoutSessionDisplayItemType is the list of allowed values for the display item type.
type CheckoutSessionDisplayItemType string

// List of values that CheckoutSessionDisplayItemType can take.
const (
	CheckoutSessionDisplayItemTypeCustom CheckoutSessionDisplayItemType = "custom"
	CheckoutSessionDisplayItemTypePlan   CheckoutSessionDisplayItemType = "plan"
	CheckoutSessionDisplayItemTypeSKU    CheckoutSessionDisplayItemType = "sku"
)

// CheckoutSessionMode is the list of allowed values for the mode on a Session.
type CheckoutSessionMode string

// List of values that CheckoutSessionMode can take.
const (
	CheckoutSessionModePayment      CheckoutSessionMode = "payment"
	CheckoutSessionModeSetup        CheckoutSessionMode = "setup"
	CheckoutSessionModeSubscription CheckoutSessionMode = "subscription"
)

// CheckoutSessionLineItemParams is the set of parameters allowed for a line item
// on a checkout session.
type CheckoutSessionLineItemParams struct {
	Amount      *int64    `form:"amount" json:"amount"`
	Currency    *string   `form:"currency" json:"currency"`
	Description *string   `form:"description" json:"description"`
	Images      []*string `form:"images" json:"images"`
	Name        *string   `form:"name" json:"name"`
	Quantity    *int64    `form:"quantity" json:"quantity"`
}

// CheckoutSessionPaymentIntentDataTransferDataParams is the set of parameters allowed for the
// transfer_data hash.
type CheckoutSessionPaymentIntentDataTransferDataParams struct {
	Destination *string `form:"destination" json:"destination"`
}

// CheckoutSessionPaymentIntentDataParams is the set of parameters allowed for the
// payment intent creation on a checkout session.
type CheckoutSessionPaymentIntentDataParams struct {
	Params               `form:"*" json:"*"`
	ApplicationFeeAmount *int64                                              `form:"application_fee_amount" json:"application_fee_amount"`
	CaptureMethod        *string                                             `form:"capture_method" json:"capture_method"`
	Description          *string                                             `form:"description" json:"description"`
	OnBehalfOf           *string                                             `form:"on_behalf_of" json:"on_behalf_of"`
	ReceiptEmail         *string                                             `form:"receipt_email" json:"receipt_email"`
	SetupFutureUsage     *string                                             `form:"setup_future_usage" json:"setup_future_usage"`
	Shipping             *ShippingDetailsParams                              `form:"shipping" json:"shipping"`
	StatementDescriptor  *string                                             `form:"statement_descriptor" json:"statement_descriptor"`
	TransferData         *CheckoutSessionPaymentIntentDataTransferDataParams `form:"transfer_data" json:"transfer_data"`
}

// CheckoutSessionSetupIntentDataParams is the set of parameters allowed for the setup intent
// creation on a checkout session.
type CheckoutSessionSetupIntentDataParams struct {
	Params      `form:"*" json:"*"`
	Description *string `form:"description" json:"description"`
	OnBehalfOf  *string `form:"on_behalf_of" json:"on_behalf_of"`
}

// CheckoutSessionSubscriptionDataItemsParams is the set of parameters allowed for one item on a
// checkout session associated with a subscription.
type CheckoutSessionSubscriptionDataItemsParams struct {
	Plan     *string `form:"plan" json:"plan"`
	Quantity *int64  `form:"quantity" json:"quantity"`
}

// CheckoutSessionSubscriptionDataParams is the set of parameters allowed for the subscription
// creation on a checkout session.
type CheckoutSessionSubscriptionDataParams struct {
	Params                `form:"*" json:"*"`
	ApplicationFeePercent *float64                                      `form:"application_fee_percent" json:"application_fee_percent"`
	Items                 []*CheckoutSessionSubscriptionDataItemsParams `form:"items" json:"items"`
	TrialEnd              *int64                                        `form:"trial_end" json:"trial_end"`
	TrialFromPlan         *bool                                         `form:"trial_from_plan" json:"trial_from_plan"`
	TrialPeriodDays       *int64                                        `form:"trial_period_days" json:"trial_period_days"`
}

// CheckoutSessionParams is the set of parameters that can be used when creating
// a checkout session.
// For more details see https://stripe.com/docs/api/checkout/sessions/create
type CheckoutSessionParams struct {
	Params                   `form:"*" json:"*"`
	BillingAddressCollection *string                                 `form:"billing_address_collection" json:"billing_address_collection"`
	CancelURL                *string                                 `form:"cancel_url" json:"cancel_url"`
	ClientReferenceID        *string                                 `form:"client_reference_id" json:"client_reference_id"`
	Customer                 *string                                 `form:"customer" json:"customer"`
	CustomerEmail            *string                                 `form:"customer_email" json:"customer_email"`
	LineItems                []*CheckoutSessionLineItemParams        `form:"line_items" json:"line_items"`
	Locale                   *string                                 `form:"locale" json:"locale"`
	Mode                     *string                                 `form:"mode" json:"mode"`
	PaymentIntentData        *CheckoutSessionPaymentIntentDataParams `form:"payment_intent_data" json:"payment_intent_data"`
	PaymentMethodTypes       []*string                               `form:"payment_method_types" json:"payment_method_types"`
	SetupIntentData          *CheckoutSessionSetupIntentDataParams   `form:"setup_intent_data" json:"setup_intent_data"`
	SubscriptionData         *CheckoutSessionSubscriptionDataParams  `form:"subscription_data" json:"subscription_data"`
	SubmitType               *string                                 `form:"submit_type" json:"submit_type"`
	SuccessURL               *string                                 `form:"success_url" json:"success_url"`
}

// CheckoutSessionDisplayItemCustom represents an item of type custom in a checkout session
type CheckoutSessionDisplayItemCustom struct {
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Name        string   `json:"name"`
}

// CheckoutSessionDisplayItem represents one of the items in a checkout session.
type CheckoutSessionDisplayItem struct {
	Amount   int64                             `json:"amount"`
	Currency Currency                          `json:"currency"`
	Custom   *CheckoutSessionDisplayItemCustom `json:"custom"`
	Quantity int64                             `json:"quantity"`
	Plan     *Plan                             `json:"plan"`
	SKU      *SKU                              `json:"sku"`
	Type     CheckoutSessionDisplayItemType    `json:"type"`
}

// CheckoutSession is the resource representing a Stripe checkout session.
// For more details see https://stripe.com/docs/api/checkout/sessions/object
type CheckoutSession struct {
	CancelURL          string                        `json:"cancel_url"`
	ClientReferenceID  string                        `json:"client_reference_id"`
	Customer           *Customer                     `json:"customer"`
	CustomerEmail      string                        `json:"customer_email"`
	Deleted            bool                          `json:"deleted"`
	DisplayItems       []*CheckoutSessionDisplayItem `json:"display_items"`
	ID                 string                        `json:"id"`
	Livemode           bool                          `json:"livemode"`
	Locale             string                        `json:"locale"`
	Metadata           map[string]string             `json:"metadata"`
	Mode               CheckoutSessionMode           `json:"mode"`
	Object             string                        `json:"object"`
	PaymentIntent      *PaymentIntent                `json:"payment_intent"`
	PaymentMethodTypes []string                      `json:"payment_method_types"`
	SetupIntent        *SetupIntent                  `json:"setup_intent"`
	Subscription       *Subscription                 `json:"subscription"`
	SubmitType         CheckoutSessionSubmitType     `json:"submit_type"`
	SuccessURL         string                        `json:"success_url"`
}

// UnmarshalJSON handles deserialization of a checkout session.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *CheckoutSession) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type session CheckoutSession
	var v session
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = CheckoutSession(v)
	return nil
}
