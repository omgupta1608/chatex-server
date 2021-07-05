# API Docs

## User APIs
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