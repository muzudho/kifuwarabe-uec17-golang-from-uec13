# Install


## Go言語のインストール

Blog: [Go言語を練習しようぜ☆（＾～＾）？](https://crieit.net/drafts/5ffc46af9214c)  

例えば `C:\go` に Go言語（本体）をインストールしてください。  
例えば 以下のように環境変数を設定してください。  

システム環境変数:  

| Name | Value                    | 別の書き方  |
| ---- | ------------------------ | ----------- |
| Path | C:\go\bin                |             |
| Path | C:\Users\むずでょ\go\bin | %GOPATH%\go |

ユーザー環境変数:  

| Name   | Value                | 別の書き方       |
| ------ | -------------------- | ---------------- |
| GOPATH | C:\Users\むずでょ\go | %USERPROFILE%\go |

（PowerShell ではなく、Command Prompt を使って）以下のコマンドを叩いてください。  

```shell
go version

        #go version go1.15.6 windows/amd64
        go version go1.25.3 windows/amd64

# go言語は、インストール時に、インストール先ディレクトリ（GOROOT）を覚えています。
go env GOROOT

        #c:\go
        C:\Program Files\Go
```


## Modules を使ったプロジェクトの作成

```shell
go mod init github.com/muzudho/kifuwarabe-uec17-golang-from-uec13
```


## Telnet

```shell
# Go言語 は 個人作成の同名のライブラリがいっぱいあるので 一番良さそうなのを検索してください。
go get -v -u github.com/ziutek/telnet
# ↓ これでもいいかも。
# go get -v -u github.com/reiver/go-telnet
```


## Gore

```shell
#go get github.com/motemen/gore/cmd/gore
go install github.com/x-motemen/gore/cmd/gore@latest

# for code completion
go get github.com/mdempsky/gocode

gore -autoimport

        #gore version 0.5.2  :help for help
        gore version 0.6.1  :help for help

gore> fmt.Println("Hello World")

        # 数秒かかるから待つ
        Hello World
        12
        nil
gore>
```

`[Ctrl] + [D]` で終了。  
