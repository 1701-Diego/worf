# Worf - It is a good day to destroy Tasks and LRPs

Worf destroys Tasks and LRPs

## To Install

```
go get github.com/1701-diego/worf
```

## To Use

First export the `RECEPTOR` address

For Ketchup:
```
export RECEPTOR=http://username:password@receptor.ketchup.cf-app.com
```

For Diego-Edge:
```
export RECEPTOR=http://receptor.192.168.11.11.xip.io
```

Then, to destroy a task:

```
worf destroy-task TASK_GUID TASK_GUID ...
```

To destroy an LRP:

```
worf destroy-lrp PROCESS_GUID PROCESS_GUID ...
```