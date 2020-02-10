package stripe

import "encoding/json"

// VerificationDocumentDetailsCode is a machine-readable code specifying the verification state of
// a document associated with a person.
type VerificationDocumentDetailsCode string

// List of values that IdentityVerificationDetailsCode can take.
const (
	VerificationDocumentDetailsCodeDocumentCorrupt               VerificationDocumentDetailsCode = "document_corrupt"
	VerificationDocumentDetailsCodeDocumentFailedCopy            VerificationDocumentDetailsCode = "document_failed_copy"
	VerificationDocumentDetailsCodeDocumentFailedGreyscale       VerificationDocumentDetailsCode = "document_failed_greyscale"
	VerificationDocumentDetailsCodeDocumentFailedOther           VerificationDocumentDetailsCode = "document_failed_other"
	VerificationDocumentDetailsCodeDocumentFailedTestMode        VerificationDocumentDetailsCode = "document_failed_test_mode"
	VerificationDocumentDetailsCodeDocumentFraudulent            VerificationDocumentDetailsCode = "document_fraudulent"
	VerificationDocumentDetailsCodeDocumentIDTypeNotSupported    VerificationDocumentDetailsCode = "document_id_type_not_supported"
	VerificationDocumentDetailsCodeDocumentIDCountryNotSupported VerificationDocumentDetailsCode = "document_id_country_not_supported"
	VerificationDocumentDetailsCodeDocumentManipulated           VerificationDocumentDetailsCode = "document_manipulated"
	VerificationDocumentDetailsCodeDocumentMissingBack           VerificationDocumentDetailsCode = "document_missing_back"
	VerificationDocumentDetailsCodeDocumentMissingFront          VerificationDocumentDetailsCode = "document_missing_front"
	VerificationDocumentDetailsCodeDocumentNotReadable           VerificationDocumentDetailsCode = "document_not_readable"
	VerificationDocumentDetailsCodeDocumentNotUploaded           VerificationDocumentDetailsCode = "document_not_uploaded"
	VerificationDocumentDetailsCodeDocumentTooLarge              VerificationDocumentDetailsCode = "document_too_large"
)

// PersonVerificationDetailsCode is a machine-readable code specifying the verification state of a
// person.
type PersonVerificationDetailsCode string

// List of values that IdentityVerificationDetailsCode can take.
const (
	PersonVerificationDetailsCodeFailedKeyedIdentity PersonVerificationDetailsCode = "failed_keyed_identity"
	PersonVerificationDetailsCodeFailedOther         PersonVerificationDetailsCode = "failed_other"
	PersonVerificationDetailsCodeScanNameMismatch    PersonVerificationDetailsCode = "scan_name_mismatch"
)

// IdentityVerificationStatus describes the different statuses for identity verification.
type IdentityVerificationStatus string

// List of values that IdentityVerificationStatus can take.
const (
	IdentityVerificationStatusPending    IdentityVerificationStatus = "pending"
	IdentityVerificationStatusUnverified IdentityVerificationStatus = "unverified"
	IdentityVerificationStatusVerified   IdentityVerificationStatus = "verified"
)

// DOBParams represents a DOB during account creation/updates.
type DOBParams struct {
	Day   *int64 `form:"day" json:"day"`
	Month *int64 `form:"month" json:"month"`
	Year  *int64 `form:"year" json:"year"`
}

// RelationshipParams is used to set the relationship between an account and a person.
type RelationshipParams struct {
	Director         *bool    `form:"director" json:"director"`
	Executive        *bool    `form:"executive" json:"executive"`
	Owner            *bool    `form:"owner" json:"owner"`
	PercentOwnership *float64 `form:"percent_ownership" json:"percent_ownership"`
	Representative   *bool    `form:"representative" json:"representative"`
	Title            *string  `form:"title" json:"title"`
}

// PersonVerificationDocumentParams represents the parameters available for the document verifying
// a person's identity.
type PersonVerificationDocumentParams struct {
	Back  *string `form:"back" json:"back"`
	Front *string `form:"front" json:"front"`
}

// PersonVerificationParams is used to represent parameters associated with a person's verification
// details.
type PersonVerificationParams struct {
	AdditionalDocument *PersonVerificationDocumentParams `form:"additional_document" json:"additional_document"`
	Document           *PersonVerificationDocumentParams `form:"document" json:"document"`
}

