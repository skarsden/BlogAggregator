## Contributing
Gator
-----
A CLI tool for aggreagating RSS feeds and viewing their posts
Requires latest Postgres and Go versions to be installed.

Config
------
create a ".gatorconfig.json" file in your home directory and enter the follwing:
  {
    "db_url":"postgres://<username>:@localhost:5432/gator?sslmode=disable"
  }
replace <username> with your username

Usage
-----
Create new user:(bash terminal)> gator register <username> 
Add an RSS feed:(bash terminal)> gator register <url> 
Begin Aggregator:(bash terminal)> gator agg <time interval in seconds>
View posts:(bash terminal)> gator browse <result limit>
List existing users:(bash terminal)> gator users
List feeds:(bash terminal)> gator feeds
Follow an existing feed as logged in user:(bash terminal)> gator follow <url>
Unfollow a followed feed:(bash terminal)> gator unfollow <url>
