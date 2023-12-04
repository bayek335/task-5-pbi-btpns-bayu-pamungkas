## Task PBI BTPNS Rakamin.id Bayu Pamungkas
Hello everyone, this is a task Project Base Internship, which the task is to build an api using ``Go-Languange``. This project still not implement a Clean Code or a Clean Architecture yet, ya it is my first project with ``Go-language`` and I spend a few days to finish it, despite the project just have a small case.
With the Project Base Internship Rakamin.id, it makes me curious to learn something new, and it encourages me know about Go fundamental too.

### Access Endpoint
Endpoints are available in the user_endpoint.md and photo_endpoint.md, which contain the request and response from the endpoint.

### Feature 
- Register
- Login
- Update user
- Remove user
- Retrieve profile image
- Upload profile image
- Update profile image
- Delete profile image
 
#### In progress feature
I will add the : 
- Email verification (SMTP) : After user registration user account still restricted to change user data, until the account is activated. User will receive Verification Code via email
- Change profile image : User can upload image more than 1, and user can use which image do they want to be a profile image


### Testing
Before running the Register test, make sure the users table is clear or does not have a record, because the User ID is Primary Key, Email is unique and I set fixed value for the ID and email  so it will conflict if running User Testing more than once before clear the data in users table in database.

Like i was text it was my first project, so i just implement ``unit testing`` for the user feature, even it only two or three files test I take a lot of time to reach the appropriate results. And the test code is too bad and does not implement appropriate Unit Testing I thought.

### Logger
Ya for the logger, I have searching a lot for how to implement logging be a part of the code and when it is using, but I don't get it enough. So I only implement the logger in User feature 


### Tools
- Go-Lang 
- PostgreSQL
- Postman
- Code Editor

### Me
- [Github](https://www.github.com/bayek335)
- [Facebook](https://www.facebook.com/bayu.p.7146)
- [Linkedin](https://www.linkedin.com/in/bayu-pamungkas-b85399221)

