# refresh [![Build status](https://ci.appveyor.com/api/projects/status/ok7detq1hwnbd5cc?svg=true)](https://ci.appveyor.com/project/tischda/refresh)

Windows utility written in [Go](https://www.golang.org) to refresh
SYSTEM environment variables.

I tried https://support.microsoft.com/en-us/kb/104011 and restarted Explorer,
it does not work for me.

With `refresh`, applications still need to be restarted to see the variables
changed, but at least I can keep my session without logging out or rebooting.

### Install

Dependencies:

* `gcc` from MinGW (adding `...\ruby-devkit\mingw\bin` to PATH is also fine)

~~~
go get github.com/tischda/refresh
~~~

### Usage

~~~
Usage of refresh:
  -title="Environment Variables": localized version of title
  -version=false: print version and exit
~~~

Example (FR):

~~~
C:\>refresh.exe -title "Variables d'environnement"
~~~

The "Environment Variables" window should pop up for a second and disappear.

### References

* [How can I modify a user’s PATH environment variable without logging out?](http://serverfault.com/questions/33681/how-can-i-modify-a-user-s-path-environment-variable-without-logging-out?rq=1)
* [How do you add a Windows environment variable without rebooting?](http://serverfault.com/questions/8855/how-do-you-add-a-windows-environment-variable-without-rebooting)
