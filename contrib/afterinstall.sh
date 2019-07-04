#!/bin/bash
export USER="nobody"
export GROUP="nogroup"
id -u $USER &>/dev/null || useradd $USER
id -g $GROUP &>/dev/null || groupadd $GROUP

# Disable systemctl for running in docker
# systemctl daemon-reload
