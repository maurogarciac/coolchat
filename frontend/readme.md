# Frontend Server

The frontend server uses HTMX + Templ, to manage templates and user interaction with only a little bit of js.

Authorizes user via JWT data and connects to the backend via a websocket connection.

### Setup:

1. Install golang 1.23
2. Install templ, add it to path
3. Run ```make install``` to get all the packages
4. Install docker-buildx

### To run:

Local Server:
* Run ```make start```  

### Build for linux server:

Build binary and create image on docker:
* Run ```make binux```

Create and run new container:
* Run ```docker compose up -d```
