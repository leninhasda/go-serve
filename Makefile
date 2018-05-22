build:
	go build -o go-serve .

build-all: clean
	./build.sh

clean:
	rm -rf dist && mkdir dist
