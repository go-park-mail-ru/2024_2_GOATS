# Определение переменных
COVERAGE_FILE = coverage.out
EXCLUDE_FILE = exclude_from_coverage.txt
COVERAGE_HTML = coverage.html

all: test coverage report html

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
