// Copyright 2019 A. Wolcott. All rights reserved.
//
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package daddy

// Consent represents the agreement to terms for a service
type Consent struct {
	AgreedAt      string
	AgreedBy      string
	AgreementKeys []string
}

// DNSRecord represents an individual DNS record
type DNSRecord struct {
	Data     string `json:"data"`
	Name     string `json:"name"`
	TTL      int    `json:"ttl"`
	Type     string `json:"type"`
}

// DomainAvailableResponse contains the current availability of a domain, as
// well as data related to purchasing the domain
type DomainAvailableResponse struct {
	Available  bool
	Currency   string
	Definitive bool
	Domain     string
	Period     int
	Price      int
}

// DomainAvailableError contains error information if a domain availability check fails
type DomainAvailableError struct {
	Code    string
	Domain  string
	Message string
	Path    string
	Status  int
}

// DomainAvailableBulkMixed is a series of arries for domain availability
// responses and errors when doing a bulk request
type DomainAvailableBulkMixed struct {
	Domains []DomainAvailableResponse
	Errors  []DomainAvailableError
}

// DomainContacts contains the four types of WHOIS contacts for a domain
type DomainContacts struct {
	Admin      Contact
	Billing    Contact
	Registrant Contact
	Tech       Contact
}

type DomainDetail struct {
	DomainSummary
	SubaccountID  string
	Verifications VerificationsDomain
}

type DomainPurchase struct {
	Consent Consent
	DomainContacts
	Domain      string
	NameServers []string
	Period      int
	Privacy     bool
	RenewAuto   bool
}

type DomainPurchaseResponse struct {
	Currency  string
	ItemCount int
	OrderID   int
	Total     int
}

type DomainRenew struct {
	Period int
}

type DomainSuggestion struct {
	Domain string
}

type DomainSummary struct {
	AuthCode string
	DomainContacts
	CreatedAt               string
	DeletedAt               string
	TransferAwayEligibileAt string
	Domain                  string
	DomainID                int
	ExpirationProtected     bool
	Expires                 string
	HoldRegistrar           bool
	Locked                  bool
	NameServers             []string
	Privacy                 bool
	RenewAuto               bool
	RenewDeadline           string
	Renewable               bool
	Status                  string
	TransferProtected       bool
}

type DomainForwardingMask struct {
	Title       string
	Description string
	Keywords    string
}

type DomainForwardingCreate struct {
	Type string
	URL  string
	Mask DomainForwardingMask
}

type DomainForwarding struct {
	FQDN string
	DomainForwardingCreate
}

type DomainTransferIn struct {
	AuthCode  string
	Consent   Consent
	Period    int
	Privacy   bool
	RenewAuto bool
	DomainContacts
}

type DomainUpdate struct {
	Locked       bool
	NameServers  []string
	RenewAuto    bool
	SubaccountID string
}

type DomainContactsBulk struct {
	DomainContacts
	ContactPresence Contact
	Domains         []string
	EntityType      string
}

type ErrorDomainContactsValidate struct {
	Code    string
	Fields  []ErrorFieldDomainContactsValidate
	Message string
	Stack   []string
}

type ErrorFieldDomainContactsValidate struct {
	ErrorField
	Domains []string
}

type PrivacyPurchase struct {
	Consent Consent
}

type RealNameValidation struct {
	Status string
}

type TldSummary struct {
	Name string
	Type string
}

type VerificationDomainName struct {
	Status string
}

type VerificationRealName struct {
	Status string
}

type VerificationsDomain struct {
	DomainName VerificationDomainName
	RealName   VerificationRealName
}
