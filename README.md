# csi-dev

> a simple csi implement for kubernetes

## feature

* create and delete volume
* push and delete volume
* lesser code
* nfs protocol

## you can learn what something

all-in-one kubernetes csi plugin implementation flow

## how to deploy it

### env

1. Good nfs server that include as well as correct `/etc/export` file,firewall config

### reproduction

1. `make docker`
2. `make deploy`
3. `make pvc`
4. `make pod`
5. `cp index.html` in your container volume path
6. browser the website