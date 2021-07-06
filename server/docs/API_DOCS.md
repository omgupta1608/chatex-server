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
				"data":               	id uid of registered user,
				"message":            	Relevant message from the server,
				"error":              	Relevant error from the server, if any (can be NULL),
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
				"data":               unique jwt token for user auth,
				"message":            Relevant message from the server,
				"error":              Relevant error from the server, if any (can be NULL),
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
				"data":               unique jwt token for user auth,
				"message":            Relevant message from the server,
				"error":              Relevant error from the server, if any (can be NULL),
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
