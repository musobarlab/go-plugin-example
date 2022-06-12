.PHONY : build build-plugins clean

build-plugins:
	@echo "building plugins"
	@go build -buildmode=plugin -o plugins/consoledumper/consoledumper.so plugins/consoledumper/console.go \
	&& go build -buildmode=plugin -o plugins/filedumper/filedumper.so plugins/filedumper/file.go

build:
	@echo "building binary"
	@go build -o runner  github.com/musobarlab/go-plugin-example/cmd/runner

clean:
	@echo "cleaning unused file"
	rm -rf runner \
	&& rm -rf *.txt