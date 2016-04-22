# Introduction

This document provides a high-level overview of the REST API. Ghenga manages
roughly the following entities:

 * People
 * Companies
 * Activities
 * Events
 * Users

# Endpoints

The API is reachable at the path `/api`

## Authentication

All requests to the API (except the next one) must be authenticated.

### GET /login/token

Log into ghenga with the given user name and password in the HTTP basic auth.
Returns an authentication token which is valid for the given period of time.
The body of response looks as follows:

```
{
  "token": "8890bb0467cfe0bde7ec8554b6b01e4174ee6217ed540fc811ef4bfac80c082e",
  "valid_for": 7200,
}
```

The token needs to be submitted in the HTTP header `X-Auth-Token` for all
requests to the API.

## People

This endpoint manages all entries for people in the database. People can be
communicated with and are assigned to a company.

### GET /person

Returns a list of all persons.

### POST /person

Create a new person. In the body, a JSON document describing the new person
must be submitted.

### GET /person/:id:

Returns the data for the specified person.

### PUT /person/:id:

Updates the entry for the person with the specified ID. The body must contain a
JSON document with the changed attributes. Attributes that are not specified
here will not be modified.

### DELETE /person/:id:

Removes the person with the given ID from the database.

## Users

This endpoint manages ghenga users.

### GET /user

Returns a list of all users known to ghenga.

### POST /user

Create a new user. In the body, a JSON document describing the user
must be submitted.

### GET /user/:id:

Returns the data for the specified user.

### PUT /user/:id:

Updates the entry for the user with the specified ID. The body must contain a
JSON document with the changed attributes. Attributes that are not specified
here will not be modified.

# Errors

When an error occurs, the server returns an appropriate HTTP response code and
an optional JSON document in the body.

For example when ghenga is unable to reach the database server, the HTTP status
code 500 (internal server error) and the following document is returned:

```
{
  "message": "Unable to connect to database",
  "code": "DATABASE_DOWN"
}
```
