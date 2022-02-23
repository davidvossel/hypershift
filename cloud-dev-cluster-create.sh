
#!/bin/bash

# https://amd64.ocp.releases.ci.openshift.org/ is helpful for finding releases
export RELEASE_IMAGE=quay.io/openshift-release-dev/ocp-release:4.9.21-x86_64
export CONTAINER_DISK_IMAGE=quay.io/containerdisks/rhcos:4.9-pre-release
export BASE_DOMAIN="gcp.devcluster.openshift.com"
bin/hypershift create cluster kubevirt --name test-cluster --pull-secret /home/dvossel/pull-secret-file.txt --ssh-key /home/dvossel/.ssh/id_rsa.pub --node-pool-replicas 5 --containerdisk $CONTAINER_DISK_IMAGE --release-image $RELEASE_IMAGE --base-domain $BASE_DOMAIN


