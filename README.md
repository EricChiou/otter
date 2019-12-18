# otter
Setting config at config.ini.  

## Make config file
Making config.ini with below content:  
```
# Server
SERVER_PORT=7000
SSL_CERT_FILE_PATH=
SSL_KEY_FILE_PATH=

# MySQL
MYSQL_ADDR=127.0.0.1
MYSQL_PORT=3306
MYSQL_USERNAME=your user name
MYSQL_PASSWORD=your user password
MYSQL_DBNAME=your db name

# JWT
JWT_KEY=your jwt key
# JWT expire time, set 1 for one day, set 2 for two days, ...
JWT_EXPIRE=1

# RSA, key file path
RSA_PUBLIC_KEY=
RSA_PRIVATE_KEY=

# Environment
ENV=dev
```
