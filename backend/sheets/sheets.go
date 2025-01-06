package sheets

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type ShaderDetail struct {
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	Tags        string `json:"tags"`
	Description string `json:"description"`
	Code        string `json:"code"`
}

func getSheetsService() *sheets.Service {
	ctx := context.Background()
	credentialOption := option.WithCredentialsFile("secrets.json")
	scopesOption := option.WithScopes("https://www.googleapis.com/auth/spreadsheets")

	service, err := sheets.NewService(ctx, credentialOption, scopesOption)
	if err != nil {
		log.Fatalf("Failed to create sheets service: %v", err)
	}

	return service
}

func GetShaderDetail(spreadSheetId string, shaderName string) (ShaderDetail, error) {
	details := ShaderDetail{}
	errorMsg := fmt.Errorf("Could not find details for shader: %v", shaderName)

	service := getSheetsService()
	readRange := "Shaders!A2:E"
	response, err := service.Spreadsheets.Values.Get(spreadSheetId, readRange).Do()
	if err != nil {
		return details, errorMsg
	}
	if len(response.Values) == 0 {
		return details, errorMsg
	}

	row, err := search(response.Values, shaderName, 1)
	if err != nil {
		return details, errorMsg
	}

	details.UserId = row[0].(string)
	details.Name = row[1].(string)
	details.Tags = row[2].(string)
	details.Description = row[3].(string)
	details.Code = row[4].(string)

	return details, nil
}

func AppendShaderDetail(spreadSheetId string, sheetId int64, details ShaderDetail) error {
	service := getSheetsService()

	nameExists, errMsg := checkEntryExists(service, spreadSheetId, details.Name, "B2:B", 0)
	if errMsg != nil {
		return errMsg
	}
	if nameExists {
		return fmt.Errorf("Name already exists: %s", details.Name)
	}

	valueRange := sheets.ValueRange{}
	valueRange.Values = make([][]interface{}, 1)
	valueRange.Values[0] = make([]interface{}, 5)
	valueRange.Values[0][0] = details.UserId
	valueRange.Values[0][1] = details.Name
	valueRange.Values[0][2] = details.Tags
	valueRange.Values[0][3] = details.Description
	valueRange.Values[0][4] = details.Code

	_, err := service.Spreadsheets.Values.Append(spreadSheetId, "A:E", &valueRange).ValueInputOption("RAW").Do()
	if err != nil {
		log.Println(err)
		return err
	}

	sortRangeRequest := new(sheets.SortRangeRequest)
	sortRangeRequest.Range = &sheets.GridRange{
		SheetId:          sheetId,
		StartRowIndex:    1,
		StartColumnIndex: 0,
		EndColumnIndex:   5,
	}
	sortRangeRequest.SortSpecs = make([]*sheets.SortSpec, 1)
	sortRangeRequest.SortSpecs[0] = &sheets.SortSpec{
		SortOrder:      "ASCENDING",
		DimensionIndex: 1,
	}

	updateRequest := sheets.BatchUpdateSpreadsheetRequest{}
	updateRequest.Requests = make([]*sheets.Request, 1)
	updateRequest.Requests[0] = &sheets.Request{SortRange: sortRangeRequest}

	_, err2 := service.Spreadsheets.BatchUpdate(spreadSheetId, &updateRequest).Do()
	if err2 != nil {
		return err2
	}

	return nil
}

func checkEntryExists(service *sheets.Service, spreadSheetId string, entry string, sheetRange string, col int) (bool, error) {
	details, err := service.Spreadsheets.Values.Get(spreadSheetId, sheetRange).Do()
	if err != nil {
		return false, err
	}

	_, err2 := search(details.Values, entry, col)
	if err2 != nil {
		return false, nil
	}

	return true, nil
}

func search(values [][]interface{}, search string, col int) ([]interface{}, error) {
	startIndex := 0
	endIndex := len(values) - 1

	for {
		midIndex := (startIndex + endIndex) / 2

		testValue := values[midIndex][col].(string)
		if testValue == search {
			return values[midIndex], nil
		}
		if values[startIndex][col].(string) == search {
			return values[startIndex], nil
		}
		if values[endIndex][col].(string) == search {
			return values[endIndex], nil
		}

		if startIndex == endIndex || startIndex+1 == endIndex {
			return make([]interface{}, 0), fmt.Errorf("Could not find: %s", search)
		}

		if testValue < search {
			startIndex = midIndex
		} else {
			endIndex = midIndex
		}
	}
}
