# Changelog

## [v1.4.2] - 01 August 2025

    * Use go modules
    * Fix refresh crashing some applications because wParam was not NULL

## [v1.4.1] - 04 June 2019

    * Fix Panic and data race calling a Windows function in user32.dll (golang #31741)
    * Appveyor hangs when running refresh.exe (removed that step) 

## [v1.3.0] - 19 July 2017

    * Replaced windows pop-up by sending WM_SETTINGCHANGED instead

## [v1.2.0] - 26 July 2015

    * Fixed keys not sent correctly
    * Added `-delay` to specify delay before sending keys
    * Removed dependencies to `github.com/AllenDang/w32` and `gcc`

## [v1.1.0] - 14 July 2015

    * Added `-title` option to specify window title for localized versions of windows
    * Cleaned up code
    * Note: this version is broken (forgot to send KEYEVENTF_KEYUP)

## [v1.0.0] - 12 June 2015

    * First version
