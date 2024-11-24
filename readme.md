# CoolChat

CoolChat is the coolest spot to chat!  
Remember to put on some sunglasses before you connect and start yapping! (background is really white)  

![alt text](https://github.com/maurogarciac/coolchat/blob/main/frontend/static/images/coolchat-preview.png?raw=true)

# How to run:

1. Make sure you have docker 27.1 installed! (or compatible version)

2. Copy the contents of .env.docker.example into a file named .env.docker
    - Optionally, you can replace SERVER_IP and SERVER_HOSTNAME with custom values. 

3. Now, in the root directory of the project, run:
```bash
docker compose up -d
```
4. Alternatively, you could run it without deamon mode, ('-d') and view the logs!

This will create all images, spin up all containers and they should connect to each-other automatically (unless some of their pre-configured ports are in use by the host machine)

The preconfigured ports are:
    - Frontend : 8000
    - Backend  : 1337

* If you want to run a single container, do:
``` bash
docker compose up -d {container_name}
```

# How to chat!

You can use any of these user credentials to log-in:
- user:'bob', password:'root'
- user:'alice', password:'root'

Now you should be able to chat in two different browsers!  

Also works in different devices in your local network, Phones and tablets included  

(btw there is no limit to how many times a user can be logged in, do with that information what you will)

Have fun!


# How to remove chat data (just in case):

Docker stores the container data inside a volume, to remove all volumes related to this project, just do run following command:

```bash
docker compose down --rmi all --volumes
```

If you re-run the compose up command, it should start the application with a brand new chat history.