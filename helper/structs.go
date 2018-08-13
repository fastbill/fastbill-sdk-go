package helper

import (
	"encoding/xml"

	"github.com/fastbill/sdk-go/customer"
	"github.com/fastbill/sdk-go/filters"
)

type FBAPI struct {
	XMLName xml.Name `xml:"FBAPI"`
}

type FBAPIWithResponse struct {
	FBAPI
	Request  Request  `xml:"REQUEST"`
	Response Response `xml:"RESPONSE"`
}

type FBAPIWithFilter struct {
	FBAPI
	Filter filters.Filter `xml:"FILTER"`
}

type Request struct {
	XMLName xml.Name `xml:"REQUEST"`
	Service Service  `xml:"SERVICE"`
}

type Response struct {
	XMLName  xml.Name            `xml:"RESPONSE"`
	Customer []customer.Customer `xml:"CUSTOMERS"`
}

type Service struct {
	Service string `xml:"SERVICE"`
}
