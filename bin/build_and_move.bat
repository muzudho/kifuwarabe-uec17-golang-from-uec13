@echo off

rem 文字化け対策。コマンドプロンプトがデフォルトで Shift-JIS なので、その文字コードを消すことで、UTF-8 にする。
chcp 65001 >nul

echo 全部任せろだぜ（＾～＾）...

call subroutines/build.bat

call subroutines/move_file_kifuwarabe.bat

echo すべて終わったぜ（＾～＾）！
pause
