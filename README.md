[![Build Status](https://github.com/tischda/refresh/actions/workflows/build.yml/badge.svg)](https://github.com/tischda/refresh/actions/workflows/build.yml)
[![Test Status](https://github.com/tischda/refresh/actions/workflows/test.yml/badge.svg)](https://github.com/tischda/refresh/actions/workflows/test.yml)
[![Coverage Status](https://coveralls.io/repos/tischda/refresh/badge.svg)](https://coveralls.io/r/tischda/refresh)
[![Linter Status](https://github.com/tischda/refresh/actions/workflows/linter.yml/badge.svg)](https://github.com/tischda/refresh/actions/workflows/linter.yml)
[![License](https://img.shields.io/github/license/tischda/refresh)](/LICENSE)
[![Release](https://img.shields.io/github/release/tischda/refresh.svg)](https://github.com/tischda/refresh/releases/latest)


# refresh

Notify applications of Windows environment variable changes.

`refresh` notifies applications that environment variables have changed by sending
them a `WM_SETTINGCHANGE` message.

Some applications will still need to be restarted to become aware of the variables
changes, but it's no more necessary to log out or reboot.


## Install

~~~
go install github.com/tischda/refresh@latest
~~~

## Usage

~~~
Usage: refresh [OPTIONS]

Refresh environment variables from the Windows registry.

OPTIONS:

  -?, --help
          display this help message
  -v, --version
          print version and exit
~~~
