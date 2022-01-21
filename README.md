# Users and their favorite Assets

## Design

This is a platform of our users that have access to a list of assets. Our users to have a peronal list of favourites, meaning assets that favourite or “star” so that they have them in their frontpage dashboard for quick access. An asset can be one the following
* Chart (that has a small title, axes titles and data)
* Insight (a small piece of text that provides some insight into a topic, e.g. "40% of millenials spend more than 3hours on social media daily")
* Audience (which is a series of characteristics, for that exercise lets focus on gender (Male, Female), birth country, age groups, hours spent daily on social media, number of purchases last month)
e.g. Males from 24-35 that spent more than 3 hours on social media daily.

We have a web server which has some endpoints to receive a user id and return a list of all the user’s favourites. Also we have endpoints that can add an asset to favourites, remove it, or edit its description. Assets obviously can share some common attributes (like their description) but they also have completely different structure and data.
