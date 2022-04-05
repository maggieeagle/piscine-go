#!/bin/bash
find -name '*.sh' -maxdepth 1 -type f -exec basename {} \; 2> /dev/null | sed 's/.sh//g'