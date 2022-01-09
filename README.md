# Article REST-API
Rest API for articles data created using Go programming language with some requirements to handle:
- Thousands of articles in database
- Many user accessing API at the same time

## Routes
- Post new article --> `POST : ROOT_URL/article`
- Get articles --> `GET : ROOT_URL/article?query=[your-query-here]&author=[your-author-here]`

## Tools
- [Elasticsearch](https://www.elastic.co/?ultron=B-Stack-Trials-APJ-Exact&gambit=Stack-Core&blade=adwords-s&hulk=paid&Device=c&thor=elasticsearch&gclid=CjwKCAiA5t-OBhByEiwAhR-hm55h2dKBhzaGYLj4s9GEzdeFVFZvTUjmSfjuQVNAcpEiD_bIZg7iXBoCN3oQAvD_BwE)
- [Httprouter](https://github.com/julienschmidt/httprouter)
- [Redis](https://redis.io/)
- [Postgresql](https://www.postgresql.org/)

## Solution Overral
- To handle thousands articles in database and searching feature with query to look inside title and body   of the article, i use elasticsearch as the search engine. So i use two database, one for storing master data and elasticsearch for searching optimization
- To handle many user access at the same time, i use redis to cache the response. So if someone do GET request, it indirectly querying to database, instead go lookup to the cache. If there is some resources cached, then simply return it as the response.

High level design (more detail --> [whimsical](https://whimsical.com/CBkzTKAYHWvCWJuDyaFu1J))
![high level design](https://github.com/zipzap11/Article-RestApi/blob/master/high_level_design.PNG)

## Run
To run this you can pull some images from dockerhub
#### Postgres
```
docker pull postgres:12-alpine
```
#### Elasticsearch
```
docker pull elasticsearch:7.16.2
```
#### Redis
```
docker pull redis:latest
```
#### Database setup (if docker compose file doesn't work)
Then run the postgres image to create database
```
docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine 
```
and create database
```
docker exec -it postgres12 psql -U root
```
then, inside postgres container terminal run
```
CREATE DATABASE golang_article;
```

#### Run APP
After all images ready, run
```
docker-compose up
```


