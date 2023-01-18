package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tidwall/gjson"
)

type Details []struct {
	EmailID    string `json:"EmailID"`
	LastName   string `json:"LastName"`
	EmployeeID string `json:"EmployeeID"`
	ZohoID     int64  `json:"Zoho_ID"`
	FirstName  string `json:"FirstName"`
}

func main() {
	jsonRaw1 := `{
		"response": {
			"result": [
				{
					"496960000005047035": [
						{
							"EmailID": "aditya@coffeebeans.io",
							"CreatedTime": "1638938845189",
							"Employee_type.id": "496960000000039959",
							"Address": "Vasant Nagar Lane no. 02, Near Mahadev Temple, Pusad, Maharashtra. Pin no. : 445204",
							"Date_of_birth": "04-Dec-1997",
							"Photo": "https://contacts.zoho.com/file?ID=766621431&fs=thumb",
							"AddedTime": "08-Dec-2021 10:17:25",
							"Marital_status": "Single",
							"ModifiedBy": "159 - Diya - Gandhi",
							"ApprovalStatus": "Approval Not Enabled",
							"Department": "DSD",
							"LocationName.ID": "496960000007578003",
							"tabularSections": {
								"Education": [
									{}
								],
								"Work experience": [
									{}
								],
								"Dependent": [
									{}
								]
							},
							"Mobile.country_code": "in",
							"AddedBy": "0 - Finance - NA",
							"Tags": "",
							"Reporting_To": "Ashwin Anupam Dalela 98",
							"Photo_downloadUrl": "https://contacts.zoho.com/file?ID=766621431&fs=thumb",
							"Source_of_hire.id": "496960000000039975",
							"Employeestatus": "Active",
							"Role": "Team member",
							"Employee_type": "Permanent",
							"AddedBy.ID": "496960000000136005",
							"Role.ID": "496960000000035635",
							"LastName": "Sahatonde",
							"EmployeeID": "136",
							"ZUID": "766621431",
							"Current_Job_Description": "",
							"Other_Email": "adityasahatonde555@gmail.com",
							"Work_location": "Pune",
							"LocationName": "Pune",
							"Nick_Name": "",
							"ModifiedTime": "1668408829965",
							"Reporting_To.MailID": "ashwin@coffeebeans.io",
							"Zoho_ID": 496960000005047035,
							"Designation.ID": "496960000006167003",
							"Source_of_hire": "Direct",
							"Designation": "Software Engineer",
							"Marital_status.id": "496960000000039985",
							"FirstName": "Aditya Sanjay",
							"Dateofjoining": "29-Nov-2021",
							"AboutMe": "",
							"Mobile": "91-9284553224",
							"Extension": "",
							"ModifiedBy.ID": "496960000005683043",
							"Reporting_To.ID": "496960000003800137",
							"Work_phone": "",
							"Department.ID": "496960000006118817",
							"Expertise": ""
						}
					]
				}
			],
			"message": "Data fetched successfully",
			"uri": "/api/forms/employee/getRecords",
			"status": 0
		}
	}`

	zohoResponse, err := processZohoResponse(jsonRaw1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(zohoResponse[0].ZohoID)
}

func processZohoResponse(rawJsonString string) (Details, error) {
	var final Details
	jsonResponseMessage := gjson.Get(rawJsonString, "response.message")
	fmt.Println(jsonResponseMessage)
	if jsonResponseMessage.String() == "Data fetched successfully" {
		jsonT := gjson.Get(rawJsonString, "response.result.0")
		if len(jsonT.String()) == 0 {
			err := errors.New("Parsing Zoho Response Issue")
			fmt.Println("Error :", err)
			return final, err
		}
		jsonT.ForEach(func(key, value gjson.Result) bool {
			_ = json.Unmarshal([]byte(value.String()), &final)
			return false // keep iterating
		})
		return final, nil
	} else {
		return final, errors.New("Zoho API Hit Unsuccessful")
	}

}
