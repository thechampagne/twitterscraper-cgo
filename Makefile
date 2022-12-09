CGO := go build
STATIC := -buildmode=c-archive
SHARED := -buildmode=c-shared
LIBS := build/twitterscraper.a build/twitterscraper.so

.PHONY: all
all: $(LIBS)

build/twitterscraper.a: twitterscraper.go
	$(CGO) $(STATIC) -o build/twitterscraper.a $<

build/twitterscraper.so: twitterscraper.go
	$(CGO) $(SHARED) -o build/twitterscraper.so $<

.PHONY: clean
clean:
	find build -type f \( -name '*.h' -o -name '*.so' -o -name '*.a' \) -delete
