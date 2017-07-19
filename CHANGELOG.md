## Changelog

##### refresh 1.3.0 - (19 July 2017)

    * Replaced windows pop-up by sending WM_SETTINGCHANGED instead

##### refresh 1.2.0 - (26 July 2015)

    * Fixed keys not sent correctly
    * Added `-delay` to specify delay before sending keys
    * Removed dependencies to `github.com/AllenDang/w32` and `gcc`

##### refresh 1.1.0 - (14 July 2015)

    * Added `-title` option to specify window title for localized versions of windows
    * Cleaned up code
    * Note: this version is broken (forgot to send KEYEVENTF_KEYUP)

##### refresh 1.0.0 - (12 June 2015)

    * First version
