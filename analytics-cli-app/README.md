# CLI Tool to Analyze Log Data

## How to run the app:
- There are two Go code to run, one to generate dummy log data (generate) and the other one is the main app (analytics)
- Go to folder "generate", run: `cd generate/`
- In main.go, there are tweakable constants, that are NUM_OF_FILES (specifying number of files to be generated) and NUM_OF_LINES (specifying number of log lines in a log file)
- Generate the dummy data by running command: `go run main.go`
- The log files will be stored inside analytics directory
- Go to folder "analytics", run: `cd ../analytics` (from generate folder)
- Build the CLI app by running: `go build`
- Make it runnable from command line by copy the executable file into /usr/local/bin, run: `cp analytics /usr/local/bin`
- Now, run this command to get the last 10 minutes worth of log data, run: `analytics -t 10m -d ./logs`

## Regarding tests:
- If you're in "generate" folder, run the tests by running command: `go test -v ./...`
- Still in the "generate" folder, run the benchmarking by running command: `go test -bench=Bench`
- Before running tests inside the analytics folder make sure you have the log files inside /logs. Otherwise, no data will be read and the tests won't pass
- Inside the "analytics" folder, run the run the tests by running command: `go test -v ./...`
- Inside the "analytics" folder, run the benchmarking by running command: `go test -bench=Bench`
