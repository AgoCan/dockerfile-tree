#!/bin/bash

# Remove all static lib with corresponding dynamic lib
base_dir=/usr/local/cuda/lib64/
for f in `find ${base_dir} -maxdepth 1 -type f -exec basename {} \;`; do
    if [[ $f =~ (.*)\.so\.(.*)\.(.*)\.(.*)$ ]]; then
        so_name=${BASH_REMATCH[1]}
        static_name=${base_dir}${so_name}_static.a
        if [[ -f ${static_name} ]]; then
            echo "deleting ${static_name}"; 
            rm ${static_name};
        fi
    fi
done