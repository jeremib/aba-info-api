# Routing Number API Service

## API Documentation

### GET /banks
Returns a list of all the banks in json format:

```
	{
		"routing_number": "324184440",
		"bank": "IDAHO UNITED CREDIT UNION",
		"address": "PO BOX 2268",
		"city": "BOISE",
		"state": "ID",
		"zip": "83701",
		"zip4": "2268",
		"phone": "2083882138"
	},
	{
		"routing_number": "324274033",
		"bank": "ELKO FEDERAL CREDIT UNION",
		"address": "2397 MTN CITY HWY",
		"city": "EIKO",
		"state": "NV",
		"zip": "89801",
		"zip4": "0000",
		"phone": "7757384083"
	},
	...
```

### GET /banks/\<RoutingNumber\>

This returns bank information for the routing number specified

```
{
    "routing_number": "031101334",
    "bank": "SoFi Bank, N.A.",
    "address": "San Francisco, CA",
    "city": "San Francisco",
    "state": "CA",
    "zip": "",
    "phone": "(855) 936-2269",
    "message": "OK"
}
```

### GET /health

Returns a status if the microservice is up (used by the load balancer)

```
"up"
```

## Building Docker Container
Use the following commands below to build the docker container.  For this to be successful, you need to use the correct `--profile` to match your set up.

```
aws ecr get-login-password --region us-east-1 --profile egov-prod | docker login --username AWS --password-stdin 859052165483.dkr.ecr.us-east-1.amazonaws.com
docker build -t routing-number-api-prod .
docker tag routing-number-api-prod:latest 859052165483.dkr.ecr.us-east-1.amazonaws.com/routing-number-api-prod:latest
docker push 859052165483.dkr.ecr.us-east-1.amazonaws.com/routing-number-api-prod:latest
```
Update the running production service:
```
aws ecs update-service --cluster egov-micro-services-cluster-prod --service routing-number-api-service-prod --force-new-deployment --region us-east-1 --profile egov-prod
```

## Adding new routing number information

Edit the `data/banks.json` file with the new bank routing information.  Then Build and update listed above.

## Terraform
This API is deployed using the Microservices Terraform script found at https://dev.azure.com/eGovernmentSolutions/DevOps/_git/eGov%20Microservices%20Terraform
