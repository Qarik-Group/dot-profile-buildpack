#!/bin/bash

[[ -f file-created-by-profiled ]] && { echo "ERROR: file already exists"; exit 1; }
touch file-created-by-profiled