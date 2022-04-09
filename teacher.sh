#!/bin/bash
grep "CLUE" crimescene | sed '0~1 s/$/\n/g'
echo $MAIN_SUSPECT