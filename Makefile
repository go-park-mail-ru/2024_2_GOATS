# Определение переменных
COVERAGE_FILE = coverage.out
EXCLUDE_FILE = exclude_from_coverage.txt
COVERAGE_HTML = coverage.html
LOCAL_BIN:=$(CURDIR)/bin

all: test coverage report

test:
	go test -coverprofile=$(COVERAGE_FILE) ./...

coverage: test
	./filter_coverage.sh $(COVERAGE_FILE) $(EXCLUDE_FILE)

report:
	go tool cover -func=$(COVERAGE_FILE)

html:
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	open $(COVERAGE_HTML)

clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)


generate-rewiew-api:
	mkdir -p review/pkg/review_v1
	protoc --proto_path review/proto \
	--go_out=review/pkg/review_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=/Users/rugarov/GolandProjects/onboarding-rugarov/bin/protoc-gen-go \
	--go-grpc_out=review/pkg/review_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=/Users/rugarov/GolandProjects/onboarding-rugarov/bin/protoc-gen-go-grpc \
	review/proto/review.proto





#
#COVERAGE_FILE = coverage.out
#EXCLUDE_FILE = exclude_from_coverage.txt
#COVERAGE_HTML = coverage.html
#
#all: test coverage report
#
#test:
#	go test -coverprofile=$(COVERAGE_FILE) ./...
#
#coverage: test
#	./filter_coverage.sh $(COVERAGE_FILE) $(EXCLUDE_FILE)
#
#report:
#	go tool cover -func=$(COVERAGE_FILE)
#
#html:
#	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
#	open $(COVERAGE_HTML)
#
#clean:
#	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)
#
install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

#get-deps:
#	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
#	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
#
#generate-auth-api:
#	mkdir -p auth_service/pkg/auth_v1
#	protoc --proto_path auth_service/proto \
#	--go_out=auth_service/pkg/auth_v1 --go_opt=paths=source_relative \
#	--plugin=protoc-gen-go=/Users/unicoyal/go/bin/protoc-gen-go \
#	--go-grpc_out=auth_service/pkg/auth_v1 --go-grpc_opt=paths=source_relative \
#	--plugin=protoc-gen-go-grpc=/Users/unicoyal/go/bin/protoc-gen-go-grpc \
#	auth_service/proto/review.proto
#
#generate-user-api:
#	mkdir -p user_service/pkg/user_v1
#	protoc --proto_path user_service/proto \
#	--go_out=user_service/pkg/user_v1 --go_opt=paths=source_relative \
#	--plugin=protoc-gen-go=/Users/unicoyal/go/bin/protoc-gen-go \
#	--go-grpc_out=user_service/pkg/user_v1 --go-grpc_opt=paths=source_relative \
#	--plugin=protoc-gen-go-grpc=/Users/unicoyal/go/bin/protoc-gen-go-grpc \
#	user_service/proto/user.proto
