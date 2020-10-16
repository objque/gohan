#!/usr/bin/env bash

if ! which revive > /dev/null; then
    echo "==> Installing revive"
    go get -u github.com/mgechev/revive@master
fi

echo "==> Checking revive..."
revive --config ./.revive.toml --formatter stylish ./cmd/...
revive --config ./.revive.toml --formatter stylish ./internal/...
revive --config ./.revive.toml --formatter stylish ./pkg/...
