// recive/return data struct

package usecasemodel

import (
	"encoding/json"
	"encoding/xml"
)

type UserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

type UserRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

type DrmTokenResponse struct {
	DrmToken string `json:"drm_token"`
}

func UnmarshalLicenseequest(data []byte) (LicenseRequest, error) {
	var r LicenseRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *LicenseRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type LicenseRequest struct {
	UserID          string `json:"userid"`
	Password        string `json:"password,omitempty"`
	ContentID       string `json:"contentid"`
	PolicyID        string `json:"policyid"`
	CurrentTime     string `json:"currenttime"`
	UniqueMachineID string `json:"uniquemachineid"`
	IPAddress       string `json:"ipaddress"`
}

type AuthResult struct {
	XMLName     xml.Name        `xml:"auth_result"`
	Text        string          `xml:",chardata"`
	Status      AuthStatus      `xml:"status"`
	LicenseInfo AuthLicenseInfo `xml:"license_info"`
}

type AuthStatus struct {
	Text  string `xml:",chardata"`
	Error string `xml:"error"`
}

type AuthLicenseInfo struct {
	Text          string `xml:",chardata"`
	ConvertID     string `xml:"convert_id"`
	StartTime     string `xml:"start_time"`
	EndTime       string `xml:"end_time"`
	MaxIssuePc    string `xml:"max_issue_pc"`
	MaxIssueCount string `xml:"max_issue_count"`
}
