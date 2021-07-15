# API Docs

## User APIs
 - Request data validation info can be found at `pkg/types/reqData`
 - Register a user
	 - ` /api/v/register `
	 - Public API
	 - POST
	 - POST Body
	  ```
		{
				"name":							Username. Required,
				"about":						About/Bio for the user. Optional,
				"email":						User's email id. Required,
				"password":					User's password. Required,
		}
		```
	 - Response
	  ```
		{
				"user":               	user data which was just created,
				"message":            	Relevant message from the server,
				"error":              	Relevant error from the server, if any (can be NULL),
				"error_fields":					Req data fields which failed validation
				"verification-route": 	Points to the route to verify a user,
		}
		```
 - Verify a user
	 - ` /api/v/register/verify `
	 - Public API
	 - POST
	 - POST Body
	  ```
		{
				"uid":								User's id uid. Required,
				"verification_code":	Verification code. Required,
		}
		```
	 - Response
	  ```
		{
				"jwt":              	unique jwt token for user auth,
				"user":								logged in user's data
				"message":            Relevant message from the server,
				"error":              Relevant error from the server, if any (can be NULL),
				"error_fields":				Req data fields which failed validation
		}
		```
 - Login a user
	 - ` /api/v/register/verify `
	 - Public API
	 - POST
	 - POST Body
	  ```
		{
				"email":								User's email id. Required,
				"password":							User's password. Required,
		}
		```
	 - Response
	  ```
		{
				"jwt":              	unique jwt token for user auth,
				"user":								logged in user's data
				"message":            Relevant message from the server,
				"error":              Relevant error from the server, if any (can be NULL),
				"error_fields":				Req data fields which failed validation
		}
		```
 - Get User By Id
   - ` /api/v/user/:uid `
   - Private API
   - GET
   - Response
    ```
    {
        data:       Data that is requested (details of the user with id uid)
        message:    Relevant message from the server
        error:      Relevant error from the server, if any (can be NULL)
    } 
    ```
 - Edit User Profile
   - ` /api/v/user/edit/:uid `
   - Private API
   - POST
   - POST Body
    ```
    {
        "name":       user's name
        "about":    User's about
        "profile_pic":      user's profile pic URL
    } 
    ```
   - Response
    ```
    {
        data:       Data that is requested (details of the user with id uid)
        message:    Relevant message from the server
        error:      Relevant error from the server, if any (can be NULL)
    } 
    ```
 - Change User Password
   - ` /api/v/user/change-password/:uid `
   - Private API
   - POST
   - POST Body
    ```
    {
        "old_password": User's current password
		"new_password": The new password
    } 
    ```
   - Response
    ```
    {
        data:       Data that is requested (details of the user with id uid)
        message:    Relevant message from the server
        error:      Relevant error from the server, if any (can be NULL)
    } 
    ```
 - Delete User By Id
   - ` /api/v/user/delete-account/:uid `
   - Private API
   - DELETE
   - Response
    ```
    {
        data:       Data that is requested (details of the user with id uid)
        message:    Relevant message from the server
        error:      Relevant error from the server, if any (can be NULL)
    } 
    ```
