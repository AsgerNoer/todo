# Todo-service
Small API Service that functions as a todo list.

## Set-up
Create docker image with:
`docker build . -t todo-service`

Run the image with:
`docker run -p 3000:3000 todo-service`

After this you can connect on localhost:3000. Request examples below illustrates this

# Requests
Two endpoints will be exposed. One for the entire todo list, second one for management of the individual tasks.
## Manage todo list
You can see or clear the entire todolist in one go.
### View entire list
```
curl --request GET 'localhost:3000/' \
    --header 'Content-Type: application/json'
```
### Clear entire list
```
curl --request DELETE 'localhost:3000/' \
    --header 'Content-Type: application/json'
```
 
## Manage individual tasks
Below is an overview of the different requests you can make towards the endpoint, that will do CRUD operations on tasks. When creating or updating a task please send the data in json. When doing operations on a task, the UUID of that task has to be in the URL paramters.

### Create Item
```
curl --request POST 'localhost:3000/items' \
    --header 'Content-Type: application/json' \
    --data-raw '{"ItemText": "NewTask"}'
```
### Read Item
```
curl --request GET 'localhost:3000/items?ID=UUID'
```
### Update Item
```
curl --request PUT 'localhost:3000/items&ID=UUID \
    --header 'Content-Type: application/json' \
    --data-raw '{"ItemText": "Updated task"}'
```
### Delete Item
```
    curl --request DELETE 'localhost:3000/items?ID=UUID'
```