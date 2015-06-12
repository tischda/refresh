# refresh [![Build status](https://ci.appveyor.com/api/projects/status/ok7detq1hwnbd5cc?svg=true)](https://ci.appveyor.com/project/tischda/refresh)

Windows utility written in [Go](https://www.golang.org) to refresh
SYSTEM environment variables.

I tried https://support.microsoft.com/en-us/kb/104011 and restarted Explorer,
it does not work for me.

With `refresh`, applications still need to be restarted to see the variables
changed, but at least I can keep my session.

### Install

Dependencies:

* `github.com/AllenDang/w32`
* `gcc` from MinGW (adding `...\ruby-devkit\mingw\bin` to PATH is also fine)

~~~
go get github.com/tischda/refresh
~~~

### Usage

Simply execute:

~~~
u:\>refresh.exe
~~~

The "Environment Variables" windows should pop up for a second and disappear.

### References

* http://serverfault.com/questions/33681/how-can-i-modify-a-user-s-path-environment-variable-without-logging-out?rq=1
* http://serverfault.com/questions/8855/how-do-you-add-a-windows-environment-variable-without-rebooting
