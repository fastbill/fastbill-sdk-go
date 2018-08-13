package customer

import (
	"time"
)

// type Customers struct {
// 	XMLName  xml.Name   `xml:"CUSTOMERS"`
// 	Customer []Customer `xml:"CUSTOMER"`
// }

// Customer represents one customer
type Customer struct {
	CustomerID                  int        `xml:"CUSTOMER_ID"`
	CustomerNumber              int        `xml:"CUSTOMER_NUMBER"`
	DaysForPayment              int        `xml:"DAYS_FOR_PAYMENT"`
	Created                     *time.Time `xml:"CREATED"`
	PaymentType                 int        `xml:"PAYMENT_TYPE"`
	BankName                    string     `xml:"BANK_NAME"`
	BankAccountNumber           string     `xml:"BANK_ACCOUNT_NUMBER"` // maybe int?
	BankCode                    string     `xml:"BANK_CODE"`           // maybe int?
	BankAccountOwner            string     `xml:"BANK_ACCOUNT_OWNER"`
	BankIBAN                    string     `xml:"BANK_IBAN"`
	BankBIC                     string     `xml:"BANK_BIC"`
	BankAccountMandateReference string     `xml:"BANK_ACCOUNT_MANDATE_REFERENCE"` // maybe int?
	ShowPaymentNotice           bool       `xml:"SHOW_PAYMENT_NOTICE"`
	AccountReceivable           string     `xml:"ACCOUNT_RECEIVABLE"` // maybe int?
	CustomerType                string     `xml:"CUSTOMER_TYPE"`      // maybe int or bool?
	Top                         int        `xml:"TOP"`                // no explanation, seems to be always 0
	Organization                string     `xml:"ORGANIZATION"`       // required if type is business
	Position                    string     `xml:"POSITION"`
	AcademicDegree              string     `xml:"ACADEMIC_DEGREE"`
	Salutation                  string     `xml:"SALUTATION"`
	FirstName                   string     `xml:"FIRST_TIME"` // required if type is consumer
	LastName                    string     `xml:"LAST_NAME"`  // required if type is consumer
	Address                     string     `xml:"ADDRESS"`
	Address2                    string     `xml:"ADDRESS2"`
	Zipcode                     string     `xml:"ZIPCODE"`
	City                        string     `xml:"CITY"`
	CountryCode                 string     `xml:"COUNTRY_CODE"`
	SecondaryAddress            string     `xml:"SECONDARY_ADDRESS"`
	Phone                       string     `xml:"PHONE"`
	Phone2                      string     `xml:"PHONE2"`
	Fax                         string     `xml:"FAX"`
	Mobile                      string     `xml:"MOBILE"`
	Email                       string     `xml:"EMAIL"`
	Website                     string     `xml:"WEBSITE"`
	VATID                       string     `xml:"VAT_ID"`
	CurrencyCode                string     `xml:"CURRENCY_CODE"`
	LastUpdate                  *time.Time `xml:"LASTUPDATE"`
	Tags                        string     `xml:"TAGS"`
	// NewsletterOptIn bool
} // TODO which ones are optional?
