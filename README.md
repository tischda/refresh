# refresh [![Build status](https://ci.appveyor.com/api/projects/status/axw7a46cbm0ro45q?svg=true)](https://ci.appveyor.com/project/tischda/refresh)

Windows utility written in [Go](https://www.golang.org) to notify applications of
environment variable changes.

`refresh` notifies applications that environment variables have changed by sending
them a `WM_SETTINGCHANGE` message.

Some applications will still need to be restarted to become aware of the variables
changes, but it's no more necessary to log out or reboot.


### Install

There are no dependencies.

~~~
go install github.com/tischda/refresh@latest
~~~

### Usage

~~~
refresh.exe
~~~
