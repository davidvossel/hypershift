#!/bin/bash

export IMG=quay.io/dvossel/hypershift:latest
make docker-build
make docker-push
