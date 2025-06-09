#!bin/bash

# find all .sh files, remove the extension and sort in descending order
find . type f -name "*.sh" | sed 's/\ //;s/\.sh$//' | sort-r