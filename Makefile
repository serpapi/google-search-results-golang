version=3.0

all: test

check:
	go vet .
	go fmt .

test:
	go test -v .

# check that everything is pushed
package:
	git status | grep "Nothing"

oobt:
	mkdir -p /tmp/oobt
	cp demo/demo.go /tmp/oobt
	cd /tmp/oobt ; \
		go get -u github.com/serpapi/google_search_results_golang ; \
		go run demo.go

release: oobt
	git tag -a ${version}
	git push origin ${version}
	@echo "create release: ${version}"

