## stargazers

illuminate your GitHub community by delving into your repo's stars

### Synopsis

Improoved version of the great spencerkimball/stargazers repo

GitHub allows visitors to star a repo to bookmark it for later
perusal. Stars represent a casual interest in a repo, and when enough
of them accumulate, it's natural to wonder what's driving interest.
Stargazers attempts to get a handle on who these users are by finding
out what else they've starred, which other repositories they've
contributed to, and who's following them on GitHub.

### Features

1. List all stargazers
2. Fetch user info for each stargazer
3. For each stargazer, get list of starred repos & subscriptions
4. For each stargazer subscription, query the repo statistics to
   get additions / deletions & commit counts for that stargazer (use `--advanced=true` flag)
5. Run analyses on stargazer data
6. Export as Google Sheet report (New Comparing to original repo)

### Basic starting point 

```
stargazers :owner/:repo --token=:access_token
```

### To export to google drive you need to have `token.json` file in the root folder. 
To obtain it perform the following steps

1. Create credentials using wizzard referenced in this guide
https://developers.google.com/sheets/api/quickstart/go  
2. Download generated OAuth2 credentials as `client_secret.json` file 
3. Run the app. You will be asked to open link in browser and authorize the app 
4. Once authorized you will be able to download token.json file. This is one time operation. 
5. Next time it will not ask you to follow the link and grant authorization
If `token.json` and `client_secret.json` files are not present the app assumes no google upload is requested


### Examples

#### Download only stargazers + profile info. Save as CSV files
```
  stargazers --repo=cockroachdb/cockroach --token=f87456b1112dadb2d831a5792bf2ca9a6afca7bc
```

#### Download only stargazers + profile info. Save as CSV files and export to Google Sheets

**Important!** Obtain `token.json` file as described above  
```
  stargazers --repo=cockroachdb/cockroach --token=f87456b1112dadb2d831a5792bf2ca9a6afca7bc --id=19agiVeJ-jsn-cbXm2WR4VwAl-fIUaiVK5KOAKKyZpLo
```

#### Download all data (takes a while for big repos). Save as CSV files and export to Google Sheets

**Important!** Obtain `token.json` file as described above  
```
  stargazers --repo=cockroachdb/cockroach --token=f87456b1112dadb2d831a5792bf2ca9a6afca7bc --id=19agiVeJ-jsn-cbXm2WR4VwAl-fIUaiVK5KOAKKyZpLo --advanced=true
```

### Options

```
      --alsologtostderr    logs at or above this threshold go to stderr (default NONE)
  -c, --cache string       directory for storing cached GitHub API responses (default "./stargazer_cache")
      --log-backtrace-at   when logging hits line file:N, emit a stack trace (default :0)
      --log-dir            if non-empty, write log files in this directory (default /var/folders/83/r_nmcwd969g5qc0b7my9wl900000gn/T/)
      --logtostderr        log to standard error instead of files (default true)
      --no-color           disable standard error log colorization
  -r, --repo string        GitHub owner and repository, formatted as :owner/:repo
  -t, --token string       GitHub access token for authorized rate limits
  -i, --id string          Id of Google SpreadSheet. If provided and if has token.json with credentials will publish report to the google file.        
  -a, --adnvanced bool     if true enables advanced long running report like repos starred by followers, contributions of stargazers etc.         
      --verbosity          log level for V logs
      --vmodule            comma-separated list of pattern=N settings for file-filtered logging
```




