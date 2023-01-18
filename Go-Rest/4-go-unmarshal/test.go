package main

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
)

type Details []struct {
	EmailID         string `json:"EmailID"`
	CreatedTime     string `json:"CreatedTime"`
	EmployeeTypeID  string `json:"Employee_type.id"`
	Address         string `json:"Address"`
	DateOfBirth     string `json:"Date_of_birth"`
	Photo           string `json:"Photo"`
	AddedTime       string `json:"AddedTime"`
	MaritalStatus   string `json:"Marital_status"`
	ModifiedBy      string `json:"ModifiedBy"`
	ApprovalStatus  string `json:"ApprovalStatus"`
	Department      string `json:"Department"`
	LocationNameID  string `json:"LocationName.ID"`
	TabularSections struct {
		Education []struct {
		} `json:"Education"`
		WorkExperience []struct {
		} `json:"Work experience"`
		Dependent []struct {
		} `json:"Dependent"`
	} `json:"tabularSections"`
	MobileCountryCode     string `json:"Mobile.country_code"`
	AddedBy               string `json:"AddedBy"`
	Tags                  string `json:"Tags"`
	ReportingTo           string `json:"Reporting_To"`
	PhotoDownloadURL      string `json:"Photo_downloadUrl"`
	SourceOfHireID        string `json:"Source_of_hire.id"`
	Employeestatus        string `json:"Employeestatus"`
	Role                  string `json:"Role"`
	EmployeeType          string `json:"Employee_type"`
	AddedByID             string `json:"AddedBy.ID"`
	RoleID                string `json:"Role.ID"`
	LastName              string `json:"LastName"`
	EmployeeID            string `json:"EmployeeID"`
	Zuid                  string `json:"ZUID"`
	CurrentJobDescription string `json:"Current_Job_Description"`
	OtherEmail            string `json:"Other_Email"`
	WorkLocation          string `json:"Work_location"`
	LocationName          string `json:"LocationName"`
	NickName              string `json:"Nick_Name"`
	ModifiedTime          string `json:"ModifiedTime"`
	ReportingToMailID     string `json:"Reporting_To.MailID"`
	ZohoID                int64  `json:"Zoho_ID"`
	DesignationID         string `json:"Designation.ID"`
	SourceOfHire          string `json:"Source_of_hire"`
	Designation           string `json:"Designation"`
	MaritalStatusID       string `json:"Marital_status.id"`
	FirstName             string `json:"FirstName"`
	Dateofjoining         string `json:"Dateofjoining"`
	AboutMe               string `json:"AboutMe"`
	Mobile                string `json:"Mobile"`
	Extension             string `json:"Extension"`
	ModifiedByID          string `json:"ModifiedBy.ID"`
	ReportingToID         string `json:"Reporting_To.ID"`
	WorkPhone             string `json:"Work_phone"`
	DepartmentID          string `json:"Department.ID"`
	Expertise             string `json:"Expertise"`
}
type Employee struct {
	Response struct {
		Result []struct {
			EmpDetails map[string]Details //`json:"496960000005885093"`
		} `json:"result"`
		Message string `json:"message"`
		URI     string `json:"uri"`
		Status  int    `json:"status"`
	} `json:"response"`
}

type Test struct {
	Employee map[string]Details `json:"496960000005885093"`
}

