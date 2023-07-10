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

# APP UI

![alt text](./documentation/appui.png?raw=true "App UI")

# Technologies
    * Go lang, Javascript, HTML, CSS
    * GORM (ORM: defines structs which have methods for storing and finding the entity from database)
    * Fiber: is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go
    * Gin Gonic is a high-performant web framework written in Go (Golang)
    * PostgreSQL Database (pgAdmin)
    * JSON (json objects)
    * REST API's
    * HTTP, HTTPS
    * ReactJS, react-router, axios (consume REST API's), moment (utils library to format), Jest (unit test)
    * Docker (is an open platform for developing, shipping, and running applications, enables you to 
    separate your applications from your infrastructure)
    

# Architecture

  MVC Architecture: Webapp (PWA)

   ## Package Diagram
  ![alt text](./documentation/PackageDiagram.png?raw=true "Package Diagram")
  
  Backend:

    * migrations: Create the table in the database when I run go
    * controller: Definition of API endpoints and access to them, in addition to their configuration.
    * models: Mapping of the database tables like alerts.
    * repository: It is the connection to the database, they also call this repositories (they call the structs that connect to the database), use pattern DAO.
    * routes: Setup API's routes (get, post, delete, put) with fiber and use repository.
    * storage: Database configuration like dbname, password, port, host, user, among others... (use Postgres driver).
    * bootstrap: setup environment variables, database and fiber configurations

  Frontend:

     * alerts: app views to management alerts like: show, create, delete and update alerts

# Database

  Database Name: go_crud

  ## Entity Relationship Diagram
  ![alt text](./documentation/EntityRelationshipDiagramCore.PNG?raw=true "Entity Relationship Diagram Core")

    * alerts: Table save the app alerts

# Deploy
## Docker
Build image with the command:
docker build -t jps-alert-crud .
(The . It is because we are located where the Dockerfile is and jps-alert-crud 
is the name that we are going to give the container)

Check the image with the command:
docker images

And finally, execute the image with the command:
docker run --name jps-alert-crud -p 8080:8080 jps-alert-crud

Tip: with upx library reduce the size of the image

# Utils
## Postman Requests
Import request for test apis, path: jps-alert-crud/documentation/Postman Requests/Alerts.postman_collection


# Run the app
## 1) run go (Backend)
Path: /jps-alert-crud/ 
Command: go run main.go

![alt text](./documentation/gorunning.png?raw=true "Go Command Running")

## 2) run react (Frontend)
Path: /jps-alert-crud/client/
Command: npm install

Path: /jps-alert-crud/client/
Command: npm start

![alt text](./documentation/reactrunning.png?raw=true "React Command Running")
