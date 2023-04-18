version=0.1.0
revision=1
bin_name={{cookiecutter.app_name}}

build=$(shell TZ=Asia/Shanghai date +%FT%T%z)
commit=$(shell git rev-parse HEAD | cut -c1-6)
ldflags="-X main.version=${version} -X main.build=${build} -X main.commit=${commit}"

build: prepare
	@echo "building ${bin_name} ${commit}"
	go build -ldflags ${ldflags} -o build/{{cookiecutter.app_name}} cmd/{{cookiecutter.app_name}}/main.go
.PHONY: build

prepare:
	test -e build/conf || mkdir -p build/conf
.PHONY: prepare

clean:
	rm -f build/*
.PHONY: clean
