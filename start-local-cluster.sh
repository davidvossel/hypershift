#!/bin/bash

# starts cluster, starts kubevirt
./kubevirtci up
# installs hypershift operator from source
./kubevirtci sync



# https://amd64.ocp.releases.ci.openshift.org/ is helpful for finding releases
export RELEASE_IMAGE=quay.io/openshift-release-dev/ocp-release:4.9.21-x86_64
export CONTAINER_DISK_IMAGE=quay.io/containerdisks/rhcos:4.9-pre-release
export KUBECONFIG=/home/dvossel/go/src/github.com/openshift/hypershift/cluster-up/_ci-configs/k8s-1.21/.kubeconfig
bin/hypershift create cluster kubevirt --name test-cluster --pull-secret /home/dvossel/pull-secret-file.txt --ssh-key /home/dvossel/.ssh/id_rsa.pub --node-pool-replicas 1 --containerdisk $CONTAINER_DISK_IMAGE --release-image $RELEASE_IMAGE


