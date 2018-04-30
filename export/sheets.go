/**
 * @license
 * Copyright Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
// [START sheets_quickstart]
package export

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}

// ToSpreadSheet uploads all csv files under rootPath/repo as sheets in a SpreadSheet
func ToSpreadSheet(rootPath string, repo string) {
	var reportFiles []string
	filepath.Walk(rootPath, func(p string, f os.FileInfo, _ error) error {
		if filepath.Ext(p) == ".csv" {
			reportFiles = append(reportFiles, f.Name())
		}
		return nil
	})

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved client_secret.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	fmt.Println(reportFiles)
	sID := "19agiVeJ-jsn-cbXm2WR4VwAl-fIUaiVK5KOAKKyZpLo"

	// newSheets := make([]*sheets.Sheet, 0)
	for _, f := range reportFiles {
		bsr := sheets.BatchUpdateSpreadsheetRequest{
			Requests: []*sheets.Request{
				&sheets.Request{
					AddSheet: &sheets.AddSheetRequest{
						Properties: &sheets.SheetProperties{
							Title: f,
						},
					},
				},
			},
		}
		_, er := srv.Spreadsheets.BatchUpdate(sID, &bsr).Do()
		if er != nil {
			fmt.Println("Tab Exists: " + f)
		}
		sheetRange := f + "!A1:Z10000"
		clearRes, _ := srv.Spreadsheets.Values.Clear(sID, sheetRange, &sheets.ClearValuesRequest{}).Do()
		fmt.Println(clearRes)

		valueRange := sheets.ValueRange{
			Values: buildRowsFromCsv(path.Join(rootPath, repo, f)),
		}
		valueInputOption := "USER_ENTERED"
		valueRes, err := srv.Spreadsheets.Values.Update(sID, sheetRange, &valueRange).ValueInputOption(valueInputOption).Do()
		fmt.Println(valueRes, err)
	}
}

func buildRowsFromCsv(csvPath string) [][]interface{} {
	fmt.Println("reading " + csvPath)
	f, _ := os.Open(csvPath)
	r := csv.NewReader(bufio.NewReader(f))
	result := make([][]interface{}, 0)
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		row := make([]interface{}, 0)
		for value := range record {
			row = append(row, record[value])
		}
		result = append(result, row)

	}
	return result
}
