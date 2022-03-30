**helping-hands**

---
**Deployment can be found at** `https://helping-hands-alpha.herokuapp.com/`

---
**Endpoints**

There are several endpoints that can be hit at `127.0.0.1:8000` :

| Package | method | full path|
|---|---|---|
|  | ping | /ping |
| /item/ | ping | /item/ping |
| /item/ | createItem | /item/createItem |
| /item/ | getItems | /item/getItems |
| /item/ | addToCart | /item/addToCart |
| /user/ | ping | /user/ping |
| /user/ | getCart | /user/getCart | 
| /user/ | getOrders | /user/getOrders |
| /user/ | createUser | /user/createUser |
| /location/ | ping | /location/ping |
| /location/ | getLocation | /location/getLocation |
| /location/ | getLocations | /location/getLocations |
| /location/ | createLocation | /location/createLocation |

The requests and responses for all endpoints can be found in `service/<package>/handler.go` file
