# Vehicleparkingapp

[Go][go] [KAFKA][kafkalink] [KAFKA CONFLUENT][kafkaconfluent] [MongoDb][mongodblink]
  
  
<img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTaehFWyMxNLhP8_spbkEFZ5Ivh-BjHMjPe8jVQ8RmZLMNEAWrsiXrvrDOq308WErjAG44&usqp=CAU" width="100">

  This App is build to serve the requirement of the parking lot. It keeps track of the vehicle, their parking time and     the billing amount.


  ## Participants
    This application is developed in GO. It uses a distributed event streaming platform called Apache KAFKA
    and DB as MongoDB. Application has 5 different services. This can run both natively and as container service.
    It has Confluent Apache Kafka, and Mongodb container aloing with docker-compose to make it container ready.
    
  ## Running Vehicleparkingapp in container
      Prerequisite
       1. Install Docker(19.03.6) and DockerCompose(1.27.4)  // you can try the lower verision
       2. Download the source code from https://github.com/abhiranjan2429/vehicleparkingapp
       3. Open a terminal at ../vehicleparkingapp
       4. Run the command 
       5. ./build.sh // this will create the image and run docker-compose
       6. sudo docker run -d --rm --network container:broker parkinglot ./servicestarter //Running serives
      This will start all the services.
      At localhost:9021 will run Confluent Control Center, follow the link to create a topic.
        https://docs.confluent.io/platform/current/quickstart/ce-docker-quickstart.html
      Create two topic in the Name of "exit_topic" and "entry_topic" // this is configurable from the configfile.
      ** find the details in Config  section of this README.md file.
      At localhost:9011 the ENTRYSERVICE will run where you can make rest call with application/json as described
      in API section of this README.md file.
      At localhost:9012 the EXITSERVICE will run where you can make rest call with application/json as described
      in API section of this README.md file.
      At localhost:9013 the BILLINGANDTIMESERVICE will run where you can make rest call with application/json as described
      in API section of this README.md file.
      
  ## Using Vehicleparkingapp in container
    Use postman of anyother tool to make rest call to ENTRYSERVICE EXITSERVICE BILLINGANDTIMESERVICE
    on their port to simulate the parkinglot system. ENTRYSERVICE EXITSERVICE will be regular
    Entry and Exit service to vehicle as to come and go.
    BILLINGANDTIMESERVICE has the api to display the amount and parkingtime
    Calling ENTRYSERVICE_api provide the Vehicle number and submit the call and the 
    responce will generate a ParkingID which will be used for later services.
    At EXITSERVICE_api call provide the ParkingID and VehicleNumber
    AT BILLINGANDTIMESERVICE goto localhost:9013/prktime for parking time enquiry by submitting ParkingID
    and goto localhost:9013/billing for Amount Payable enquiry  by submitting ParkingID
    
    Add "plot" as db into Mongodb and create the collections as vehicleentry,vehicleexit and billing.
    Auth. is not implemented yet.
 
 ## API
    ENTRYSERVICE at localhost:9011/entry will accept api call with get method
    and body as {
                        "vehiclenumber":"XXXXX"
                }
     and in the response we get the ParkingID
     EXITSERVICE at localhost:9012/exit will accept api call with get method
     and body as {
                       "VehicleNumber": "XXXXX",
                      "ParkingID": "XXXXX"
                  }
    this will submit the query of vehicle exiting
    
    BILLINGANDTIMESERVICE has 2 Apis to provide /billing and /prktime
    BILLINGANDTIMESERVICE at localhost:9013/billing will accept api call with get method 
    and body as {
                   "parkingid": "XXXXX"

                }
    this will give the billing amount.
    BILLINGANDTIMESERVICE at localhost:9013/prktime will accept api call with get method 
    and body as {
                   "parkingid": "XXXXX"

                }
    this will give the parking time amount.

  ## Config
    Before the running the application goto pkg/pconstant/constant.go. Here you can change the Configuration to
    serve you environment.
      // excuse the way the config file is implemented in comming time will be moved to ENV variable and .yaml file.
   

  ## Known Issue
    ENTRYSERVICE and EXITSERVICE in response, the time property with empty value this will be removed.
    Better implementaion of config file.
    No log implemented
   

  ## Architecture
    The architecture diagram will be uploaded to /Docs folder
   


[Parkinglot]: https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTaehFWyMxNLhP8_spbkEFZ5Ivh-BjHMjPe8jVQ8RmZLMNEAWrsiXrvrDOq308WErjAG44&usqp=CAU
[kafkalink]: https://kafka.apache.org
[kafkaconfluent]:https://docs.confluent.io/platform/current/quickstart/ce-docker-quickstart.html
[go]:https://golang.org
[mongodblink]:https://www.mongodb.com/2
[docker]: https://docs.docker.com/get-docker/
