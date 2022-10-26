#bin/bash
go test -count=1 \
    ./... \
    -coverprofile coverage.out -covermode count 
COVERAGE=`go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+'`
echo "Current test coverage : $COVERAGE %"
                