package utils

import (
	"sort"

	"github.com/samber/lo"
)

func CustomerPropertyMap() map[string]string {
	propertyMap := make(map[string]string, 0)

	// propertyMap["Title"] = "Title"
	// propertyMap["First name"] = "FirstName"
	// propertyMap["Last name"] = "LastName"
	// propertyMap["Middle name"] = "MiddleName"
	// propertyMap["Suffix"] = "Suffix"
	// propertyMap["Display name"] = "DisplayName"
	// propertyMap["Company name"] = "CompanyName"
	// propertyMap["Business ID"] = "BusinessID"
	// propertyMap["Phone number"] = "Phone"
	// propertyMap["Country code of phone number"] = "PhoneCountry"
	// propertyMap["Mobile number"] = "Mobile"
	// propertyMap["Country code of mobile number"] = "MobileCountry"
	// propertyMap["Email"] = "Email"
	// propertyMap["Other contact"] = "OtherContact"
	// propertyMap["Website address"] = "WebsiteAddress"
	// propertyMap["Customer address"] = "CustomerAddress"
	// propertyMap["Note"] = "Note"
	// propertyMap["Tax information"] = "TaxInformation"
	// propertyMap["Street(Billing address)"] = "BillingAddressStreet"
	// propertyMap["City(Billing address)"] = "BillingAddressCity"
	// propertyMap["State(Billing address)"] = "BillingAddressState"
	// propertyMap["Postal code(Billing address)"] = "BillingAddressPostalCode"
	// propertyMap["Postal code(Billing address)"] = "BillingAddressCountry"
	// propertyMap["Street(Shipping address)"] = "ShippingAddressStreet"
	// propertyMap["City(Shipping address)"] = "ShippingAddressCity"
	// propertyMap["State(Shipping address)"] = "ShippingAddressState"
	// propertyMap["Postal code(Shipping address)"] = "ShippingAddressPostalCode"
	// propertyMap["Country(Shipping address)"] = "ShippingAddressCountry"
	propertyMap["Title"] = "Title"
	propertyMap["FirstName"] = "First Name"
	propertyMap["LastName"] = "Last Name"
	propertyMap["MiddleName"] = "Middle Name"
	propertyMap["Suffix"] = "Suffix"
	propertyMap["DisplayName"] = "Display Name"
	propertyMap["CompanyName"] = "Company Name"
	propertyMap["BusinessID"] = "Business ID"
	propertyMap["Phone"] = "Phone Number"
	propertyMap["PhoneCountry"] = "Country code of phone number"
	propertyMap["Mobile"] = "Mobile"
	propertyMap["MobileCountry"] = "Country code of mobile number"
	propertyMap["Email"] = "Email"
	propertyMap["OtherContact"] = "Other Contact"
	propertyMap["WebsiteAddress"] = "Website Address"
	propertyMap["CustomerAddress"] = "Customer Address"
	propertyMap["Note"] = "Note"
	propertyMap["TaxInformation"] = "Tax Information"
	propertyMap["BillingAddressStreet"] = "Street(Billing address)"
	propertyMap["BillingAddressCity"] = "City(Billing address))"
	propertyMap["BillingAddressState"] = "State(Billing address)"
	propertyMap["BillingAddressPostalCode"] = "Postal code(Billing address)"
	propertyMap["BillingAddressCountry"] = "Country(Billing address)"
	propertyMap["ShippingAddressStreet"] = "Street(Shipping address)"
	propertyMap["ShippingAddressCity"] = "City(Shipping address)"
	propertyMap["ShippingAddressState"] = "State(Shipping address)"
	propertyMap["ShippingAddressPostalCode"] = "Postal code(Shipping address)"
	propertyMap["ShippingAddressCountry"] = "Country(Shipping address)"
	propertyMap["Language"] = "Language"
	propertyMap["Status"] = "Status"

	return propertyMap
}

func RequiredField() map[string]bool {
	requiredFields := make(map[string]bool, 0)
	requiredFields["Display Name"] = true
	return requiredFields
}

func Demo() map[string]string {
	propertyMap := make(map[string]string, 0)
	propertyMap1 := make(map[string]string, 0)
	propertyMap1["Name"] = "Name4"
	propertyMap1["Age"] = "Age"
	propertyMap1["Institution Name"] = "Institution.Name1"
	propertyMap1["Institution Address"] = "Institution.Address"
	propertyMap1["Dept"] = "Institution.Faculty.Dept"
	propertyMap1["ID"] = "Institution.Faculty.ID"
	propertyMap1["CGPA"] = "Institution.Faculty.CGPA"
	propertyMap1["Height"] = "Height"
	propertyMap1["Company Name"] = "Company.Name3"
	propertyMap1["Salary"] = "Company.Salary"
	propertyMap1["Currency"] = "Company.Currency"
	propertyMap1["Designation Name"] = "Company.Designation.Name2"
	propertyMap1["Dev"] = "Company.Designation.Dev"

	propertyMap["Name4"] = "Name"
	propertyMap["Age"] = "Age"
	propertyMap["Institution.Name1"] = "Institution Name"
	propertyMap["Institution.Address"] = "Institution Address"
	propertyMap["Institution.Faculty.Dept"] = "Dept"
	propertyMap["Institution.Faculty.ID"] = "ID"
	propertyMap["Institution.Faculty.CGPA"] = "CGPA"
	propertyMap["Height"] = "Height"
	propertyMap["Company.Name3"] = "Company Name"
	propertyMap["Company.Salary"] = "Salary"
	propertyMap["Company.Currency"] = "Currency"
	propertyMap["Company.Designation.Name2"] = "Designation Name"
	propertyMap["Company.Designation.Dev"] = "Dev"
	propertyMap["Company.Designation.Test"] = "Designation test"
	propertyMap["Company.Designation.RequireField"] = "Require Field"
	return propertyMap
}

func FilterIDs(oldIDs []int, requestIDs []int) ([]int, []int, []int) {
	sort.Ints(oldIDs)
	sort.Ints(requestIDs)
	oldIDs = lo.Uniq(oldIDs)
	requestIDs = lo.Uniq(requestIDs)

	visited := make(map[int]bool, 0)
	for _, id := range oldIDs {
		visited[id] = true
	}

	newIDs := make([]int, 0)
	updateIDs := make([]int, 0)
	for _, id := range requestIDs {
		if !visited[id] {
			newIDs = append(newIDs, id)
		} else {
			visited[id] = false
			updateIDs = append(updateIDs, id)
		}
	}

	removeIDs := make([]int, 0)
	for _, id := range oldIDs {
		if visited[id] {
			removeIDs = append(removeIDs, id)
		}
	}

	return newIDs, removeIDs, updateIDs
}
