# Reach

## Data Orchestration Tool

Reach is a Data Orchestration Tool that automates processes related to managing data, such as bringing / sending + receiving data together from multiple sources, and preparing it for data analysis. It includes tasks like provisioning resources, consuming data, send + receive (based on consumed) data and monitoring.

## Prerequisites

We need MariaDB for the Reach API:

```shell
sudo docker pull mariadb
sudo docker run --name sch-mariadb -p 3306:3306 -e MYSQL_ROOT_PASSWORD=some_pass -d mariadb:latest
# MYSQL_ROOT_PASSWORD needs to be updated in .env and Dockerfile
```

## Docker

```shell
 docker build -t reach .
 docker run -p 1989:1989 reach
 docker save reach > reach.tar
```

## Endpoints

- Swagger - [http://localhost:1989/api/](http://localhost:1989/api/)

## Stay in touch

- Authors - [Stefan](https://github.com/stefanciprian)
- Email - contact@schdigital.co.uk
- Website - [schdigital](https://schdigital.co.uk)

## License

Reach is an MIT-licensed open source project [MIT licensed](LICENSE).
