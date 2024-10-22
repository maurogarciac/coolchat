# CoolChat

CoolChat is the coolest spot to chat!
Make sure to put on some sunglasses before you connect and start yapping! (background is really white)

# How to run:

1. Make sure you have docker 27.1 installed!

2. Copy the contents of .env.example into a file named .env

3. Copy the contents of .env.docker.example into a file named .env.docker

4. You will have to create a docker network, run:
``` docker network create coolchat-network ```

5. In the root directory of the project, run:
```docker compose up -d```

This will create all images and spin up all containers and they should connect automatically (unless some of their preconfigured ports are in use by the host machine)

The preconfigured ports are:
    - Frontend : 8000
    - Backend  : 1337

* If you want to run a single container, do:
```docker compose up -d {container_name}```

# How to chat!

You can use these user credentials to log-in:
- user:'bob', password:'root'
- user:'alice', password:'root'

Now you should be able to chat in two different browsers!

(There is no limit to how many times a user can be logged in, do with that information what you will)

Have fun!