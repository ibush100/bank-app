# bank-app
Bank App is a banking clone project. It is a CRUD app that features users, authentication, and transactions. 

## Modules
### API 
- This module hold the routers for the various endpoints. The app runs on port 3000 by default. 

### Database 
- This module holds the logic for connecting and interacting with the database. The module was put in place so that the modules that call the functions in Database don’t have to worry about the database implementation details. 

### Helpers 
- The Helpers module contains functions for handling errors and reading requests. 

### Interfaces 
- The Interfaces module holds the interfaces and models used by the DB and the request handlers

### Transactions  
- The Transaction module methods related to making transactions 

 ### Users  
- The Users module holds methods related to the creation, updating, and authentication of users
