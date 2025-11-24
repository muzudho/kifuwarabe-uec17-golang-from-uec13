# ビルドの前に

（PowerShell ではなく、Command Prompt を使って）以下のコマンドを叩いてください。  


## 手順１

もし git を使っていて、信頼性が確認できていないドライブ、例えば外付けの NAS などをローカル・リポジトリーにしているときは、  
安全なディレクトリーである旨を登録する必要があります。  

一例：  

```shell
git config --global --add safe.directory Z:/muzudho-github.com/muzudho/kifuwarabe-uec17-golang-from-uec13
```


## 手順２

```shell
go mod tidy
```

👆 ソースコードの掃除。使っていないパッケージを削除。 📄 `go.mod` と 📄 `go.sum` ファイルを更新します。  
