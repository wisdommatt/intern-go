// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Account interface {
	IsAccount()
}

type AuthResponse struct {
	AuthToken string  `json:"authToken"`
	Account   Account `json:"account"`
}

type Job struct {
	ID          string    `json:"id"`
	Role        string    `json:"role"`
	Location    string    `json:"location"`
	CompanyName string    `json:"companyName"`
	JobType     string    `json:"jobType"`
	Skills      []string  `json:"skills"`
	Description string    `json:"description"`
	EndDate     time.Time `json:"endDate"`
	TimeAdded   string    `json:"timeAdded"`
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

type NewStudent struct {
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	ResumeURL string   `json:"resumeUrl"`
	Skills    []string `json:"skills"`
	Bio       string   `json:"bio"`
}

type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Description string `json:"description"`
}

func (Organization) IsAccount() {}

type Pagination struct {
	AfterID *string `json:"afterId"`
	//  default pagination limit is 100
	Limit *int `json:"limit"`
}

type Student struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	ResumeURL string   `json:"resumeUrl"`
	Skills    []string `json:"skills"`
	Bio       string   `json:"bio"`
}

func (Student) IsAccount() {}

type AccountCategory string

const (
	AccountCategoryStudent      AccountCategory = "Student"
	AccountCategoryOrganization AccountCategory = "Organization"
)

var AllAccountCategory = []AccountCategory{
	AccountCategoryStudent,
	AccountCategoryOrganization,
}

func (e AccountCategory) IsValid() bool {
	switch e {
	case AccountCategoryStudent, AccountCategoryOrganization:
		return true
	}
	return false
}

func (e AccountCategory) String() string {
	return string(e)
}

func (e *AccountCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AccountCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AccountCategory", str)
	}
	return nil
}

func (e AccountCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
