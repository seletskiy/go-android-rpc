#!/bin/bash

base_url=${1:-https://github.com/seletskiy/go-android-rpc/blob/master/}

grep '@'TODO -r . -nI | awk -vbase_url=$base_url -f <(cat <<-'END'
    BEGIN {
        FS=":"
    }

    {
        if ($4) {
            $4 = substr($4,2);
        } else {
            $4 = substr($1,3);
        }

        link="[" $4 "] (" base_url substr($1,3) "#L" $2 ")";

        print "- [ ] " link;
    }
END
)
