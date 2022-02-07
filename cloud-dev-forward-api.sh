
#!/bin/bash

oc port-forward svc/kube-apiserver 6443:6443 -n clusters-test-cluster