// PersonParams is the set of parameters that can be used when creating or updating a person.
// For more details see https://stripe.com/docs/api#create_person.
type PersonParams struct {
	Params         `form:"*" json:"*"`
	Account        *string                   `form:"-" json:"-"` // Included in URL
	Address        *AccountAddressParams     `form:"address" json:"address"`
	AddressKana    *AccountAddressParams     `form:"address_kana" json:"address_kana"`
	AddressKanji   *AccountAddressParams     `form:"address_kanji" json:"address_kanji"`
	DOB            *DOBParams                `form:"dob" json:"dob"`
	Email          *string                   `form:"email" json:"email"`
	FirstName      *string                   `form:"first_name" json:"first_name"`
	FirstNameKana  *string                   `form:"first_name_kana" json:"first_name_kana"`
	FirstNameKanji *string                   `form:"first_name_kanji" json:"first_name_kanji"`
	Gender         *string                   `form:"gender" json:"gender"`
	IDNumber       *string                   `form:"id_number" json:"id_number"`
	LastName       *string                   `form:"last_name" json:"last_name"`
	LastNameKana   *string                   `form:"last_name_kana" json:"last_name_kana"`
	LastNameKanji  *string                   `form:"last_name_kanji" json:"last_name_kanji"`
	MaidenName     *string                   `form:"maiden_name" json:"maiden_name"`
	PersonToken    *string                   `form:"person_token" json:"person_token"`
	Phone          *string                   `form:"phone" json:"phone"`
	Relationship   *RelationshipParams       `form:"relationship" json:"relationship"`
	SSNLast4       *string                   `form:"ssn_last_4" json:"ssn_last_4"`
	Verification   *PersonVerificationParams `form:"verification" json:"verification"`
}

// RelationshipListParams is used to filter persons by the relationship
type RelationshipListParams struct {
	Director       *bool `form:"director" json:"director"`
	Executive      *bool `form:"executive" json:"executive"`
	Owner          *bool `form:"owner" json:"owner"`
	Representative *bool `form:"representative" json:"representative"`
}

// PersonListParams is the set of parameters that can be used when listing persons.
// For more detail see https://stripe.com/docs/api#list_persons.
type PersonListParams struct {
	ListParams   `form:"*" json:"*"`
	Account      *string                 `form:"-" json:"-"` // Included in URL
	Relationship *RelationshipListParams `form:"relationship" json:"relationship"`
}

// DOB represents a Person's date of birth.
type DOB struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

// Relationship represents how the Person relates to the business.
type Relationship struct {
	Director         bool    `json:"director"`
	Executive        bool    `json:"executive"`
	Owner            bool    `json:"owner"`
	PercentOwnership float64 `json:"percent_ownership"`
	Representative   bool    `json:"representative"`
	Title            string  `json:"title"`
}

// Requirements represents what's missing to verify a Person.
type Requirements struct {
	CurrentlyDue        []string `json:"currently_due"`
	EventuallyDue       []string `json:"eventually_due"`
	PastDue             []string `json:"past_due"`
	PendingVerification []string `json:"pending_verification"`
}

// PersonVerificationDocument represents the documents associated with a Person.
type PersonVerificationDocument struct {
	Back        *File                           `json:"back"`
	Details     string                          `json:"details"`
	DetailsCode VerificationDocumentDetailsCode `json:"details_code"`
	Front       *File                           `json:"front"`
}

// PersonVerification is the structure for a person's verification details.
type PersonVerification struct {
	AdditionalDocument *PersonVerificationDocument   `json:"additional_document"`
	Details            string                        `json:"details"`
	DetailsCode        PersonVerificationDetailsCode `json:"details_code"`
	Document           *PersonVerificationDocument   `json:"document"`
	Status             IdentityVerificationStatus    `json:"status"`
}

// Person is the resource representing a Stripe person.
// For more details see https://stripe.com/docs/api#persons.
type Person struct {
	Account          string              `json:"account"`
	Address          *AccountAddress     `json:"address"`
	AddressKana      *AccountAddress     `json:"address_kana"`
	AddressKanji     *AccountAddress     `json:"address_kanji"`
	Deleted          bool                `json:"deleted"`
	DOB              *DOB                `json:"dob"`
	Email            string              `json:"email"`
	FirstName        string              `json:"first_name"`
	FirstNameKana    string              `json:"first_name_kana"`
	FirstNameKanji   string              `json:"first_name_kanji"`
	Gender           string              `json:"gender"`
	ID               string              `json:"id"`
	IDNumberProvided bool                `json:"id_number_provided"`
	LastName         string              `json:"last_name"`
	LastNameKana     string              `json:"last_name_kana"`
	LastNameKanji    string              `json:"last_name_kanji"`
	MaidenName       string              `json:"maiden_name"`
	Metadata         map[string]string   `json:"metadata"`
	Object           string              `json:"object"`
	Phone            string              `json:"phone"`
	Relationship     *Relationship       `json:"relationship"`
	Requirements     *Requirements       `json:"requirements"`
	SSNLast4Provided bool                `json:"ssn_last_4_provided"`
	Verification     *PersonVerification `json:"verification"`
}

// PersonList is a list of persons as retrieved from a list endpoint.
type PersonList struct {
	ListMeta
	Data []*Person `json:"data"`
}

// UnmarshalJSON handles deserialization of a Person.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Person) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type person Person
	var v person
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Person(v)
	return nil
}
