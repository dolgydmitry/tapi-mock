# Description
## This app is implement part of TAPI v2.4.0
## Release 1.0.0 implment part of connective-service. 
## Check controllers/interfaces.go to understand current implementstion

## Implemented routes:
	POST /tapi-common:context/tapi-connectivity:connectivity-context/
	GET /tapi-common:context/tapi-connectivity:connectivity-context/
	GET /tapi-common:context/tapi-connectivity:connectivity-context/:uuid
	PATCH /tapi-common:context/tapi-connectivity:connectivity-context/:uuid


# AAA
## app use JWT to autorisied user
## inital key located in: utils/keys - !!!! only for testing use !!!!
## for Prod put your keys 



