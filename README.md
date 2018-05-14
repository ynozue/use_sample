# use_sample
dep で自作ライブラリを利用するためのサンプル

## 使い方
1. dep をインストール
```
$ go get -u github.com/golang/dep/cmd/dep
```
2. 依存を解決
```
$ dep ensure
```
3. 状態を確認
```
$ dep status

PROJECT                       CONSTRAINT     VERSION  REVISION  LATEST         PKGS USED
github.com/ynozue/lib_sample  branch master  v1.0.0   3061ef6   branch master  2
```
4. 実行
```
$ $ go run main.go

subpackage A [hoge]
subpackage B [hoge]
```
