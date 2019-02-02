@echo off
setlocal enabledelayedexpansion
set file=%cd%\test.ini
set file_tmp=%cd%\test_tmp.ini
set file_bak=%cd%\test_bak.ini 
set sourceStr=·½°¸
set replacedStr=999

for /f "delims=" %%i in (%file%) do (
    set str=%%i
        set "str=!str:%sourceStr%=%replacedStr%!" 
        echo !str!>>%file_tmp%
)
copy "%file%" "%file_bak%" >nul 2>nul
move "%file_tmp%" "%file%"