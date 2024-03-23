package helpers

import (
	"bufio"
	"flexeraCodeTest/app/pkg/constants"
	"flexeraCodeTest/app/pkg/models"
	"fmt"
	"log"
	"os"
	"strings"
)

func OpenFile(fileName string) (*bufio.Scanner, func(), error) {
	// opens and scans the csv file
	f, err := os.Open(fileName)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open file with error: %v", err)
	}
	s := bufio.NewScanner(f)
	return s, func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}, nil
}

func FetchRecords(text, applicationId string) (*models.RecordRows, error) {
	record := strings.Split(text, ",")
	// to throw error if columns are more or less than the predefined 5 columns
	if len(record) != 5 {
		return nil, fmt.Errorf("columns are either less or more than the predefined data set")
	}

	// to skip rows which don't belong to the given application id
	if record[2] != applicationId {
		return nil, nil
	}

	// return matching record rows
	return &models.RecordRows{
		ComputerID:    record[0],
		UserID:        record[1],
		ApplicationID: record[2],
		ComputerType:  constants.ComputerTypes(strings.ToUpper(record[3])),
		Comment:       record[4],
	}, nil
}
