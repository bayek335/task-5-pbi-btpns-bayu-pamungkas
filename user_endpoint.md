# User Endpoint
For user update and delete are restricted by JWT Oauth.

In order to consume those endpoints make sure have a JWT Token and add the token value in the Header with key Authorization

### `POST` : ``v1/users/register`` 
#### Request Body
```json
    Success : 
    {
        "username" : "Bayu Pamungkas",
        "email":"bayu@gmail.com",
        "password": "123123"
    }

    Fail : 
    {
       "username" : "Bayu Pamungkas",
        "email":"bayugmail.com",
        "password": "123123"
    }
```
#### Response
```json
    Success 200 :
    {
        "success" : true,
        "message" : "user succcessfully created",
        "data" : {
            "id" : "@uuid",
            "username" : "Bayu Pamungkas",
            "email":"bayu@gmail.com",
            "is_active":"false"
        }
    }

    Fail 400 :
    {
        "success" : false,
        "message" : "field email must be type of email"
        
    }
```


### `POST` : ``v1/users/login`` 
#### Request Body
```json
    Success :
    {
        "email":"bayu@gmail.com",
        "password": "123123"
    }

    Fail : 
    {
        "email":"bayu@gmail.com",
        "password": "12312"
    }
```
#### Response
```json
    Success 200 :
    {
        "success" : true,
        "message" : "user succcessfully loged in",
        "data" : {
            "token" : "@token"
        }
    }

    Fail 400 :
    {
        "success" : false,
        "message" : "wrong email or password"
        
    }
```


### `PUT` : ``v1/users/:id`` 
#### Request Body
```json
    Success :
    {
        "username" : "Bayu Pamungkas Update",
        "email":"bayu@gmail.com",
        "password": "123123"
    }

    Fail : 
    {
       "username" : "Bayu Pamungkas Update",
        "email":"bayu@gmail.com",
        "password": "123123"
    }
```
#### Response
```json
    Success 200 :
    {
        "success" : true,
        "message" : "user succcessfully updated",
        "data" : {
            "id" : "@uuid",
            "username" : "Bayu Pamungkas",
            "email":"bayu@gmail.com",
            "is_active":"false"
        }
    }

    Fail 404 :
    {
        "success" : false,
        "message" : "user does not exist"
        
    }
```


### `DELETE` : ``v1/users/:1`` 
#### Response
```json
    Success 200 : 
    {
        "success" : true,
        "message" : "User succcessfully deleted",
        "data" : {
            "id" : "@uuid",
            "username" : "Bayu Pamungkas",
            "email":"bayu@gmail.com",
            "is_active":"false"
        }
    }

    Fail 500 :
    {
        "success" : false,
        "message" : "Internal server error"
        
    }
```
### Unauthorized
### `@METHOD`
#### Response
```json
    Fail 401 :
    {
        "success" : false,
        "message" : "unauthorize, login first"
        
    }
```