package sheets

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type Response struct {
	Name string `json:"name"`
	Code string `json:"code"`
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

func GetSheetData() []Response {
	service := getSheetsService()
	spreadSheetId := "1kReHyEqnuPh9QTlSt7FeVno9DDn6Qfc0hpbYlzpC6hc"
	readRange := "Sheet1!A2:E"
	response, err := service.Spreadsheets.Values.Get(spreadSheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	details := make([]Response, 0)
	for _, row := range response.Values {
		var temp Response
		temp.Name = row[0].(string)
		temp.Code = row[1].(string)
		details = append(details, temp)
	}

	return details
}
