package sheets

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/google/uuid"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type ShaderDetail struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	UserName     string `json:"userName"`
	Name         string `json:"name"`
	Tags         string `json:"tags"`
	Description  string `json:"description"`
	Code         string `json:"code"`
	CreationDate int    `json:"creationDate"`
}

type ShaderDetailRequest struct {
	Email        string `json:"email"`
	UserName     string `json:"userName"`
	Name         string `json:"name"`
	Tags         string `json:"tags"`
	Description  string `json:"description"`
	Code         string `json:"code"`
	CreationDate int    `json:"creationDate"`
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

func GetShaderDetail(spreadSheetId string, shaderId string) (ShaderDetail, error) {
	errorMsg := fmt.Errorf("Could not find details for shader: %v", shaderId)

	service := getSheetsService()
	readRange := "Shaders!A2:H"
	response, err := service.Spreadsheets.Values.Get(spreadSheetId, readRange).Do()
	if err != nil {
		return ShaderDetail{}, errorMsg
	}
	if len(response.Values) == 0 {
		return ShaderDetail{}, errorMsg
	}

	row, err := search(response.Values, shaderId, 0)
	if err != nil {
		return ShaderDetail{}, errorMsg
	}

	details := getShaderDetailObject(row)
	return details, nil
}

func GetShaderDetails(spreadSheetId string, startRow int, endRow int) ([]ShaderDetail, error) {
	details := make([]ShaderDetail, 0)
	service := getSheetsService()
	readRange := fmt.Sprintf("Shaders!A%d:H%d", startRow, endRow)
	response, err := service.Spreadsheets.Values.Get(spreadSheetId, readRange).Do()
	if err != nil {
		return details, err
	}
	if len(response.Values) == 0 {
		return details, nil
	}

	for _, row := range response.Values {
		temp := getShaderDetailObject(row)
		details = append(details, temp)
	}

	return details, nil
}

func getShaderDetailObject(row []interface{}) ShaderDetail {
	creationDate, err := strconv.Atoi(row[7].(string))
	if err != nil {
		creationDate = 0
	}
	temp := ShaderDetail{
		Id:           row[0].(string),
		Email:        row[1].(string),
		UserName:     row[2].(string),
		Name:         row[3].(string),
		Tags:         row[4].(string),
		Description:  row[5].(string),
		Code:         row[6].(string),
		CreationDate: creationDate,
	}

	return temp
}

func AppendShaderDetail(spreadSheetId string, sheetId int, details ShaderDetailRequest) (ShaderDetail, error) {
	service := getSheetsService()

	shaderId := uuid.NewString()
	valueRange := sheets.ValueRange{}
	valueRange.Values = make([][]interface{}, 1)
	valueRange.Values[0] = make([]interface{}, 8)
	valueRange.Values[0][0] = shaderId
	valueRange.Values[0][1] = details.Email
	valueRange.Values[0][2] = details.UserName
	valueRange.Values[0][3] = details.Name
	valueRange.Values[0][4] = details.Tags
	valueRange.Values[0][5] = details.Description
	valueRange.Values[0][6] = details.Code
	valueRange.Values[0][7] = details.CreationDate

	_, err := service.Spreadsheets.Values.Append(spreadSheetId, "A:H", &valueRange).ValueInputOption("RAW").Do()
	if err != nil {
		log.Println(err)
		return ShaderDetail{}, err
	}

	sortRangeRequest := new(sheets.SortRangeRequest)
	sortRangeRequest.Range = &sheets.GridRange{
		SheetId:          int64(sheetId),
		StartRowIndex:    1,
		StartColumnIndex: 0,
		EndColumnIndex:   8,
	}
	sortRangeRequest.SortSpecs = make([]*sheets.SortSpec, 1)
	sortRangeRequest.SortSpecs[0] = &sheets.SortSpec{
		SortOrder:      "ASCENDING",
		DimensionIndex: 0,
	}

	updateRequest := sheets.BatchUpdateSpreadsheetRequest{}
	updateRequest.Requests = make([]*sheets.Request, 1)
	updateRequest.Requests[0] = &sheets.Request{SortRange: sortRangeRequest}

	_, err2 := service.Spreadsheets.BatchUpdate(spreadSheetId, &updateRequest).Do()
	if err2 != nil {
		return ShaderDetail{}, err2
	}

	updateDetail := ShaderDetail{
		Id:           shaderId,
		Email:        details.Email,
		UserName:     details.UserName,
		Name:         details.Name,
		Tags:         details.Tags,
		Description:  details.Description,
		Code:         details.Code,
		CreationDate: details.CreationDate,
	}

	return updateDetail, nil
}

func checkEntryExists(service *sheets.Service, spreadSheetId string, entry string, sheetRange string, col int) (bool, error) {
	details, err := service.Spreadsheets.Values.Get(spreadSheetId, sheetRange).Do()
	if err != nil {
		return false, err
	}
	if len(details.Values) == 0 {
		return false, nil
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
