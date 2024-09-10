## Code structure

The code is split into 3 parts based on the functionality.
The endpoints in this project do not interact with the database, they evaluate the received message and if the are no issues with the request, then an external function will complete the request asynchronously.

The tests for each part are located in the same folder as the code.

### Main

Main is the start of the application, connection the functions to the endpoints stated. 

### Consents

Consents contains the endpoint regarding user consent on saving the dialog data. This only contains the endpoint, because not offering consent will result in the deletion of the data from the database, therefore the deletion happens in Data.

### Data

Here the endpoints regarding the addition of dialog entries and fetching of data are located. Helper functions that interact with the database are also kept here. Data structures relevant to the database along with the database variable are present in [data_structures.go](data_structures.go).

## Code scalability

The interaction with the database and the endpoints is currently slow, so if this would be deployed and used with large ammounts of data, the request with take too long to be completed. A centralized database will help with this, since we will be able to store the data on disk and access it much faster.
