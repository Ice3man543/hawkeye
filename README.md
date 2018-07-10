# Hawkeye
[![License](https://img.shields.io/badge/license-MIT-_red.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/Ice3man543/hawkeye)](https://goreportcard.com/report/github.com/Ice3man543/hawkeye) 
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/Ice3man543/hawkeye/issues)

HawkEye is a simple tool to crawl the filesystem or a directory looking for interesting stuff like SSH Keys, Log Files, Sqlite Database, password files, etc. Hawkeye uses a fast filesystem crawler to look through files recursively and then sends them for analysis in real time and presents the data in both json format and simple console output. The tool is built with a modular approach making it easy to use and easily extensible. 

It can be used during pentests as a privilege escalation tool to look through the filesystem finding configuration files or ssh keys sometimes left by the sys-admins. 

[![Hawkeye](http://i.imgur.com/C4prGfK.png)](https://asciinema.org/a/D1sINGdcAhJKlpzaRyexrxO1Y)]
 # Features
 
 - Simple and modular code base making it easy to contribute.
 - Fast And Powerful Directory crawling module doing real-time analysis
 - Easily extensible and vast scanner (Thanks to Gitrob)
 - Outputs in various formats 

# Installation Instructions

The installation is easy. Git clone the repo and run go build.

```bash
go get github.com/Ice3man543/hawkeye
```

## Upgrading
If you wish to upgrade the package you can use:
```bash
go get -u github.com/Ice3man543/hawkeye
```

# Usage

Hawkeye needs a directory to begin with. A directory can be supplied with `-d` flag. For example - 
```bash
./hawkeye -d <directory>
```

To run it against my home directory, i can pass /home/ice3man as the argument. 
```bash
./hawkeye -d /home/ice3man

 ice3man@TheDaemon  ~/tmp  ./hawkeye -d /home/ice3man  

 _  _                _    ___           
| || | __ _ __ __ __| |__| __|_  _  ___ 
| __ |/ _  |\ V  V /| / /| _|| || |/ -_)
|_||_|\__,_| \_/\_/ |_\_\|___|\_, |\___|
                              |__/     
	    Analysis v1.0 - by @Ice3man

[13:31:59] HawkEye : An advance filesystem analysis tool
[13:31:59] Written By : @Ice3man
[13:31:59] Github : https://github.com/Ice3man543


[Log file] /home/ice3man/.tplmap/tplmap.log
[Log file] /home/ice3man/burpsuite-master/hs_err_pid3028.log
[Log file] /home/ice3man/.log/jack/jackdbus.log
[Shell command history file] /home/ice3man/oldvps/root/.bash_history
[Shell configuration file] /home/ice3man/oldvps/root/.bashrc

```

You can use `-v` flag to show verbose output. You can also get json output using `-o` flag.
```javascript
[
    {
        "path": "/home/ice3man/oldvps/root/.bash_history",
        "description": "Shell command history file",
        "comment": ""
    },
    {
        "path": "/home/ice3man/oldvps/root/.profile",
        "description": "Shell profile configuration file",
        "comment": "Shell configuration files can contain passwords, API keys, hostnames and other goodies"
    },
    {
        "path": "/home/ice3man/oldvps/root/.bashrc",
        "description": "Shell configuration file",
        "comment": "Shell configuration files can contain passwords, API keys, hostnames and other goodies"
    },
]
```

You can list the signatures present in the tool's database by using `-l` option.
```bash
[-] Signatures present in Database:
	-> CryptoFiles
	-> PasswordFiles
	-> ConfigurationFiles
	-> DatabaseFiles
	-> MiscFiles
```

You can specify the signatures to be used by the tool by passing the `--sig` flag. It takes a comma-separated list of signatures to be used. You can also specify exclusion of certain signatures using `--exclude-sig` flag.
```bash
ice3man@TheDaemon  ~/tmp  ./hawkeye -d /home/ice3man -sig cryptofiles
ice3man@TheDaemon  ~/tmp  ./hawkeye -d /home/ice3man -exclude-sig miscfiles
```


# License

HawkEye is made with 🖤 by [Ice3man](https://github.com/Ice3man543).

By me a coffee if you appreciate my work.

[![ko-fi](https://www.ko-fi.com/img/donate_sm.png)](https://ko-fi.com/M4M7FAVC)

See the **License** file for more details.

# Thanks

HawkEye uses signatures from the awesome [Gitrob](https://github.com/michenriksen/gitrob) project by Michenriksen. Thanks to him :)
