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

## Endpoints 
* **URL**

  `/riddles/jug?x=<x>&y=<y>&z=<z>`

* **Method:**

  `GET` 
  
*  **URL Params**
  
   **Required:**
   `x`:  X jug capacity
   `y`:  Y jug capacity
   `z`:  Z goal

* **Success Response:**

  * **Code:** 200 
    **Content:** steps to solve the riddle: `[ {"x": ${x}, "y": ${y}}, {"x": ${x}, "y": ${y}}, .... ]`
