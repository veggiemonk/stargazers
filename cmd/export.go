// Copyright 2016 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Author: Serhii Kuts (sergeykuc@gmail.com)

package cmd

import (
	"errors"
	"log"

	"github.com/drmegavolt/stargazers/export"
	"github.com/spf13/cobra"
)

// ExportCmd analyzes previously fetched GitHub stargazer data.
var ExportCmd = &cobra.Command{
	Use:   "export --repo=:owner/:repo",
	Short: "export previously analyzed reports to Google Sheets",
	Long: `

Exports generated reports as a SpreadSheet, with each report done as sheet  
`,
	Example: `  stargazers export --repo=cockroachdb/cockroach`,
	RunE:    RunExportToSheets,
}

// RunExportToSheets fetches saved stargazer info for the specified repo and
// runs the analysis reports.
func RunExportToSheets(cmd *cobra.Command, args []string) error {
	if len(Repo) == 0 {
		return errors.New("repository not specified; use --repo=:owner/:repo")
	}

	log.Printf("exporting to Google Sheets GitHub data for repository %s", Repo)
	export.ToSpreadSheet(CacheDir, Repo)
	return nil
}
