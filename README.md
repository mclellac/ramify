# Ramify is a microservice test project

# Build
```bash 
$ git clone https://github.com/mclellac/ramify
$ cd ramify
$ make deps && make proto && make
$ make install
```

# Start MariaDB/MySQL server and create the DB
``` SQL
mysql> CREATE DATABASE posts;
```

# postd configuration file
Copy the example configuration file from ./services/post, and modify what you need to.

# Start the postd server
``` bash 
$ ramify-postd
```

# Add a post with the ramify client
``` bash
$  ramify add "I'm a test post title" "I'm the body of the post."
```
