#!/bin/bash
find -name '*.sh' | sed 's/.sh//g' | cut -c 2-