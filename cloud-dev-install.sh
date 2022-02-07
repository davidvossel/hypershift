#!/bin/bash

export IMG=quay.io/dvossel/hypershift:latest
./bin/hypershift install --hypershift-image=${IMG}
