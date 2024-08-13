# Arithmetic calculator

Web platform to provide a simple calculator functionality (addition, subtraction, multiplication, division, square root, and a random string generation) 
where each functionality will have a separate cost per request.

User’s will have a starting credit/balance. Each request will be deducted from the user’s balance. If the user’s balance isn’t enough to cover the 
request cost, the request shall be denied.

## Endpoints

Import the [`Calc.postman_collection.json`](/Calc.postman_collection.json) file into **Postman** to view the entire information and payload of these endpoints.

* Login -> `POST /v1/login`
* Create operation -> `POST /v1/operations`
* Perform operation -> `POST /v1/operations/{operation_id}/perform`
* Fetch records -> `GET /v1/records`
* Delete record -> `DELETE /v1/records/{record_id}`

### Perform Operation

To perform an operation you first need to create it, then send a `perform operation` request with the expected payload; the payload varies depending on the operation type (name).

* Addition -> `{ "num1": "1", "num2": "4" }`
* Subtraction -> `{ "num1": "5", "num2": "3" }`
* Multiplication -> `{ "num1": "5", "num2": "2" }`
* Division -> `{ "num1": "10", "num2": "5" }`
* Square root -> `{ "num1": "1" }`
* Random string -> `No payload`


## Developer guide

In order to start developing on this project, you'd need to:

1. Install `docker` and `docker-compose`.
2. Optionally configure the `docker-compose.yml` and the `.env` file to your needs.
4. Run `docker compose up -d` to start the services.
5. Start making changes or querying the API.
