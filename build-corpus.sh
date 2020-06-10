#!/bin/bash

rm corpus/*-bin

OSS="darwin linux"

for OS in $OSS; do
  GOOS=darwin GOARCH=amd64 find src -name '*.go' \
    -execdir go build -o ../corpus/{}-${OS}-amd64-standard-bin {} \; \
    -execdir go build -gcflags='-l' -o ../corpus/{}-${OS}-amd64-no-inline-bin {} \; \
    -execdir go build -gcflags='-N' -o ../corpus/{}-${OS}-amd64-no-optimise-bin {} \; \
    -execdir go build -gcflags='-l -N' -o ../corpus/{}-${OS}-amd64-no-inline-optimise-bin {} \; \
    -execdir go build -ldflags='-s' -o ../corpus/{}-${OS}-amd64-no-sybmols-bin {} \; \
    -execdir go build -ldflags='-w' -o ../corpus/{}-${OS}-amd64-no-dwarf-bin {} \; \
    -execdir go build -ldflags='-s -w' -o ../corpus/{}-${OS}-amd64-no-symbols-dwarf-bin {} \; \
    -execdir go build -ldflags='-s -w' -gcflags='-l -N' -o ../corpus/{}-${OS}-amd64-no-all-bin {} \;
done