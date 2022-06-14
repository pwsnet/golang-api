# Golang API Template

Create a golang API with the following structure:
* `/`: root path
* `/api`: API Package contains all the endpoints and middlewares
* `/endpoints`: Endpoints package contains all the endpoints
* `/application`: Application package contains all the application logic, getters
* `/infrastructure`: Infrastructure package contains all the infrastructure logic, fetchers
* `/infrastructure/builders`: Builders package contains all the builders for fetchers (Postgres, MySQL) prepared statements
* `/infrastructure/{engine}`: Engines package contains all the engines for fetchers (Postgres, MySQL)
* `/lib/{lib}`: Libs package contains all the libs for integrations and utils (Railway, AWS, etc)


## Usage

To run the main pipeline, run the following command:

```bash
make env=[development | production]
```