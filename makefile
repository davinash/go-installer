ifeq ($(OS), Windows_NT)
	EXE=.exe
else
	EXE=
endif

build:
	if [ ! -a $(GOPATH)/bin/binclude ]; then go get -u github.com/lu4p/binclude/cmd/binclude; fi;
	go generate ./...
	go build -o out/installer$(EXE) binclude.go launcher.go
