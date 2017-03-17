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

# postd.yaml configuration file
Copy the example configuration file from ./services/post, and modify what you need to.

# Start the Auth, Post, and API gateway services
``` bash
$ ./ramify-auth&
$ ./ramify-post&
$ ./ramify-api -auth $AUTH_HOST:3000 -post $POST_HOST:3001
```

The default ports for the auth and post service are 3000 and 3001. You can change them by passing the port number you wish for the service to run on as an argument when starting the service.

Example:
```bash
$ ./ramify-auth 5000&
```
To start the auth service on port 5000.


# Add a post with the ramify client
``` bash
$  ramify add "I'm a test post title" "I'm the body of the post."
```
