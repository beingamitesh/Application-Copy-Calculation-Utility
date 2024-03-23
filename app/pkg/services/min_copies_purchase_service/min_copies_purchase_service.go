package minCopiesPurchaseService

import (
	"flexeraCodeTest/app/pkg/constants"
	"flexeraCodeTest/app/pkg/helpers"
	"flexeraCodeTest/app/pkg/models"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

// calculate the minimum copies of the application with ID 374 a company must purchase from a csv file
func MinCopiesPurchase(csvFileName string) (int, error) {
	duplicateRecordCheckerMap := make(map[string]int)
	userDevicesCountMap := make(map[string]map[string]int)

	s, close, err := helpers.OpenFile(csvFileName)
	if err != nil {
		return 0, err
	}
	defer close()

	// skipping first row, which is header
	s.Scan()

	// mutex avoids different go routines working on a critical section, i.e. race condition
	m := &sync.Mutex{}

	// to spawn n number of goroutines
	g := new(errgroup.Group)

	// read records
	for s.Scan() {
		err = s.Err()
		if err != nil {
			return 0, fmt.Errorf("error while scanning file, error: %v", err)
		}
		// variable defined to store text from the row
		text := s.Text()
		g.Go(func() error {
			// to only get the records belonging to application id 374
			record, err := fetchRecords(text, constants.ApplicationID)
			if err != nil {
				return err
			}

			m.Lock()
			// check duplicate and update userDevicesCountMap
			if record != nil {
				userDevicesCountMap = checkAndUpdateUserDevicesCountMap(duplicateRecordCheckerMap, record, userDevicesCountMap)
			}
			m.Unlock()
			return nil
		})
	}

	// waiting for all goroutines to finish execution
	err = g.Wait()
	if err != nil {
		return 0, fmt.Errorf("error while executing goroutines, error: %v", err)
	}

	return minApplicationRequiredByUserWithDeviceMapping(userDevicesCountMap)
}

func fetchRecords(text string, applicationId string) (*models.RecordRows, error) {
	record, err := helpers.FetchRecords(text, applicationId)
	if err != nil {
		return nil, err
	} else if record == nil {
		return nil, nil
	}
	return record, err
}

func checkAndUpdateUserDevicesCountMap(duplicateRecordCheckerMap map[string]int, record *models.RecordRows, userDevicesCountMap map[string]map[string]int) map[string]map[string]int {
	// to ignore duplicates treating Column Id as unique
	if _, ok := duplicateRecordCheckerMap[record.ComputerID]; !ok {
		duplicateRecordCheckerMap[record.ComputerID] = 1

		// initialising user device count map for each user
		if _, ok := userDevicesCountMap[record.UserID]; !ok {
			userDevicesCountMap[record.UserID] = make(map[string]int)
		}

		// counting total no of desktops and laptops each user has
		switch record.ComputerType {
		case constants.DESKTOP:
			userDevicesCountMap[record.UserID][string(constants.DESKTOP)]++
		case constants.LAPTOP:
			userDevicesCountMap[record.UserID][string(constants.LAPTOP)]++
		}
	}
	return userDevicesCountMap
}

// returns minimum copies purchased by company by adding the max of the count of ComputerType of each user
func minApplicationRequiredByUserWithDeviceMapping(usersDeviceCountMap map[string]map[string]int) (int, error) {
	var totalCount int
	for _, val := range usersDeviceCountMap {
		// to add max of count of ComputerType
		if val[string(constants.DESKTOP)] > val[string(constants.LAPTOP)] {
			totalCount = totalCount + val[string(constants.DESKTOP)]
		} else {
			totalCount = totalCount + val[string(constants.LAPTOP)]
		}
	}
	return totalCount, nil
}
