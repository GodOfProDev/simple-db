# Simple DB
A Simple go application that demonstrates how to use a postgres db

# API Documentation

### User JSON Data
* **id**: The unique identifier of the task.
* **name**: The username.
* **createdAt** The date where the user was created
* **updatedAt** The date where the user was updated
* #### GET ``/users``

```
Returns all the users
```
* #### GET ``/users/:id``
```
Returns the associated user with that id
```
* #### POST ``/users``
```
Adds a user to the db

Params:

name: The username
```
* #### PUT ``/users/:id``
```
Updates the associated user

Params:

name: The username
```
* #### DELETE ``/users/:id``
```
Deletes the associated user
```