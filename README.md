# エラーで学ぶ Go

Goで起きやすい実行時の振る舞いの誤りを、失敗する標準テストから再現し、最小修正と回帰テストで学ぶ教材です。各章の`main`は成功状態で、失敗する状態はGit履歴に残します。

## 開始方法

```bash
git clone https://github.com/tonbiattack/error-driven-go-lab.git
cd error-driven-go-lab
go test ./...
```

## 基礎コース

| # | テーマ | バグコミット |
| ---: | --- | --- |
| G001 | 破壊的なスライス操作 | `b4317cb` |
| G002 | mapの未知キー | `66b8514` |
| G003 | 整数除算 | `b60647a` |

Red → Greenの詳細は[`fundamentals/README.md`](fundamentals/README.md)を参照してください。

## 検証

```bash
gofmt -w *.go
go vet ./...
go test ./...
```

## 文書

| 文書 | 内容 |
| --- | --- |
| [SUMMARY.md](SUMMARY.md) | コース目次 |
| [DESIGN.md](DESIGN.md) | Go固有の教材設計 |
| [coverage-matrix.md](coverage-matrix.md) | 実装済み・未着手テーマ |

## References

[1] [Learn Go with Tests](https://github.com/quii/learn-go-with-tests)

[2] [The Go Programming Language](https://go.dev/doc/)
