# Coding challenge - Data API

## Introduction

### Problem Statement

At Swisscom we have a chatbot where we process data from our customers in real time.
A background job pushes the data (text messages) provided by the customers via HTTP to a customer data management API (that we'll call Data API).  

Data API is used by data scientists to further improve the chatbot.

In order to be compilant with the data regulation laws, our customers must give an
explicit consent to process the data - Data API also manages that.

### Task

Given the above mentioned problem statement, your goal would be to create
this Data API service that fullfills the requirements and API specification 
mentioned in _Requirements for Data API_.

## Requirements for DataAPI

Create an application with an API layer with the following endpoints:

### POST `/data/:customerId/:dialogId`
With the payload
```json
{
    "text": "the text from the customer",
    "language": "EN"
}
```

This is the endpoint used by the background job to push each customer input during their dialogue with our chatbot.


### POST `/consents/:dialogId`
With the payload  
`true` or `false`  

This endpoint is called AT THE END of the dialogue when the customer is asked if they gives consent for us to store and use their data for further improving the chatbot.  
  
If false it should delete the customer's data

### GET `/data/(?language=:language|customerId=:customerId)`

This endpoint is used by data scientists to retrieve data to improve the chatbot, it should return all the datapoints:

- that match the query params (if any)
- for which we have consent for
- and sorted by most recent data first
- implement pagination for the returned data

## Things to take into consideration

- For data storage please use a relational database (E.g. PostgreSQL, MySQL/MariaDB)
- Please make sure data is processed asyncronously

## Providing the solution

- Store your solution on a VCS that is either publicly accessible or the internal company VCS 
- Make sure you add commits as you would generally do while working on a real project.
- Add a reasonable amount of automated tests to the application
- Include documentation on how to run/deploy the application
- Include a document explaining the decisions you took during the implementation
- Think about and explain in the documentation how to improve the system for running at scale (e.g: > 1M of datapoints)
- If you use a database or any external service make sure to include a `docker-compose.yml` (or similar) file, so we can run the full stack on our own

## Expected outcomes

- Code fulfills the above mentioned requirements
- Documentation on how to run / deploy the application is provided
- Commits are meaningful
- Non-obvious decisions are documented
- Code is tested
- Code is [clean](https://www.geeksforgeeks.org/characteristics-of-a-clean-code/), and uses [design patterns](https://www.geeksforgeeks.org/software-design-patterns/)

