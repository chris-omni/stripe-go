package stripe

import (
	"encoding/json"

	"github.com/stripe/stripe-go/form"
)

// SourceCodeVerificationFlowStatus represents the possible statuses of a code verification flow.
type SourceCodeVerificationFlowStatus string

// List of values that SourceCodeVerificationFlowStatus can take.
const (
	SourceCodeVerificationFlowStatusFailed    SourceCodeVerificationFlowStatus = "failed"
	SourceCodeVerificationFlowStatusPending   SourceCodeVerificationFlowStatus = "pending"
	SourceCodeVerificationFlowStatusSucceeded SourceCodeVerificationFlowStatus = "succeeded"
)

// SourceFlow represents the possible flows of a source object.
type SourceFlow string

// List of values that SourceFlow can take.
const (
	SourceFlowCodeVerification SourceFlow = "code_verification"
	SourceFlowNone             SourceFlow = "none"
	SourceFlowReceiver         SourceFlow = "receiver"
	SourceFlowRedirect         SourceFlow = "redirect"
)

// SourceMandateAcceptanceStatus represents the possible failure reasons of a redirect flow.
type SourceMandateAcceptanceStatus string

// List of values that SourceMandateAcceptanceStatus can take.
const (
	SourceMandateAcceptanceStatusAccepted SourceMandateAcceptanceStatus = "accepted"
	SourceMandateAcceptanceStatusRefused  SourceMandateAcceptanceStatus = "refused"
)

// SourceMandateNotificationMethod represents the possible methods of notification for a mandate.
type SourceMandateNotificationMethod string

// List of values that SourceMandateNotificationMethod can take.
const (
	SourceMandateNotificationMethodEmail  SourceMandateNotificationMethod = "email"
	SourceMandateNotificationMethodManual SourceMandateNotificationMethod = "manual"
	SourceMandateNotificationMethodNone   SourceMandateNotificationMethod = "none"
)

// SourceSourceOrderItemType describes the type of source order items on source
// orders for sources.
type SourceSourceOrderItemType string

// The list of possible values for source order item types.
const (
	SourceSourceOrderItemTypeDiscount SourceSourceOrderItemType = "discount"
	SourceSourceOrderItemTypeSKU      SourceSourceOrderItemType = "sku"
	SourceSourceOrderItemTypeShipping SourceSourceOrderItemType = "shipping"
	SourceSourceOrderItemTypeTax      SourceSourceOrderItemType = "tax"
)

// SourceRedirectFlowFailureReason represents the possible failure reasons of a redirect flow.
type SourceRedirectFlowFailureReason string

// List of values that SourceRedirectFlowFailureReason can take.
const (
	SourceRedirectFlowFailureReasonDeclined        SourceRedirectFlowFailureReason = "declined"
	SourceRedirectFlowFailureReasonProcessingError SourceRedirectFlowFailureReason = "processing_error"
	SourceRedirectFlowFailureReasonUserAbort       SourceRedirectFlowFailureReason = "user_abort"
)

// SourceRedirectFlowStatus represents the possible statuses of a redirect flow.
type SourceRedirectFlowStatus string

// List of values that SourceRedirectFlowStatus can take.
const (
	SourceRedirectFlowStatusFailed      SourceRedirectFlowStatus = "failed"
	SourceRedirectFlowStatusNotRequired SourceRedirectFlowStatus = "not_required"
	SourceRedirectFlowStatusPending     SourceRedirectFlowStatus = "pending"
	SourceRedirectFlowStatusSucceeded   SourceRedirectFlowStatus = "succeeded"
)

// SourceRefundAttributesMethod are the possible method to retrieve a receiver's refund attributes.
type SourceRefundAttributesMethod string

// List of values that SourceRefundAttributesMethod can take.
const (
	SourceRefundAttributesMethodEmail  SourceRefundAttributesMethod = "email"
	SourceRefundAttributesMethodManual SourceRefundAttributesMethod = "manual"
)

// SourceRefundAttributesStatus are the possible status of a receiver's refund attributes.
type SourceRefundAttributesStatus string

// List of values that SourceRefundAttributesStatus can take.
const (
	SourceRefundAttributesStatusAvailable SourceRefundAttributesStatus = "available"
	SourceRefundAttributesStatusMissing   SourceRefundAttributesStatus = "missing"
	SourceRefundAttributesStatusRequested SourceRefundAttributesStatus = "requested"
)

// SourceStatus represents the possible statuses of a source object.
type SourceStatus string

// List of values that SourceStatus can take.
const (
	SourceStatusCanceled   SourceStatus = "canceled"
	SourceStatusChargeable SourceStatus = "chargeable"
	SourceStatusConsumed   SourceStatus = "consumed"
	SourceStatusFailed     SourceStatus = "failed"
	SourceStatusPending    SourceStatus = "pending"
)

// SourceUsage represents the possible usages of a source object.
type SourceUsage string

