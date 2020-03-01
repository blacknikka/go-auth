USER=$DB_USER_NAME
PASSWORD=$DB_USER_PASSWORD
HOST=$DB_HOST
DB=$MYSQL_DATABASE

migrate.linux-amd64 -path ./migrate -database 'mysql://mysql:secret@tcp(mysql:3306)/db' up