func main() {
	fmt.Println("Test Files...")
	jsonRaw1 := `{"response":{"result":[{"496960000005885093":[{"EmailID":"faiza@coffeebeans.io","CreatedTime":"1655285850134","Employee_type.id":"496960000000039959","Address":"Raheja Vistas Phase - 3,NIBM Post office road, Mohammed Wadi, Pune 411048","Date_of_birth":"24-May-1997","Photo":"https://contacts.zoho.com/file?ID=782130219&fs=thumb","AddedTime":"15-Jun-2022 15:07:30","Marital_status":"Married","ModifiedBy":"109 - Aditi - Bhate","ApprovalStatus":"Approval Not Enabled","Department":"Recruitment","LocationName.ID":"496960000007578003","tabularSections":{"Education":[{}],"Work experience":[{}],"Dependent":[{}]},"Mobile.country_code":"in","AddedBy":"159 - Diya - Gandhi","Tags":"","Reporting_To":"Rajesh B 135","Photo_downloadUrl":"https://contacts.zoho.com/file?ID=782130219&fs=thumb","Source_of_hire.id":"496960000000039977","Employeestatus":"Active","Role":"Team member","Employee_type":"Permanent","AddedBy.ID":"496960000005683043","Role.ID":"496960000000035635","LastName":"A H","EmployeeID":"204","ZUID":"782130219","Current_Job_Description":"","Other_Email":"faiza.akthar5@gmail.com","Work_location":"Pune","LocationName":"Pune","Nick_Name":"","ModifiedTime":"1668664335035","Reporting_To.MailID":"rajeshb@coffeebeans.io","Zoho_ID":496960000006413059,"Designation.ID":"496960000004365027","Source_of_hire":"Referral","Designation":"Recruiter","Marital_status.id":"496960000000039987","FirstName":"Faiza Banu","Dateofjoining":"13-Jun-2022","AboutMe":"","Mobile":"91-9578678677","Extension":"","ModifiedBy.ID":"496960000003491133","Reporting_To.ID":"496960000004936246","Work_phone":"9578678677","Department.ID":"496960000003698230","Expertise":""}]}],"message":"Data fetched successfully","uri":"/api/forms/employee/getRecords","status":0}}`
	jsonRaw := json.RawMessage(`[{
			"EmailID": "rahil@coffeebeans.io",
			"CreatedTime": "1649318693885",
			"Employee_type.id": "496960000000039959",
			"Address": "201, Vimal Apartments, Swedhganga Society, Canal Road, Warje, Pune\nDist. Pune, Maharashtra - 411058",
			"Date_of_birth": "11-Nov-1997",
			"Photo": "https://contacts.zoho.com/file?ID=775949004&fs=thumb",
			"AddedTime": "07-Apr-2022 13:34:53",
			"Marital_status": "Single",
			"ModifiedBy": "159 - Diya - Gandhi",
			"ApprovalStatus": "Approval Not Enabled",
			"Department": "Café",
			"LocationName.ID": "496960000007578003",
			"tabularSections": {
				"Education": [{}],
				"Work experience": [{}],
				"Dependent": [{}]
			},
			"Mobile.country_code": "",
			"AddedBy": "0 - Finance - NA",
			"Tags": "",
			"Reporting_To": "Sheetal Jain 95",
			"Photo_downloadUrl": "https://contacts.zoho.com/file?ID=775949004&fs=thumb",
			"Source_of_hire.id": "496960000000039979",
			"Employeestatus": "Active",
			"Role": "Team member",
			"Employee_type": "Permanent",
			"AddedBy.ID": "496960000000136005",
			"Role.ID": "496960000000035635",
			"LastName": "Jaiswal",
			"EmployeeID": "172",
			"ZUID": "775949004",
			"Current_Job_Description": "",
			"Other_Email": "rahilravijaiswal@gmail.com",
			"Work_location": "Pune",
			"LocationName": "Pune",
			"Nick_Name": "",
			"ModifiedTime": "1668492530492",
			"Reporting_To.MailID": "sheetal@coffeebeans.io",
			"Zoho_ID": 496960000005885093,
			"Designation.ID": "496960000006167003",
			"Source_of_hire": "Web",
			"Designation": "Software Engineer",
			"Marital_status.id": "496960000000039985",
			"FirstName": "Rahil",
			"Dateofjoining": "04-Apr-2022",
			"AboutMe": "",
			"Mobile": "9403382619",
			"Extension": "",
			"ModifiedBy.ID": "496960000005683043",
			"Reporting_To.ID": "496960000003622005",
			"Work_phone": "",
			"Department.ID": "496960000006122249",
			"Expertise": ""
		}]`)

	jsonRaw2 := json.RawMessage(`{
        "496960000005885093": [{
            "EmailID": "rahil@coffeebeans.io",
            "CreatedTime": "1649318693885",
            "Employee_type.id": "496960000000039959",
            "Address": "201, Vimal Apartments, Swedhganga Society, Canal Road, Warje, Pune\nDist. Pune, Maharashtra - 411058",
            "Date_of_birth": "11-Nov-1997",
            "Photo": "https://contacts.zoho.com/file?ID=775949004&fs=thumb",
            "AddedTime": "07-Apr-2022 13:34:53",
            "Marital_status": "Single",
            "ModifiedBy": "159 - Diya - Gandhi",
            "ApprovalStatus": "Approval Not Enabled",
            "Department": "Café",
            "LocationName.ID": "496960000007578003",
            "tabularSections": {
                "Education": [{}],
                "Work experience": [{}],
                "Dependent": [{}]
            },
            "Mobile.country_code": "",
            "AddedBy": "0 - Finance - NA",
            "Tags": "",
            "Reporting_To": "Sheetal Jain 95",
            "Photo_downloadUrl": "https://contacts.zoho.com/file?ID=775949004&fs=thumb",
            "Source_of_hire.id": "496960000000039979",
            "Employeestatus": "Active",
            "Role": "Team member",
            "Employee_type": "Permanent",
            "AddedBy.ID": "496960000000136005",
            "Role.ID": "496960000000035635",
            "LastName": "Jaiswal",
            "EmployeeID": "172",
            "ZUID": "775949004",
            "Current_Job_Description": "",
            "Other_Email": "rahilravijaiswal@gmail.com",
            "Work_location": "Pune",
            "LocationName": "Pune",
            "Nick_Name": "",
            "ModifiedTime": "1668492530492",
            "Reporting_To.MailID": "sheetal@coffeebeans.io",
            "Zoho_ID": 496960000005885093,
            "Designation.ID": "496960000006167003",
            "Source_of_hire": "Web",
            "Designation": "Software Engineer",
            "Marital_status.id": "496960000000039985",
            "FirstName": "Rahil",
            "Dateofjoining": "04-Apr-2022",
            "AboutMe": "",
            "Mobile": "9403382619",
            "Extension": "",
            "ModifiedBy.ID": "496960000005683043",
            "Reporting_To.ID": "496960000003622005",
            "Work_phone": "",
            "Department.ID": "496960000006122249",
            "Expertise": ""
        }]
    }`)

	var zohoResponse Details
	err := json.Unmarshal([]byte(jsonRaw), &zohoResponse)
	if err != nil {
		fmt.Println("Error Aaya :", err)
	}

	//fmt.Println(zohoResponse)

	testMap := make(map[string]Details)
	testMap["496960000005885093"] = zohoResponse

	// for k, v := range testMap {
	// 	fmt.Println(k, "  =>  ", v)
	// }

	zohoResponse2 := make(map[string]Details)
	//var zohoResponse2 Test
	err = json.Unmarshal([]byte(jsonRaw2), &zohoResponse2)
	if err != nil {
		fmt.Println("Error Aaya :", err)
	}
	//fmt.Println(zohoResponse2["496960000005885093"][0].EmailID)

	zohoResponse3 := make(map[string]interface{})
	err = json.Unmarshal([]byte(jsonRaw1), &zohoResponse3)
	if err != nil {
		fmt.Println("Error Aaya :", err)
	}
	fmt.Println(zohoResponse3["response"].(map[string]interface{})["result"])
	//zohoResponse4 := make(map[string]Details)
	zohoResponse4 := (zohoResponse3["response"].(map[string]interface{})["result"])
	fmt.Printf("%T %v", zohoResponse4, zohoResponse4)
	var final Details
	jsonT := gjson.Get(jsonRaw1, "response.result.0")
	fmt.Println("\n \n \n GJson Output", jsonT)
	jsonT.ForEach(func(key, value gjson.Result) bool {
		fmt.Printf("\n\n")
		fmt.Println(value.String())
		fmt.Printf("%T ", value.String())
		_ = json.Unmarshal([]byte(value.String()), &final)
		fmt.Println("\n\n ", final[0].ZohoID)
		return false // keep iterating
	})
	//fmt.Println("lenght of result : ", len(zohoResponse4[""]))
	// email := zohoResponse.Response.Result[0].Employee[0]["EmailID"]
	// firstName := zohoResponse.Response.Result[0].Employee[0]["FirstName"]
	// lastname := zohoResponse.Response.Result[0].Employee[0]["LastName"]
	// empId := zohoResponse.Response.Result[0].Employee[0]["EmployeeID"]

	//var zohoId int64
	//zohoId := fmt.Sprintf("%f", zohoResponse.Response.Result[0].Employee[0])
	//fmt.Printf(" Employeee Details %v %v %v %v %v", email, firstName, lastname, empId, zohoId)
}

// 496960000005885093
// 496960000005885093
