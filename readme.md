## Task PBI BTPNS Rakamin.id Bayu Pamungkas
Hello everyone, this is a task project base experience, which the task is to build an api using ``Go-Languange``. This project still not implements Clean Code and Clean Architecture yet, because it is my first project with ``Go-language`` it spend a few days to finish it despite the project just a small project.
With the project base experience or Virtual Internsip Experience Rakamin.id, it makes me to always learn something new, and it makes me know about Go fundamental too.

### Feature 
- Register
- Login
- Update user
- Remove user
- Upload profile image
- Update profile image
- Delete profile image
 
#### In progress feature
I will add the : 
- Email verification (SMTP) : After user registration user account still restricted to change user data, until the account is activated. User will receive Verification Code via email
- Change profile image : User can upload image more than 1, and user can use which image do they want to be a profile image

### Testing
Before running the Register test, make sure the users table is clear or does not have a record, because the id is PK and Email is unique. If test runs once then it will create an account and if test runs again it will return err cause the id and email already exist. Except if the id and email are different.
For Update and Delete user make sure user id is exist on the database

### Requirements
- Go-Lang 
- PostgreSQL
- Postman
- Code Editor

### Structure
```

```

### Installation
```
    go run main.go
```
- Endpoints are available in the user_endpoint.md and photo_endpoint.md

### Me
- [Github](https://www.github.com/bayek335)
- [Facebook](https://www.facebook.com/bayu.p.7146)
- [Linkedin](https://www.linkedin.com/in/bayu-pamungkas-b85399221)

