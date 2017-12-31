# go-lang
Exploration of the Go language

# Setup on Unix System
- Download and install the correct go package for your operating system
- Create the necessary folders for the below commands
- Add these to your ~/.bash_profile 
  - `export PATH = $PATH:/usr/local/go/bin`
  - (optional) `export GOPATH=$HOME/go`
  - (optional) `export GOBIN=$HOME/go/bin`
- run the command `source ~/.bash_profile`
- `cd $HOME/go` and `mkdir src; mkdir hello; cd src/hello`
- copy the file `hello.go` from this repo and into `src/hello/`
- run `go build` and you should see a `hello` executable
- run the executable with `./hello`

# Running go commands
- `go run <myfile.go>` will run the go file
- `gofmt -w <file.go>` will format code

- concurrency - simple concurrency with a go routine
- cryptofun - todo
- fibonacci - simple memoization and tabulation using fibonacci
- function - random generate two numbers and determine larger number
- hashcollider - testing hash comparison and hash searching
- pi_montecarlo - montecarlo solver to determine value of pi
- swap - swapping without temp variable
