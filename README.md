# newggw 
[![Build Status](https://travis-ci.org/Flukas88/newggw.svg?branch=master)](https://travis-ci.org/Flukas88/newggw)

gRPC client/server for my old gogetwheater project


## How to build it
    $ make
    $ make build_client
    $ make build_server


## How to run it
    $ export GGW=<> ./cmd/server/server
    $ ./cmd/client/client -city Milan -degrees C
