.PHONY: fuzz
fuzz:
    go install github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-build
	go-fuzz-build .
	go-fuzz

.PHONY: corpus
corpus:
	./build-corpus.sh