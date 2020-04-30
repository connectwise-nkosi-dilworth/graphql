package pocfex63

import (
	"context"
	"fmt"
)

type HideFieldMap struct {
	IDField         string // partnerID-clientID
	PartnerID       string
	ClientID        string
	HideTheseFields *[]string
	// AllowedDirectives *[]Directive
}

var hideFieldsMap = []*HideFieldMap{
	{
		IDField:   "50016865-50121580",
		PartnerID: "50016865",
		ClientID:  "50121580",
		HideTheseFields: &[]string{
			"Sites",
		},
	},
	{
		IDField:   "50016865-6500023",
		PartnerID: "50016865",
		ClientID:  "6500023",
		HideTheseFields: &[]string{
			"Endpoints",
			"sites",
			"siteId",
			"endpointId",
			"Site",
			"Endpoint",
		},
	},
}

var hiddenFieldsMap = make(map[string]*HideFieldMap)

func init() {
	fmt.Println("HiddenFields init() function")
	for _, hFld := range hideFieldsMap {
		hiddenFieldsMap[hFld.IDField] = hFld
	}
	fmt.Printf("INIT -done- Print hiddenFieldsMap: %+v\n\n", hiddenFieldsMap)
}

// PocInputParams : data into func
type PocInputParams struct {
	// Schema         Schema
	Root           interface{}
	VariableValues map[string]interface{}
	Context        context.Context
}

// IsHiddenPartnerClientField : Hide fields from schema output by PartnerID/ClientID
func IsHiddenPartnerClientField(p PocInputParams, fieldName string) bool {
	if len(fieldName) == 0 {
		return false
	}
	ctx := p.Context

	//
	partnerID := ctx.Value("PartnerID").(string)
	clientID := ctx.Value("ClientID").(string)
	IDField := fmt.Sprintf("%v-%v", partnerID, clientID)
	fmt.Printf("Print IDField (partnerID-clientID): %v\n\n", IDField)

	partnerClientHideFieldMap := hiddenFieldsMap[IDField]
	fmt.Printf("Print HideFieldMap: %+v\n\n", partnerClientHideFieldMap)

	fmt.Printf("Print Search Fieldname: %v\n", fieldName)
	if partnerClientHideFieldMap != nil {
		for i := 0; i < len(*partnerClientHideFieldMap.HideTheseFields); i++ {
			hideField := (*partnerClientHideFieldMap.HideTheseFields)[i]
			// Found hidden field
			if hideField == fieldName {
				fmt.Printf("Print FOUND HIDDEN Fieldname: %v = %v\n", hideField, fieldName)
				return true
			}
		}
	}

	return false
}

// RedoResultMap : take out hidden fields
func RedoResultMap(p PocInputParams, inResult []interface{}) interface{} {
	newResult := make([]interface{}, 1 /*len(*inResult)*/)

	return newResult
}
