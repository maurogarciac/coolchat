# Backend Server

The chat's backend server takes care of user authentication with a JWT and provides a websocket connection for each client that comes from the frontend server.

Recieved messages from the websocket are formatted and then stored on a PostgreSQL db.

### Setup:

1. Install golang 1.23
2. Run ```make install```

### To run:

Server:
* Run ```make start```  

Linter:
* Run ```make lint```
