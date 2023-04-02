package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Address string
}

type Transaction struct {
	gorm.Model
	Proposal Proposal
	ProposalID uint
	Amount uint
	To string
}

type Proposal struct {
	gorm.Model
	Title string
	Body string
	Organization Organization
	OrganizationID uint
	Transactions []Transaction
	Complete bool
}


type Organization struct{
	gorm.Model
	Name string
	Creator User
	CreatorID uint
	Description string
	Proposals []Proposal

}
