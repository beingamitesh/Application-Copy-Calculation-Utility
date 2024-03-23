package models

import "flexeraCodeTest/app/pkg/constants"

type RecordRows struct {
	ComputerID    string
	UserID        string
	ApplicationID string
	ComputerType  constants.ComputerTypes
	Comment       string
}
