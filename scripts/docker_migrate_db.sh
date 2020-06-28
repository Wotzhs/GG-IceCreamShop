#!/bin/sh

wget https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz
wget -P /bin https://raw.githubusercontent.com/eficode/wait-for/master/wait-for

chmod +x /bin/wait-for

tar -C /bin -xvf migrate.linux-amd64.tar.gz

sh -c 'wait-for $DB_HOST:5432 -- migrate.linux-amd64 -database "$DB_URL" -path /home/migrations up'