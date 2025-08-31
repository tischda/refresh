[![Build Status](https://github.com/tischda/refresh/actions/workflows/build.yml/badge.svg)](https://github.com/tischda/refresh/actions/workflows/build.yml)
[![Test Status](https://github.com/tischda/refresh/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/refresh/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/tischda/refresh)](https://goreportcard.com/report/github.com/tischda/refresh)

# refresh

Notify applications of Windows environment variable changes.

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
