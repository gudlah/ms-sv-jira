# ms-sv-jira

###### version: v0.5

###### Last Update: 2024-08-13

### Programming Language

- [go1.20 windows/amd64]

### Framework

- [GIN]

#### The diagram:

![golang clean architecture](clean-architecture.png)

The explanation about this project's structure  can read from this medium's post : https://medium.com/@imantumorang/golang-clean-archithecture-efd6d7c43047

### How to Deploy on Localhost

```
go run .
```

### How to Test on Localhost or Development Server

```
go test bitbucket.org/ms-sv-jira/test/
```

### Environment Variables (Development)
Merupakan suatu variabel yang ditetapkan diluar program melalui fungsionalitas yang dibangun ke dalam OS atau Microservices. Tim DevOps perlu mengetahui Nama Env. Variables, Value Env. Variables, Description dari Env. Variables, dan Locationnya.

| Name                              | Description                                   |
|-----------------------------------|-----------------------------------------------|
| GIN_MODE                          | Gin Mode                                      |
| DB_SCHEMA                         | Database Schema                               |
| DB_DRIVER                         | Database Driver                               |
| DB_HOST                           | Database Host                                 |
| DB_PORT                           | Database Port                                 |
| DB_USER                           | Database User                                 |
| DB_PASSWORD                       | Database Password                             |
| DB_SSL_MODE                       | Database SSL Mode                             |
| SERVICE_HOST                      | Host Service                                  |
| SERVICE_PORT                      | Port Service                                  |
| ELASTIC_APM_SERVER_URL            | Elastic APM URL                               |
| ALLOW_ORIGIN                      | Allow Origin                                  |
| BRIGATE_URL                       | Brigate URL                                   |
| BRIGATE_USERNAME                  | Brigate Username                              |
| BRIGATE_PASSWORD                  | Brigate Password                              |
| ESB_URL                           | ESB URL                                       |
| ESB_CHANNEL_ID                    | ESB Channel Id                                |
| ESB_GET_CREDIT_CARD_SERVICE_ID    | ESB Get Credit Card Service Id                |
| JWT_EXPIRED                       | JWT Expiration Time                           |