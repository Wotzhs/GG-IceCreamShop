#!/bin/sh

wget -P /bin https://raw.githubusercontent.com/eficode/wait-for/master/wait-for

chmod +x /bin/wait-for

sh -c 'wait-for $DB_HOST:5432 -- ./home/ice_cream_import'