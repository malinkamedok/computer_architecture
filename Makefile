
all: tests compiler vrt
	./comp -f file.asm
	./vm -f result.json

tests: compiler_test vrt_test

compiler:
	go build -o comp ./cmd/comp/main.go

vrt:
	go build -o vm ./cmd/vm/main.go

compiler_test:
	go test internal/comp_test/trans_test.go

vrt_test:
	go test internal/vm_test/vm_test.go

clean:
	rm comp vm result.json