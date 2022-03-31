// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Account struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	ResumeURL   string      `json:"resumeUrl"`
	Skills      []string    `json:"skills"`
	Bio         string      `json:"bio"`
	Description string      `json:"description"`
	AccountType AccountType `json:"accountType"`
	TimeAdded   time.Time   `json:"timeAdded"`
}

type AuthResponse struct {
	Account *Account `json:"account"`
}

type Job struct {
	ID          string    `json:"id"`
	Role        string    `json:"role"`
	Location    string    `json:"location"`
	CompanyName string    `json:"companyName"`
	JobType     string    `json:"jobType"`
	Skills      []string  `json:"skills"`
	Description string    `json:"description"`
	TimeAdded   time.Time `json:"timeAdded"`
}

type NewAccount struct {
	Name        string      `json:"name"`
	Email       string      `json:"email"`
	Password    string      `json:"password"`
	ResumeURL   string      `json:"resumeUrl"`
	Skills      []string    `json:"skills"`
	Bio         string      `json:"bio"`
	AccountType AccountType `json:"accountType"`
	Description string      `json:"description"`
}

type NewJob struct {
	Role        string   `json:"role"`
	Location    string   `json:"location"`
	CompanyName string   `json:"companyName"`
	JobType     string   `json:"jobType"`
	Skills      []string `json:"skills"`
	Description string   `json:"description"`
}

type NewOrganization struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

type AccountType string

const (
	AccountTypeStudent      AccountType = "Student"
	AccountTypeOrganization AccountType = "Organization"
)

var AllAccountType = []AccountType{
	AccountTypeStudent,
	AccountTypeOrganization,
}

func (e AccountType) IsValid() bool {
	switch e {
	case AccountTypeStudent, AccountTypeOrganization:
		return true
	}
	return false
}

func (e AccountType) String() string {
	return string(e)
}

func (e *AccountType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AccountType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AccountType", str)
	}
	return nil
}

func (e AccountType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
