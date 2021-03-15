test:
	go test ./...
test-cover:
	go test ./... -cover
test-cover-file:
	go test ./... -coverprofile=size_coverage.out
test-tool:
	go tool cover -func=size_coverage.out
test-tool-html:
	go tool cover -html=size_coverage.out