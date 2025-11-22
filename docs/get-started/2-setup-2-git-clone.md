# リモート・リポジトリーからソースをプルするなら

プロジェクトを新規作成するなら、このステップはスキップして構いません。  


## 手順１

go 言語では、ソースコードを置くディレクトリーの構成はある程度決められています。  
プロジェクト・ディレクトリーの置き場所に注意してください。  

```shell
go env GOROOT

        C:\Program Files\Go
```

👆 これは Goルート。プロジェクトではなく、言語仕様のシステム・ライブラリーのようなものです。  


```shell
gomod

        set GO111MODULE=
```

👆 GO111MODULE を on に設定してください。（モジュール・モード）  


## 手順２

📁 `D:\github.com\muzudho\kifuwarabe-uec17-golang-from-uec13`  

👆 例えば、ローカル・リポジトリーは上記のような場所にします。（一例）  

```shell
cd D:\github.com\muzudho\kifuwarabe-uec17-golang-from-uec13
```


## 手順３

```shell
git clone https://github.com/muzudho/kifuwarabe-uec17-golang-from-uec13.git
```

👆 例えば、git を使ってクローンします。  

