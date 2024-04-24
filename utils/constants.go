package utils

import (
	"fmt"
	"strings"
)

const (
	CheckTheName = "check the name"
	LoadEnvFromPath = "/home/rdkvx/documents/_projects/KollectionManager/.env"
	Port            = ":3000"
	ServerStatus    = "Running"
	FilterByName = "name =?"
	FilterByDeleted = "deleted = ?"
	DevUpdatedSuccess = "developer updated successfully"
	DevDeletedSuccess = "developer deleted successfully"
)

func LoadEnvErr(envFilePath string, err error) error{
	return	fmt.Errorf("cant load env from path: %s, err {%+v}", envFilePath, err)
}

func FailedTo(operation string, model string, name string) string{
	return fmt.Sprintf("failed to %s %s %s: check the name", operation, model, strings.ToUpper(name))
}