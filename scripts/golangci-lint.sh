#!/usr/bin/env bash

if ! which golangci-lint > /dev/null; then
    echo "==> Installing golangci-lint"
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.31.0
fi

echo "==> Checking golangci-ling..."
golangci-lint -v run