// List of values that SourceUsage can take.
const (
	SourceUsageReusable  SourceUsage = "reusable"
	SourceUsageSingleUse SourceUsage = "single_use"
)

// SourceOwnerParams is the set of parameters allowed for the owner hash on
// source creation or update.
type SourceOwnerParams struct {
	Address *AddressParams `form:"address" json:"address"`
	Email   *string        `form:"email" json:"email"`
	Name    *string        `form:"name" json:"name"`
	Phone   *string        `form:"phone" json:"phone"`
}

// RedirectParams is the set of parameters allowed for the redirect hash on
// source creation or update.
type RedirectParams struct {
	ReturnURL *string `form:"return_url" json:"return_url"`
}

// SourceOrderItemsParams is the set of parameters allowed for the items on a
// source order for a source.
type SourceOrderItemsParams struct {
	Amount      *int64  `form:"amount" json:"amount"`
	Currency    *string `form:"currency" json:"currency"`
	Description *string `form:"description" json:"description"`
	Parent      *string `form:"parent" json:"parent"`
	Quantity    *int64  `form:"quantity" json:"quantity"`
	Type        *string `form:"type" json:"type"`
}

// SourceOrderParams is the set of parameters allowed for the source order of a
// source.
type SourceOrderParams struct {
	Items    []*SourceOrderItemsParams `form:"items" json:"items"`
	Shipping *ShippingDetailsParams    `form:"shipping" json:"shipping"`
}

// SourceObjectParams is the set of parameters allowed on source creation or update.
type SourceObjectParams struct {
	Params              `form:"*" json:"*"`
	Amount              *int64                `form:"amount" json:"amount"`
	Currency            *string               `form:"currency" json:"currency"`
	Customer            *string               `form:"customer" json:"customer"`
	Flow                *string               `form:"flow" json:"flow"`
	Mandate             *SourceMandateParams  `form:"mandate" json:"mandate"`
	OriginalSource      *string               `form:"original_source" json:"original_source"`
	Owner               *SourceOwnerParams    `form:"owner" json:"owner"`
	Receiver            *SourceReceiverParams `form:"receiver" json:"receiver"`
	Redirect            *RedirectParams       `form:"redirect" json:"redirect"`
	SourceOrder         *SourceOrderParams    `form:"source_order" json:"source_order"`
	StatementDescriptor *string               `form:"statement_descriptor" json:"statement_descriptor"`
	Token               *string               `form:"token" json:"token"`
	Type                *string               `form:"type" json:"type"`
	TypeData            map[string]string     `form:"-" json:"-"`
	Usage               *string               `form:"usage" json:"usage"`
}

// SourceMandateAcceptanceParams describes the set of parameters allowed for the `acceptance`
// hash on source creation or update.
type SourceMandateAcceptanceParams struct {
	Date      *int64                                `form:"date" json:"date"`
	IP        *string                               `form:"ip" json:"ip"`
	Offline   *SourceMandateAcceptanceOfflineParams `form:"offline" json:"offline"`
	Online    *SourceMandateAcceptanceOnlineParams  `form:"online" json:"online"`
	Status    *string                               `form:"status" json:"status"`
	Type      *string                               `form:"type" json:"type"`
	UserAgent *string                               `form:"user_agent" json:"user_agent"`
}

// SourceMandateAcceptanceOnlineParams describes the set of parameters for online accepted mandate
type SourceMandateAcceptanceOnlineParams struct {
	Date      *int64  `form:"date" json:"date"`
	IP        *string `form:"ip" json:"ip"`
	UserAgent *string `form:"user_agent" json:"user_agent"`
}

// SourceMandateAcceptanceOfflineParams describes the set of parameters for offline accepted mandate
type SourceMandateAcceptanceOfflineParams struct {
	ContactEmail *string `form:"contact_email" json:"contact_email"`
}

// SourceMandateParams describes the set of parameters allowed for the `mandate` hash on
// source creation or update.
type SourceMandateParams struct {
	Amount             *int64                         `form:"amount" json:"amount"`
	Acceptance         *SourceMandateAcceptanceParams `form:"acceptance" json:"acceptance"`
	Currency           *string                        `form:"currency" json:"currency"`
	Interval           *string                        `form:"interval" json:"interval"`
	NotificationMethod *string                        `form:"notification_method" json:"notification_method"`
}

// SourceReceiverParams is the set of parameters allowed for the `receiver` hash on
// source creation or update.
type SourceReceiverParams struct {
	RefundAttributesMethod *string `form:"refund_attributes_method" json:"refund_attributes_method"`
}

// SourceObjectDetachParams is the set of parameters that can be used when detaching
// a source from a customer.
type SourceObjectDetachParams struct {
	Params   `form:"*" json:"*"`
	Customer *string `form:"-" json:"-"`
}

