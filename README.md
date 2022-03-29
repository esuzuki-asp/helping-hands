**helping-hands**

---

**Project Setup**

This project uses 
- go 1.17
- postgres
- docker-compose


---
To create the container for our database you will need the `docker-compose` terminal command. Start the container with the following command:

```
docker-compose up -d
```





Rebuild project (optional): This project contains a build already, to rebuild the project you can use the following command:
```
go build helping-hands
```

Run the server:
```
./helping-hands runServer
```

---
**Endpoints**

There are several endpoints that can be hit at `127.0.0.1:8000` :

| Package | method | full path|
|---|---|---|
| /items/ | ping | /items/ping |
| /user/ | ping | /user/ping |
| /user/ | getCart | /user/getCart | 
| /user/ | getItems | /user/getOrders |