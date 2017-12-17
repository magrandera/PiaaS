# Basic Guide

## General Info

Currently supported languages:
- Python3
- NodeJs

Projects have to contain a Procfile that contain the following:
```
web: %command that starts the server%
```

Examples of projects:
- [Python](https://github.com/heroku/python-sample)
- [NodeJs](https://github.com/heroku/node-js-sample)

## Prerequisites
- Docker has to be installed

## Installing PiaaS

**PiaaS has to be installed on the Raspberry Pi & not your development machine**

You can either download one of the precompiled binaries or use go install.
If you download the precompiled binaries you have to add its location to the $PATH environment variable to be able to execute it anywhere.

## General Process

To deploy a new application, execute the following on the Pi:
```shell
PiaaS app add %name%
```
This command will then setup the folders and print out the location of the git repository. Then to deploy you application you have to add the git repository as a remote on your development machine.
```shell
git remote add PiaaS user@ip:~/PiaaS-Data/Applications/%name%/repo
```
After executing the above you can just push to the git repository and the the application will be automatically deployed. The port will be printed

## Redeploying without new push

You have to be logged in on the Raspberry Pi or server.
To redeploy:
```shell
PiaaS app deploy %name%
```

## Starting and Stopping Applications

You have to be logged in on the Raspberry Pi or server.
To stop an application:
```shell
PiaaS app stop %name%
```
To start an application:
```shell
PiaaS app start %name%
```

## Removing an application
You have to be logged in on the Raspberry Pi or server.
To remove an application:
```shell
PiaaS app remove %name%
```