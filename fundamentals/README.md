# 基礎Goエラーコース

各章は、バグコミットで標準の`go test`を失敗させ、`main`で同じテストを成功させます。

| # | テーマ | バグコミット | 完成実装 |
| ---: | --- | --- | --- |
| G001 | 破壊的なスライス操作 | `b4317cb` | `ranked_scores.go` |
| G002 | mapの未知キー | `66b8514` | `discounts.go` |
| G003 | 整数除算 | `b60647a` | `average.go` |

## G001: スライスの破壊的な並べ替え

```bash
git checkout b4317cb
go test -run '^TestG001'
git checkout main
```

`sort`は渡されたスライスの配列を直接変更します。入力を守るにはコピーしてから並べ替えます。

## G002: mapの未知キー

```bash
git checkout 66b8514
go test -run '^TestG002'
git checkout main
```

mapの添字アクセスは未知キーでも型のゼロ値を返します。未知の割引コードを拒否する契約では、comma-ok形式で存在を確認します。

## G003: 整数除算

```bash
git checkout b60647a
go test -run '^TestG003'
git checkout main
```

整数同士の除算は端数を切り捨てます。除算前に両方を`float64`へ変換します。

## 全章を実行する

```bash
go test ./...
```

## References

[1] [Learn Go with Tests](https://github.com/quii/learn-go-with-tests)

[2] [The Go Programming Language](https://go.dev/doc/)
