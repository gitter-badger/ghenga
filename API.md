# Introduction

This document provides a high-level overview of the REST API. Ghenga manages
roughly the following entities:

 * People
 * Companies
 * Activities
 * Events
 * Users

# Endpoints

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

## Authentication

### POST /login

Log into ghenga with the given user name and password. Returned is an
authentication token which is valid for the given period of time. The body of the POST request must be structured as follows:

```
{
  "username": "foo",
  "password": "s3kr1t"
}
```