// SourceOwner describes the owner hash on a source.
type SourceOwner struct {
	Address         *Address `json:"address,omitempty"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	Phone           string   `json:"phone"`
	VerifiedAddress *Address `json:"verified_address,omitempty"`
	VerifiedEmail   string   `json:"verified_email"`
	VerifiedName    string   `json:"verified_name"`
	VerifiedPhone   string   `json:"verified_phone"`
}

// RedirectFlow informs of the state of a redirect authentication flow.
type RedirectFlow struct {
	FailureReason SourceRedirectFlowFailureReason `json:"failure_reason"`
	ReturnURL     string                          `json:"return_url"`
	Status        SourceRedirectFlowStatus        `json:"status"`
	URL           string                          `json:"url"`
}

// ReceiverFlow informs of the state of a receiver authentication flow.
type ReceiverFlow struct {
	Address                string                       `json:"address"`
	AmountCharged          int64                        `json:"amount_charged"`
	AmountReceived         int64                        `json:"amount_received"`
	AmountReturned         int64                        `json:"amount_returned"`
	RefundAttributesMethod SourceRefundAttributesMethod `json:"refund_attributes_method"`
	RefundAttributesStatus SourceRefundAttributesStatus `json:"refund_attributes_status"`
}

// CodeVerificationFlow informs of the state of a verification authentication flow.
type CodeVerificationFlow struct {
	AttemptsRemaining int64                            `json:"attempts_remaining"`
	Status            SourceCodeVerificationFlowStatus `json:"status"`
}

// SourceMandateAcceptance describes a source mandate acceptance state.
type SourceMandateAcceptance struct {
	Date      int64                         `json:"date"`
	IP        string                        `json:"ip"`
	Status    SourceMandateAcceptanceStatus `json:"status"`
	UserAgent string                        `json:"user_agent"`
}

// SourceMandate describes a source mandate.
type SourceMandate struct {
	Acceptance         *SourceMandateAcceptance        `json:"acceptance"`
	NotificationMethod SourceMandateNotificationMethod `json:"notification_method"`
	Reference          string                          `json:"reference"`
	URL                string                          `json:"url"`
}

// SourceSourceOrderItems describes the items on source orders for sources.
type SourceSourceOrderItems struct {
	Amount      int64                     `json:"amount"`
	Currency    Currency                  `json:"currency"`
	Description string                    `json:"description"`
	Quantity    int64                     `json:"quantity"`
	Type        SourceSourceOrderItemType `json:"type"`
}

// SourceSourceOrder describes a source order for a source.
type SourceSourceOrder struct {
	Amount   int64                     `json:"amount"`
	Currency Currency                  `json:"currency"`
	Email    string                    `json:"email"`
	Items    *[]SourceSourceOrderItems `json:"items"`
	Shipping *ShippingDetails          `json:"shipping"`
}

// Source is the resource representing a Source.
// For more details see https://stripe.com/docs/api#sources.
type Source struct {
	Amount              int64                 `json:"amount"`
	ClientSecret        string                `json:"client_secret"`
	CodeVerification    *CodeVerificationFlow `json:"code_verification,omitempty"`
	Created             int64                 `json:"created"`
	Currency            Currency              `json:"currency"`
	Customer            string                `json:"customer"`
	Flow                SourceFlow            `json:"flow"`
	ID                  string                `json:"id"`
	Livemode            bool                  `json:"livemode"`
	Mandate             *SourceMandate        `json:"mandate"`
	Metadata            map[string]string     `json:"metadata"`
	Owner               *SourceOwner          `json:"owner"`
	Receiver            *ReceiverFlow         `json:"receiver,omitempty"`
	Redirect            *RedirectFlow         `json:"redirect,omitempty"`
	StatementDescriptor string                `json:"statement_descriptor"`
	SourceOrder         *SourceSourceOrder    `json:"source_order"`
	Status              SourceStatus          `json:"status"`
	Type                string                `json:"type"`
	TypeData            map[string]interface{}
	Usage               SourceUsage `json:"usage"`
}

// AppendTo implements custom encoding logic for SourceObjectParams so that the special
// "TypeData" value for is sent as the correct parameter based on the Source type
func (p *SourceObjectParams) AppendTo(body *form.Values, keyParts []string) {
	if len(p.TypeData) > 0 && p.Type == nil {
		panic("You can not fill TypeData if you don't explicitly set Type")
	}

	for k, vs := range p.TypeData {
		body.Add(form.FormatKey(append(keyParts, StringValue(p.Type), k)), vs)
	}
}

// UnmarshalJSON handles deserialization of an Source. This custom unmarshaling
// is needed to extract the type specific data (accessible under `TypeData`)
// but stored in JSON under a hash named after the `type` of the source.
func (s *Source) UnmarshalJSON(data []byte) error {
	type source Source
	var v source
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*s = Source(v)

	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if d, ok := raw[s.Type]; ok {
		if m, ok := d.(map[string]interface{}); ok {
			s.TypeData = m
		}
	}

	return nil
}
