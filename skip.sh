#!bin/bash

# list files and skip every other line (starting with the first)
ls -l | awk 'NR % 2 == 0'