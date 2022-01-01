#!/bin/bash
dir=$(ls /data/sql)
# read -p "Enter your username, please: " username

# read -p "Enter your password, please: " pswd

user=root
password=12345
host=localhost

mysql -h $host -u$user -p$password --default-character-set=utf8mb4 < create_db.sql
# mysql -h $host -u$user --default-character-set=utf8mb4 < create_db.sql

echo "use mf_admin;" >all.sql
for i in $dir; do
    echo "source /data/sql/$i;" >>all.sql
done

mysql -h $host -u $user -p$password --default-character-set=utf8mb4 <all.sql
