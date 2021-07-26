# Clovers server

Server solution for the jug riddle problem

## Configuration variables
Port where server should be listen to.
```
(optional) SERVER_PORT=3000
```

## Project setup

```bash
 sudo docker build . -t clover_server
 sudo docker run --network host -it clover_server
```
