/*
Copyright IBM Corp. 2022 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package tenable

type User struct {
	ID                 interface{} `json:"id"`
	Status             string      `json:"status,omitempty"`             // "0",
	Username           string      `json:"username,omitempty"`           // "admin",
	LDAPUsername       string      `json:"ldapUsername,omitempty"`       // "",
	Firstname          string      `json:"firstname,omitempty"`          // "Admin",
	Lastname           string      `json:"lastname,omitempty"`           // "User",
	Title              string      `json:"title,omitempty"`              // "Application Administrator",
	Email              string      `json:"email,omitempty"`              // "",
	Address            string      `json:"address,omitempty"`            // "",
	City               string      `json:"city,omitempty"`               // "",
	State              string      `json:"state,omitempty"`              // "",
	Country            string      `json:"country,omitempty"`            // "",
	Phone              string      `json:"phone,omitempty"`              // "",
	Fax                string      `json:"fax,omitempty"`                // "",
	CreatedTime        string      `json:"createdTime,omitempty"`        // "1432921843",
	ModifiedTime       string      `json:"modifiedTime,omitempty"`       // "1453473716",
	LastLogin          string      `json:"lastLogin,omitempty"`          // "1454350174",
	LastLoginIP        string      `json:"lastLoginIP,omitempty"`        // "172.20.0.0",
	MustChangePassword string      `json:"mustChangePassword,omitempty"` // "false",
	Locked             string      `json:"locked,omitempty"`             // "false",
	FailedLogin        string      `json:"failedLogins,omitempty"`       // "0",
	AuthType           string      `json:"authType,omitempty"`           // "tns",
	Fingerprint        string      `json:"fingerprint,omitempty"`        // null,
	Password           string      `json:"password,omitempty"`           // "SET",
	Preferences        []PreferenceItem
	Organization       Organization
	OrgName            string `json:"orgName,omitempty"`
	UserPrefs          []UserPrefItem
	UUID               string `json:"uuid,omitempty"`
}

type PreferenceItem struct {
	Name  string
	Value string
	Tag   string
}

type Organization struct {
	ID          interface{}
	Name        string
	Description string
}

type UserPrefItem struct {
	Name  string
	Value string
	Tag   string
}

type Role struct {
	ID          interface{}
	Name        string
	Description string
}
