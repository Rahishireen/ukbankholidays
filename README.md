# UK Bank Holidays
Go application that interfaces with the Bank Holiday API to fetch raw data. Subsequently, the application is tasked with processing this data according to specific requirements, filtering it based on predefined criteria. Finally, the processed and filtered datasets are exposed through API endpoints.

## Table of Contents
-[Requirements](#requirements)

-[Getting Started](#getting-started)

   -[Installation](#installation)
   
   -[Running the Application](#running-the-application)
  
-[Endpoints](#endpoints)

-[Testing](#testing)

-[Logging](#logging)

-[Error Handling](#error-handling)

-[Future Enhancements](#future-enhancements)

## Requirements

To run the application, you need the following

- Go version 1.21.4
- Below third party packages
  - github.com/julienschmidt/httprouter v1.3.0
  - github.com/sirupsen/logrus

  ## Getting Started

  ### Installation

  Clone the repository to your local machine

  ### Running the application

  1. cd ukbankholidays
  2. go build 
  3. ukbankholidays.exe
     
  Note: Tested in windows machine

## Endpoints
![image](https://github.com/Rahishireen/ukbankholidays/assets/81709473/dd921985-84c2-4d34-94df-a3b38b469ac4)

Note: 
1. APIs can be triggered through postman(preferred) / web browser
   eg:
   
   -http://localhost:8080/holidays-by-year?year=2023
   
   -http://localhost:8080/holidays-england-wales
   
   -http://localhost:8080/holidays-by-year-title-date?year=2023

## Testing
go test ./...

- Added unit test cases only for key parts of the application i.e controllers and services

## Logging 
- Incorporated the logrus third-party package in Go to enrich log entries with essential details such as timestamps
- filenames, line numbers and log levels.
- The default log level is set to "info".

## Error handling

Currently handling few http status errors like
-  Invalid request
-  Internal server error

## Future Enhancements

### 1. Authentication Implementation
Integrate a robust authentiaction mechanism to control access to the application

### 2. SSL Communication
Implement SSL communication to ensure that the data transmitted between the client and server is encrypted 

### 3. Handle Application Errors
Implement logic to handle application errors, if any

  
