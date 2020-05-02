// Code generated by go-swagger; DO NOT EDIT.

package specialties

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSearchSpecialtyParams creates a new SearchSpecialtyParams object
// with the default values initialized.
func NewSearchSpecialtyParams() SearchSpecialtyParams {

	var (
		// initialize parameters with default values

		categoryDefault    = string("")
		subCategoryDefault = string("")
	)

	return SearchSpecialtyParams{
		Category: &categoryDefault,

		SubCategory: &subCategoryDefault,
	}
}

// SearchSpecialtyParams contains all the bound params for the search specialty operation
// typically these are obtained from a http.Request
//
// swagger:parameters searchSpecialty
type SearchSpecialtyParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*main category
	  In: query
	  Default: ""
	*/
	Category *string
	/*secundary category
	  In: query
	  Default: ""
	*/
	SubCategory *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSearchSpecialtyParams() beforehand.
func (o *SearchSpecialtyParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCategory, qhkCategory, _ := qs.GetOK("category")
	if err := o.bindCategory(qCategory, qhkCategory, route.Formats); err != nil {
		res = append(res, err)
	}

	qSubCategory, qhkSubCategory, _ := qs.GetOK("subCategory")
	if err := o.bindSubCategory(qSubCategory, qhkSubCategory, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCategory binds and validates parameter Category from query.
func (o *SearchSpecialtyParams) bindCategory(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewSearchSpecialtyParams()
		return nil
	}

	o.Category = &raw

	return nil
}

// bindSubCategory binds and validates parameter SubCategory from query.
func (o *SearchSpecialtyParams) bindSubCategory(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewSearchSpecialtyParams()
		return nil
	}

	o.SubCategory = &raw

	return nil
}