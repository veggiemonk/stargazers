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
// Author: Spencer Kimball (spencer.kimball@gmail.com)

package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// AccessToken is used to access repo stars and gain non-authorized
// rate limits.
var AccessToken string

// AccessTokenDesc describes usage.
const AccessTokenDesc = "GitHub access token for authorized rate limits"

func getAccessToken() (string, error) {
	if len(AccessToken) == 0 {
		return "", errors.New(`An access token must be specified via --token.

To generate an access token for accessing repo stars and gaining authorized
rate limits, see:

https://help.github.com/articles/creating-an-access-token-for-command-line-use/

When creating a token, ensure that only the public_repo permission is enabled.
`)
	}
	return AccessToken, nil
}

// CacheDir specifies where to store cached JSON responses.
var CacheDir string

// CacheDirDesc describes usage.
const CacheDirDesc = "directory for storing cached GitHub API responses"

// Repo specifies the the owner and repository in :owner/:repo format.
var Repo string

// ReportTypes enabled reports to be generated
var ReportTypes *[]string

// SpreadSheetID Google drive of folder to move file (optional)
var SpreadSheetID string

// SpreadSheetIDDesc describes usage of folder-id.
const SpreadSheetIDDesc = "Google SpreadSheet id to update"

// RepoDesc describes usage.
const RepoDesc = "GitHub owner and repository, formatted as :owner/:repo"

func mustUsage(cmd *cobra.Command) {
	if err := cmd.Usage(); err != nil {
		panic(err)
	}
}
