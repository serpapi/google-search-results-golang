version=3.0

all: test

test:
	go test -v .

oobt:
	cd demo ; go run demo.go

release: oobt
	git tag -a ${version}
	git push origin ${version}
	@echo "create release: ${version}"

