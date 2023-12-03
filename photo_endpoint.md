# User Endpoint
These all methods are restricted by JWT Oauth.

In order to consume those endpoints make sure have a JWT Token and add the token value in the Header with key Authorization

### `GET` : ``v1/photos/``
#### Response
```json
    Success 200 :
    {
        {
            "success": true,
            "message": "photos successfully taken",
            "data": []
        }
    }
```


### `POST` : ``v1/photos/`` 
#### Request Body
```json
    Form-data
    Success :
    {
        "user_id":"@uuid",
        "caption": "Last trying",
        "image" :"@file"
        
    }

    Fail : 
    {
        "user_id":"@uuid",
        "caption": "Last trying",
    }
```
#### Response
```json
    Success 200 :
    {
        "success": true,
        "message": "photo successfully created",
        "data": {
            "id": "@uuid",
            "title": "@randomstring.filext",
            "caption": "Last trying",
            "image_url": "@host:port/path/title",
            "user_id": "@uuid",
            "profile_image": false,
            "created_at": "@timestamp",
            "updated_at": "@timestamp"
        }
    }

    Fail 400 :
    {
        "success": false,
        "message": "the field image is required"
    }
```


### `PUT` : ``v1/photos/:id`` 
#### Request Body
```json
    Form-data
    Success :
    {
        "user_id":"@uuid",
        "caption": "Update",
        "image" :"@file"
        
    }

    Fail : 
    {
        "user_id":"@uuid",
        "caption": "Update",
        "image" :"@nonImageFile"
        
    }
```
#### Response
```json
    Success 200 :
    {
        "success": true,
        "message": "photo successfully updated",
        "data": {
            "id": "@uuid",
            "title": "@randomstring.filext",
            "caption": "Update",
            "image_url": "@host:port/path/title",
            "user_id": "@uuid",
            "profile_image": false,
            "created_at": "@timestamp",
            "updated_at": "@timestamp"
        }
    }

    Fail 400 :
    {
        "success": false,
        "message": "image file must jpg, jpeg or png"
    }
```


### `DELETE` : ``v1/photos/:id`` 
#### Response
```json
    Success 200 : 
    {
        "success": true,
        "message": "photo successfully deleted",
        "data": {
            "id": "@uuid",
            "title": "@randomstring.filext",
            "caption": "Update",
            "image_url": "@host:port/path/title",
            "user_id": "@uuid",
            "profile_image": false,
            "created_at": "@timestamp",
            "updated_at": "@timestamp"
        }
    }

    Fail 404 :
    {
        "success": false,
        "message": "photo does not exist"
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