# Repositorio: jps-alert-crud

# Mercado Libre Interview: Alerts
Para avanzar a la siguiente etapa del proceso, te comparto el challenge que debes desarrollar. Tendrás 7 días para hacerlo, si no llegas a hacerlo al 100%, no te preocupes, puedes enviarnos lo que puedas lograr hacer.
No tenemos un lenguaje específico para la realización del challenge, pero estaría bueno hacerlo en React y Go (que es nuestra stack). Utilización de docker es un diferencial.

Al final de los 7 días tienes que enviarnos el repositorio de github, si es posible con todas las informaciones del sistema y la documentación de cómo ejecutar local.

El challenge se trata de un portal web el cual recibe por api datos sobre distintos eventos.

El formato de entrada de estas alertas es así: {"type": "xxxx", "description": "ddddd", "created_at": "hh-mm-ss-dd-mm-yyyy", "country: "argentina|brasil|colombia"}

Esta información se tiene que guardar en una base de datos relacional, luego puede ser accedida tanto por api como por un portal web.

En el portal web, sin necesidad de autenticarse, un usuario debe poder buscar alertas en una barra de búsqueda, ya sea por descripción del event así como también por nombre del país.

El type nunca cambia y como deseable, el product owner solicitó que se pueda acceder a todas los eventos de un type buscando a este por nombre. Se deben tener en cuenta los aspectos relacionados a testing, usabilidad, paginación, visualización de la información.

Del lado de la api, aparte de poder consultar los eventos, es necesario también tener un endpoint que a modo de métrica/estadística te liste los 3 países con la mayor cantidad de eventos en el último mes, junto con un valor entero que especifique esa cantidad.

# Technologies
    * Go lang
    * GORM (ORM: defines structs which have methods for storing and finding the entity from database)
    * PostgreSQL Database (pgAdmin)
    * JSON (json objects)
    * REST API's
    * HTTP, HTTPS
    * Sonarqube (continuous inspection of code quality) (http://localhost:9000/projects) and install plugin sonarlint in your IDE
    (After config proyect, execute the command: clean org.jacoco:jacoco-maven-plugin:prepare-agent install)
    And each update execute: sonar:sonar -Dsonar.login=9feb9fcf40d00ce1aa1faf7529966d68e7b1000a
    * Docker (is an open platform for developing, shipping, and running applications, enables you to 
    separate your applications from your infrastructure)
    

# Architecture
   ## Package Diagram
  ![alt text](./documentation/PackageDiagram.jpg?raw=true "Package Diagram")
  
    * config: Is to make all settings of the microservice, how to configure the api with spring security so that non-validated or identified users do not access it, JWT configuration, define the number of transactions per day, condfiguration documentation and consumption of the CAD api.
    * security: Manage security within the entire microservice, in this microservice we are using spring security with JWT and for example in this package we generate and validate the authentication token
    * controller: Definition of API endpoints and access to them, in addition to their configuration.
    * dto: Data transfer object and its objective is to separate the api from the database, because if the model changes, it does not affect the api (POJO classes).
    * error: Contains the parameterized errors in the application of which we already know why they happen.
    * exception: Handle exceptions within the entire microservice depending on whether they are handled or unhandled exceptions.
    * model: Mapping of the database tables (POJO classes: classes that have get, sets, constructors... (encapsulation) ).
    * repository: It is the connection to the database, they also call this repositories (they call the classes that connect to the database), 
    use pattern DAO.
    * service: Interfaces used by controllers to perform operations such as insert, delete, query, update, perform certain functionality.
    * service.impl: Is the implementation of the services interfaces, this to isolate the definition of the implementation.

# Database
  ## Entity Relationship Diagram
  ![alt text](./documentation/EntityRelationshipDiagramCore.PNG?raw=true "Entity Relationship Diagram Core")

    * account: Table save the accounts of clients
    * movements: Table save the transactions of clients

  ![alt text](./documentation/EntityRelationshipDiagramSecurity.PNG?raw=true "Entity Relationship Diagram Security")

    * users: Table save the users of microservice
    * roles: Table save the roles of microservice
    * user_roles: Detail table save the user with role of microservice

# Deploy
## Docker
Build image with the command:
docker build -t midocker .
(The . It is because we are located where the Dockerfile is and midocker 
is the name that we are going to give the container)

Check the image with the command:
docker images

And finally, execute the image with the command:
docker run -it -p 8080:8080 midocker

# Utils
## Postman Requests
Import request for test apis, path: FundsTransfers/documentation/Postman Requests/FundsTranfers.postman_collection
