#!/usr/bin/env bash

set -e

export KO_DOCKER_REPO=us-docker.pkg.dev/suan-vmware/dump/status-check

ko build github.com/andrew-su/status-check
