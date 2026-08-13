#!/usr/bin/env bash
set -euo pipefail

for file in README.md SUMMARY.md DESIGN.md coverage-matrix.md fundamentals/README.md; do
  test -s "$file" || { echo "必要な教材ファイルがありません: $file" >&2; exit 1; }
done

git diff --check
gofmt -w *.go
go vet ./...
go test ./...

echo "Goエラー学習コースの検証に成功しました。"
