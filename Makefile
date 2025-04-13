short-test:
	go test -short -v ./...

test:
	mkdir -p coverage/covdata
# Use the new binary format to ensure cross-package calls are counted towards coverage
# https://go.dev/blog/integration-test-coverage
# -p 1 disable parallel testing in favor of streaming log output - https://github.com/golang/go/issues/24929#issuecomment-384484654
	go test -race -cover -covermode atomic -coverpkg ./... -v -vet=all -timeout 15m -p 1\
		./... \
		-args -test.gocoverdir="${PWD}/coverage/covdata" \
		| ts -s
# NB: ts command requires moreutils package; awk trick from https://stackoverflow.com/a/25764579 doesn't stream output
	go tool covdata percent -i=./coverage/covdata
	# Convert to old text format for coveralls upload
	go tool covdata textfmt -i=./coverage/covdata -o ./coverage/covprofile
	go tool cover -html=./coverage/covprofile -o ./coverage/coverage.html

checks: check_changes check_tidy

check_changes:
# make sure .next.version contains the intended next version
# if the following fails, update either the next version or undo any unintended api changes
	go run golang.org/x/exp/cmd/gorelease@latest -version $(shell cat .next.version)

check_tidy:
	go mod tidy
	# Verify that `go mod tidy` didn't introduce any changes. Run go mod tidy before pushing.
	git diff --exit-code --stat go.mod go.sum
