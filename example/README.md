
## BUILD DEPENDENCIES

* GOX - to build Linux binaries on a Mac ( https://github.com/mitchellh/gox )

## BUILDING WITH GO 1.5+

env GOOS=linux GOARCH=amd64 go build -v .
env GOOS=linux GOARCH=arm go build -v .
