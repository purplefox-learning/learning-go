# Package Structure

as per convention
* the main package is found in the `cmd` director at the root level


# How to run 
for all the following commands, by default we run it at the root directory i.e. `/Users/ycm/go/src/github.com/purplefox81/learning-go`
* run `go build ./...` or run `go build` in a specific dir. `go build` wont produce an output file, instead it saves the compiled package in the local build cache
* run `go install` in the `hello-cmd` directory, or run a full command, `go install github.com/purplefox81/learning-go`. It produces an executable called `hello-cmd` in `~/go/bin` which is exactly the same as the directory name containing the main package. run `hello-cmd` to verify the binary is working

# Download/Install 3rd party package dependencies
* `go get xxx` e.g. `go get github.com/jinzhu/gorm`
* 

# References

for more Go commands, see `https://golang.org/cmd/go/`
for more Go best practices, see ``
for more Go Test, see ``
for ...