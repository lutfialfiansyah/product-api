# PRODUCT API

## Run
 Before run please create the database in postgresql with database name :
```bash
product_api
```
 After create database run migration :
```bash
make run_migrate
```
copy env.template and rename it to .env
after that adjust your database environment
```bash
make run_local
```
or
```bash
make run_docker
```

## Documentation
- Collection POSTMAN :
```bash
https://www.getpostman.com/collections/80a6c8bef7328e5a5c38
```
## Architecture
- DDD (Domain-Driven Design) :
```bash
- Dapat mengurangi terjadi cyclic dependency.
- Memisahkan service sesuai kebutuhan bisnis tersebut. 
- Masing - masin service tersebut terdapat sub-service contoh untuk mengambil data dll. 

```
