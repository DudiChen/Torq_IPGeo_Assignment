The Service requires the following Evironment Variables:

DATABASE_TYPE =“csv”
DATABASE_PATH = <path to data.csv> 
RATE_REQUESTS = 1
RATE_INTERVAL = 1s

The above values are to run the service using a local csv type SB store that contains a single dummy data entry:  
67.250.186.196,New-York,United-States

So in order to get a successfull hit, need to use a key (IP Address) of 67.250.186.196

Note: The project does not include all CRUD needed funtionality to maintain the DB store as the assignment did not require it.

Please see partial Unit-Testing available for CSV Store under csv/store_test.go

The project includes a Dockerfile and bash script ["run_ipgeo_service.sh"].
These facilitate a prompt execution by using a docker container.

Usefult command to run by Docker container:

Get Container ID:
sudo docker ps | grep ipgeo | awk '{print $1}'


To get the docker IP address:
sudo docker inspect \                                              
  -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' <CONTAINER_ID>


Send a HTTP request:
curl --location '<CONTAINER_IP_ADDRESS>:8080/v1/find-country?ip=67.250.186.196'
