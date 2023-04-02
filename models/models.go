package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string
	Image         string
	Address       string
	Organizations []*Organization `gorm:"many2many:user_organizations;"`
}

type Transaction struct {
	gorm.Model
	Proposal   Proposal
	ProposalID uint
	Amount     uint
	To         string
}

type Proposal struct {
	gorm.Model
	Title          string
	Body           string
	Organization   Organization
	OrganizationID uint
	Transactions   []Transaction
	Complete       bool
}

type Organization struct {
	gorm.Model
	Name        string
	Creator     *User `gorm:"foreignKey:CreatorID"`
	Image       string
	CreatorID   uint
	Description string
	Proposals   []*Proposal `gorm:"foreignKey:OrganizationID"`
	Members     []*User     `gorm:"many2many:organization_members;"`
}
