# Users and their favorite Assets

## Description

This is a platform of our users that have access to a list of assets. Our users to have a peronal list of favourites, meaning assets that favourite or “star” so that they have them in their frontpage dashboard for quick access. An asset can be one the following
* Chart (that has a small title, axes titles and data)
* Insight (a small piece of text that provides some insight into a topic, e.g. "40% of millenials spend more than 3hours on social media daily")
* Audience (which is a series of characteristics, for that exercise lets focus on gender (Male, Female), birth country, age groups, hours spent daily on social media, number of purchases last month)
e.g. Males from 24-35 that spent more than 3 hours on social media daily.

We have a web server which has some endpoints to receive a user id and return a list of all the user’s favourites. Also we have endpoints that can add an asset to favourites, remove it, or edit its description. Assets obviously can share some common attributes (like their description) but they also have completely different structure and data.

## Design

An asset will be either a chart, or an insight, or an audience. Its chart has a unique Id, as well as a description
Three Databases have been used to hold the state of the program:
* A map of assetID to an Asset type
* A map of a userID to an Asset list

The createdd packages are 
* controllers for routing the endpoints and return the final responses,
* favourites to add favourite (stared) assets of each user, 
* assets to load already created assets, 
* utils to edit update the two DBs with the new description
* test packages
* integration test package

In the beginning Assets of all types and Users are loaded to the in-memory databases.
Then the application runs.

Interfaces have been used to decouple the actual code Databases from the custom Databases created for tests.

# Rest endpoints and how to use
Execute the command
```
go run main.go
```
|Method|Endpoint|Request Body|JSON Response|
|-------|:-----------------------------------------:|-----------------------------------:|-----------------------------------:|
|GET  |   localhost:8080/getuserfavourites/<user id>|    -    |user data and a list of the favourite assets or error|
|POST |   localhost:8080/addassettofav/<user id>|    {"favourites": [ <asset id>, ... ]}        |user data and list of assets types and ids |
|PATCH|   localhost:8080/editasset/<asset id>|  {"description":"new description"}  | asset type, id, and the new description|

All Ids have to be a number and an existing Id, else an error json response will be returned.
