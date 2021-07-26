# Clovers ui

Ui solution for the jug riddle problem

## Configuration variables
To configure the server that the application should connect to.
```
(optional) SERVER_URL=localhost:3000
```

## Project setup

### Docker setup:
```bash
 sudo docker build . -t clover_ui
 sudo docker run --network host -it clover_ui
```

### Local setup: 
```
npm install
```

#### Compiles and hot-reloads for development
```
npm run serve
```

#### Compiles and minifies for production
```
npm run build
```

#### Lints and fixes files
```
npm run lint
```

