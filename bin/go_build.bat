@echo off

rem 文字化け対策。コマンドプロンプトがデフォルトで Shift-JIS なので、その文字コードを消すことで、UTF-8 にする。
chcp 65001 >nul

echo go build するぜ（＾～＾）...
cd ..
go build
cd ./bin
echo go build したぜ（＾～＾）！

call move_file_kifuwarabe.bat
