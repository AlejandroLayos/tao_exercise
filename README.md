# Tao Exercise

## Proyect Architechture 
This project follows hexagonal architecture and DDD patterns. 
The files are organized in this way :
 
```sh
|-- config
|-- logs
|-- public
|-- src
|   |-- core
|   |    |-- question domain
|   |    |   |-- application
|   |    |   |-- domain
|   |    |   |-- infrastructure
|   |-- services
|   |    |-- logger
|   |    |-- server

```
- **config**: contains the classes than allows the user to configure the API. The user can choose the server port, 
the input data type or some logs options
    
- **logs**: output directory for logs

- **public**: our main class is here. In this class the application is initialized and the different 
components are configured

- **src**: 
    - **core**:  
        - **application**: this directory contains the use cases associated with the entity 
        - **domain**: this directory contains our entity and the classes associated with it 
        - **infrastructure**: this last layer implements the handlers, the loggers and the data of the entity 
    - **services**: third party classes 



## Launch the client 

```sh
make apiRest-pre
```

```sh
./application
```

## Import Postman

https://www.getpostman.com/collections/384b3516af0c95659385