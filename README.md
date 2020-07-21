# ShopMe
 
This is an online shopping website, there are three parts: two frontend package(user and admin) and backend package.

To access it online, please see my report. Otherwise, you can try to deploy it but it may be time-costing and memory-costing to you. Since I've deployed it online, I didn't use docker for your quick deployment, there may be docker file but it will not work together. (I'm trying to build it but may be finished after the deadline of the homework).

## Deployment

### Frontend & Frontend-admin

Make sure you've installed npm and yarn.

To install and run:
``` sh
yarn
yarn serve
```

### Go

Make sure you've installed go.

To run:

``` sh
go build main.go
./main
```

## Optional Deployment

Since I've built files for deployment, you can run my built files in `dist` files.
But go project is built for windows x64, if you are using mac or else, this will never work.

### Frontend & Frontend-admin

Make sure you've installed `serve` in nodejs.

In related folder, run:
``` sh
serve -s <folder-name>
```

### Go

Double click exe file to run.