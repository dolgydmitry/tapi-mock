# Description
# ------------------------------------------------------------------------------------
## This app is implement part of TAPI v2.4.0 as MOCK for developer purpose
## More detail see this: https://github.com/OpenNetworkingFoundation/TAPI
## Release 1.0.0 implment part of connective-service. 
## Check controllers/interfaces.go to understand current implementation

# Implemented routes:
# ------------------------------------------------------------------------------------
    POST /login/
	POST /tapi-common:context/tapi-connectivity:connectivity-context/
	GET /tapi-common:context/tapi-connectivity:connectivity-context/
	GET /tapi-common:context/tapi-connectivity:connectivity-context/:uuid
	PATCH /tapi-common:context/tapi-connectivity:connectivity-context/:uuid


# AAA
# ------------------------------------------------------------------------------------
## app use JWT to autorisied user
## inital key located in: utils/keys - !!!! only for testing use !!!!
## for Prod put your keys 

## Login request:
### POST {{HOST}}/login/
### Payload:
    {
        "username": "test",
        "password": "test"
    }
## Login response:
### Status code: 200
### token: "eyJpZCI6IjU0NTE2NTM2LWI3MWYtNDQyNy...."
### token expired after 180s

### Get start
1. clone repo
2. make build-docker
3. make docker-deploy

### See smaple of docker compose in docker-compose.yml
### See postman collection on tapi mock.postman_collection.json






