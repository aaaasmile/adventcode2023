#### Running

1. Create `daynN/test.txt`
2. Create `dayN/input.txt`
3. Run: `go run dayN/main.go -part <1 or 2> [-test]`

`-test` will run the `test.txt` inputs instead of `input.txt`

#### Compiling and running

1. Create `daynN/test.txt` 
2. Create `dayN/input.txt`
3. `mkdir bin/`
4. Build: `go build -o bin/dayN dayN/main.go`
5. Run: `bin/dayN -part <1 or 2> [-test]`

`-test` will run the `test.txt` inputs instead of `input.txt`

Make sure that you use LF instead of CRLF in test.txt and input.txt

#### Template a day directory

Run `./create-day.sh <N>` where `N` is the day number (under windows use WSL or create a Powershell Script)

## Credits
Copied (first 3 dfays) and modified from https://github.com/viliusan/aoc2023-go and https://github.com/Stogas/aoc2023-go