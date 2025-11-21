@echo off

rem 文字化け対策。コマンドプロンプトがデフォルトで Shift-JIS なので、その文字コードを消すことで、UTF-8 にする。
chcp 65001 >nul

echo 全部任せろだぜ（＾～＾）...

rem プロジェクト・ルートで作業する。
cd ..

rem プロジェクト・ルートに 📁 `go-to-championship` フォルダーを作成。
mkdir go-to-championship

rem さっきビルドしてルート・ディレクトリに生成した 📄 `{アプリケーション名}.exe` ファイルを 📁 `go-to-championship` フォルダーへ移動。
move /Y .\kifuwarabe-uec17-golang-from-uec13.exe .\go-to-championship\

rem プロジェクト・ルートにある 📁 `input` フォルダーには設定ファイルが入ってるから、フォルダーを丸ごと 📁 `go-to-championship` フォルダーへコピー。
rem     /E サブディレクトリ（空のディレクトリーも含める）コピーするオプション。
rem     /I コピー先がフォルダーの場合に確認しないオプション。
rem     /Y 上書き確認をしないオプション。
xcopy /E /I /Y .\input .\go-to-championship\input

rem 📁 `go-to-championship` フォルダーの中に 📁 `logs` フォルダー作成。中に 📄 `.gitkeep` ファイルを作成して、空のフォルダーを Git 管理できるようにする。
mkdir go-to-championship\logs
echo. > go-to-championship\logs\.gitkeep

echo よし、できたぜ（＾～＾）！
