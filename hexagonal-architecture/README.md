# Hexagonal Architecture in Go

</br>

## List of Contents:
### 1. [A Hexagonal Software Architecture Implementation using Golang and MongoDB](#content-1)
### 2. [Hexagonal Architecture in Go](#content-2)

</br>

---

## Contents:

## [A Hexagonal Software Architecture Implementation using Golang and MongoDB](https://www.linkedin.com/pulse/hexagonal-software-architecture-implementation-using-golang-ramaboli/) <span id="content-1"></span>

### The Overview
- Hexagonal software architecture seeks to enable developers to create independently maintainable and testable code.
- Layout of hexagonal architecture:
  ![Layout of hexagonal architecture](https://media-exp1.licdn.com/dms/image/C4D12AQFWwt0Nvc7u9g/article-inline_image-shrink_1000_1488/0/1612880516976?e=1639612800&v=beta&t=S261vMIcSF3Qhvf8RwUDihIrPlKgJ1ekRy04QSkq_x8)
- At the core of the architecture is the business logic, which describes and encapsulates the business rules within a given business context.
- According to domain-driven design (DDD), the vocabulary that should be used to encode the business rules must be commonly understood between software developers and business domain experts, thus leading to a ubiquitous language that fosters a shared understanding of the business problem.
- A hexagonal software architecture--as shown in Fig 1--isolates the business logic from external actors (repositories, user interfaces, APIs, etc) and only allows communication through ports and adapters.
- Ports are the interfaces that are defined within the business logic boundaries, specifying a set of operations that can be performed on the business domain objects.
- Concrete implementations of the methods specified by the business logic ports effectively create adapters.
- Outside concerns, such as repositories and user interfaces, use the adapters to plug into the business domain services [logic].
- It is the use of ports and adapters in a hexagonal software architecture that leads to highly decoupled software components that can be tested and maintained independently.
- Consequently, external actors, i.e., frameworks and infrastructure, can be swapped out of the application environment without impacting the business logic.
- This creates great flexibility for infrastructure modernisation and optimisation.
- For instance, an organisation that wishes to replace a SQL database with a non-SQL one can do so without the need to make changes to the application logic.
- The application logic is completely independent of outside concerns.
- The application logic is completely independent of outside concerns. If anything, the external agents depend on the interfaces that the application logic defines, resulting in dependencies that point inwards to the core of the hexagonal architecture, which is where the business logic lives.

### The Build
- It's time to build! The 'construction' of the product catalogue microservice begins at the core of the hexagonal architecture, defining the domain model, interfaces and logic.
- The domain model is a simple Product struct that has three properties, and it is defined in the `product.go` source file.
- `product.go`
  ```go
  package domain

  type Product struct {
      Code  string  `json:"code" bson:"code"`
      Name  string  `json:"name" bson:"name"`
      Price float32 `json:"price" bson:"price"`
  }
  ```
- The interfaces defined for the domain logic are Repository and Service, and they are laid out in the `service.go` source file. 
- The interfaces are the ports that external actors will plug their adapters into to drive the application.
- The Repository interface enables the product catalogue service to connect to data store adapters to persist and query the product catalogue data.
- The Service interface is an abstraction of the main application logic, which is driven by external agents such as the REST APIs that we build later in the article.
- `service.go`
  ```go
  package domain

  type Service interface {
    Find(code string) (*Product, error)
    Store(product *Product) error
    Update(product *Product) error
    FindAll() ([]*Product, error)
    Delete(code string) error
  }

  type Repository interface {
    Find(code string) (*Product, error)
    Store(product *Product) error
    Update(product *Product) error
    FindAll() ([]*Product, error)
    Delete(code string) error
  }
  ```
- To build the application logic, the Product model and the interfaces are stitched together in the `logic.go` source.
- The logic presented in the source simply defines a service struct that attaches to a repository via the Repository interface and then creates a product catalogue service by implementing the Service interface.
- Note that the repository infrastructure details are not reflected in the domain logic implementation.
- Therefore, the service can flexibly connect to any data store infrastructure an organisation may so choose.
- `logic.go`
  ```go
  type service struct {
    productrepo Repository
  }

  func NewProductService(productrepo Repository) Service {

    return &service{productrepo: productrepo}
  }

  func (s *service) Find(code string) (*Product, error) {

    return s.productrepo.Find(code)

  }

  func (s *service) Store(product *Product) error {

    return s.productrepo.Store(product)

  }
  func (s *service) Update(product *Product) error {
    return s.productrepo.Update(product)
  }

  func (s *service) FindAll() ([]*Product, error) {
    return s.productrepo.FindAll()
  }

  func (s *service) Delete(code string) error {

    return s.productrepo.Delete(code)
  }
  ```
- In this article, the product catalogue service is connected to mongoDB as the data store infrastructure. 
- To accomplish this, a mongoDB adapter is created by constructing a mongo repository based on the mongoRepository struct, which implements the Repository interface.
- The mongo.go source file lays out the details of the implementation. For the sake of brevity, only the Create (Store) operation in CRUD is shown in the code snippet.
- It's interesting to note that whenever a need arises to replace the mongo repository with mysql or redis, all that's required is to create a mysql or redis adapter and then attach it to the Repository port with no changes to the application logic.
- `mongo.go`
  ```go
  package repository
  import (
    "context"
    "log"
    "time"

    "github.com/pkg/errors"
    "github.com/projects/hexagonal-architecture/domain"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
  )
  type mongoRepository struct {
    client  *mongo.Client
    db      string
    timeout time.Duration
  }
  //NewMongoRepository ...
  func NewMongoRepository(serverURL, dB string, timeout int) (domain.Repository, error) {
    mongoClient, err := newMongClient(serverURL, timeout)
    repo := &mongoRepository{
      client:  mongoClient,
      db:      dB,
      timeout: time.Duration(timeout) * time.Second,
    }
    if err != nil {
      return nil, errors.Wrap(err, "client error")
    }

    return repo, nil
  }
  func (r *mongoRepository) Store(product *domain.Product) error {
    ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
    defer cancel()
    collection := r.client.Database(r.db).Collection("items")
    _, err := collection.InsertOne(
      ctx,
      bson.M{
        "code":  product.Code,
        "name":  product.Name,
        "price": product.Price,
      },
    )
    if err != nil {
      return errors.Wrap(err, "Error writing to repository")
    }
    return nil
  }
  ```
- The product catalogue service is built and it has a mongoDB data store attached to it. It's time to create an http handler to drive the application by making REST API calls to the service.
- The handler is a simple struct consisting of a product catalogue service and implementing the http methods defined by the ProductHandler interface.
- `rest-handler.go`
  ```go
  package api

  import "net/http"

  type ProductHandler interface {
    Get(http.ResponseWriter, *http.Request)
    Post(http.ResponseWriter, *http.Request)
    Put(http.ResponseWriter, *http.Request)
    Delete(http.ResponseWriter, *http.Request)
    GetAll(http.ResponseWriter, *http.Request)
  }
  type handler struct {
    productService domain.Service
  }
  func NewHandler(productService domain.Service) ProductHandler {

    return &handler{productService: productService}

  }
  func (h *handler) Post(w http.ResponseWriter, r *http.Request) {
    p := &domain.Product{}
    err := json.NewDecoder(r.Body).Decode(p)
    if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
    }
    err = h.productService.Store(p)
    if err != nil {
      http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
      return
    }
    json.NewEncoder(w).Encode(p)
  }
  ```
- It's been quite an effort to get to the finish line, but here we are in the beloved main.go, weaving together everything we have built and anxiously getting ready to fire "go run main.go." 
- In the application's entry point (the main method), the server and database settings are loaded.
- The database settings are used to create a mongo repository, which is then attached to the product catalogue service. Next, the service is bound to an http handler.
- Finally, an http router is created to route requests to the handler. The router is based on the go-chi library. Finally, we ListenAndServe on the server port that's read from the loaded configuration file.
- `main.go`
  ```go
  package main

  import (
    "log"
    "net/http"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/projects/hexagonal-architecture/api"
    "github.com/projects/hexagonal-architecture/config"
    "github.com/projects/hexagonal-architecture/domain"
    "github.com/projects/hexagonal-architecture/repository"
  )

  func main() {

    conf, _ := config.NewConfig("./config/config.yaml")
    repo, _ := repository.NewMongoRepository(conf.Database.URL, conf.Database.DB, conf.Database.Timeout)
    service := domain.NewProductService(repo)
    handler := api.NewHandler(service)

    r := chi.NewRouter()
    r.Use(middleware.RequestID)
    r.Use(middleware.RealIP)
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    r.Get("/products/{code}", handler.Get)
    r.Post("/products", handler.Post)
    r.Delete("/products/{code}", handler.Delete)
    r.Get("/products", handler.GetAll)
    r.Put("/products", handler.Put)
    log.Fatal(http.ListenAndServe(conf.Server.Port, r))

  }
  ```

**[⬆ back to top](#list-of-contents)**

</br>

---

## [Hexagonal Architecture in Go](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3) <span id="content-2"></span>

### Hexagonal Architecture - Core
- In this architecture, everything is surrounding the core of the application. It is a technology agnostic component that contains all the business logic.
- In other words, the core shouldn’t be aware of how the application is served or where the the data is actually hold.
- The core could be viewed as a “box” (represented as a hexagon) capable of resolve all the business logic independently of the infrastructure in which the application is mounted.
- This approach allow us to test the core in isolation and give us the ability to easily change infrastructure components.

### Hexagonal Architecture - Actors
- Actors are real world things that want to interact with the core.
- These things could be humans, databases or even other applications.
- Actors can be categorized into two groups, depending on who triggers the interaction:
  - Drivers (or primary) actors, are those who trigger the communication with the core. They do so to invoke a specific service on the core. A human or a CLI (command line interface) are perfect examples of drivers actors.
  - Driven (or secondary) actors, are those who are expecting the core to be the one who trigger the communication. In this case, is the core who needs something that the actor provides, so it sends a request to the actor and invoke a specific action on it. For example, if the core needs to save data into a MySQL database, then the core trigger the communication to execute an INSERT query on the MySQL client.
- Notice that the actors and the core “speak” different languages.
- An external application sends a request over http to perform a core service call (which does not understand what http means).
- Another example is when the core (which is technology agnostic) wants to save data into a mysql database (which speaks SQL).

### Hexagonal Architecture - Ports
- In one hand, we have the ports which are interfaces that define how the communication between an actor and the core has to be done.
- Depending on the actor, the ports has different nature:
  - Ports for driver actors, define the set of actions that the core provides and expose to the outside. Each action generally correspond with a specific case of use.
  - Ports for driven actors, define the set of actions that the actor has to implement.
- Notice that the ports belongs to the core. It is important, due to the core is the one who define which interactions are needed to achieve the business logic goals.

### Hexagonal Architecture - Adapters
- In the other hand, we have the adapters that are responsible of the transformation between a request from the actor to the core, and vice versa.
- This is necessary, because as we said earlier the actors and the core “speaks” different languages.
- An adapter for a driver port, transforms a specific technology request into a call on a core service.
- An adapter for a driven port, transforms a technology agnostic request from the core into an a specific technology request on the actor.

### Hexagonal Architecture - Dependency Injection
- After the implementation is done, then it is necessary to connect, somehow, the adapters to the corresponding ports.
- This could be done when the application starts and it allow us to decide which adapter has to be connected in each port, this is what we call “Dependency injection”.
- For example, if we want to save data into a mysql database, then we just have to plug an adapter for a mysql database into the corresponding port

### Case of study: MinesWeeper API
- We are going to build an API for the popular game called Minesweeper.
- As we mention earlier, in this architecture everything is surrounding the core of the application, therefore, it is important to start by building the business logic.
- At this point, just forget for a moment where the data actually will be hold or how the application will be served. Just put all your energy implementing and testing the core.

### Case of study: MinesWeeper API - Application structure
- Directory structure:
  ```text
  ├── cmd
  ├── pkg
  └── internal
      ├── core
      │   ├── domain
      │   │   ├── game.go
      │   │   └── board.go
      │   ├── ports
      │   │   ├── repositories.go
      │   │   └── services.go
      │   └── services
      │       └── gamesrv
      │           └── service.go
      ├── handlers
      └── repositories
  ```

### Case of study: MinesWeeper API - Core
- All the core components (services, domain and ports) will be placed in the directory ./internal/core.

### Case of study: MinesWeeper API - Domain
- All the domain models will be placed in the directory ./internal/core/domain.
- It contains the go struct definition of each entity that is part of the domain problem and can be used across the application.
- Note: not every go struct is a domain model. Just the structs that are involved in the business logic.
- `domain.go`
  ```go
  // ./internal/core/domain/domain.go

  package domain

  type Game struct {
    ID            string        `json:"id"`
    Name          string        `json:"name"`
    State         string        `json:"state"`
    BoardSettings BoardSettings `json:"board_settings"`
    Board         Board         `json:"board"`
  }

  type BoardSettings struct {
    Size  uint `json:"size"`
    Bombs uint `json:"bombs"`
  }

  type Board [][]string
  ```

### Case of study: MinesWeeper API - Ports
- It contains the interfaces definition used to communicate with actors.
- `ports.go`
  ```go
  // ./internal/core/ports/ports.go

  package ports

  type GamesRepository interface {
      Get(id string) (domain.Game, error)
      Save(domain.Game) error
  }

  type GamesService interface {
      Get(id string) (domain.Game, error)
      Create(name string, size uint, bombs uint) (domain.Game, error)
      Reveal(id string, row uint, col uint) (domain.Game, error)
  }
  ```

### Case of study: MinesWeeper API - Services
- The services are our entry points to the core and each one of them implements the corresponding port.
- `service.go`
  ```go
  // ./internal/core/services/gamesrv/service.go

  package gamesrv

  type service struct {}

  func New() *service {
    return &service{}
  }

  func (srv *service) Get(id string) (domain.Game, error) {
    return domain.Game{}, nil
  }
  ```
- We know that somehow the game is saved in a storage. Any kind of storage.
  ```go
  // ./internal/core/services/gamesrv/service.go

  package gamesrv

  type service struct {
    gamesRepository ports.GamesRepository
  }

  func New(gamesRepository ports.GamesRepository) *service {
    return &service{
      gamesRepository: gamesRepository,
    }
  }

  func (srv *service) Get(id string) (domain.Game, error) {
    game, err := srv.gamesRepository.Get(id)
    if err != nil {
      return domain.Game{}, errors.New("get game from repository has failed")
    }

    return game, nil
  }
  ```
  ```go
  // ./internal/core/services/gamesrv/service.go

  package gamesrv

  type service struct {
    gamesRepository ports.GamesRepository
    uidGen          uidgen.UIDGen
  }

  func New(gamesRepository ports.GamesRepository, uidGen uidgen.UIDGen) *service {
    return &service{
      gamesRepository: gamesRepository,
      uidGen:          uidGen,
    }
  }

  func (srv *service) Get(id string) (domain.Game, error) {...}

  func (srv *service) Create(name string, size uint, bombs uint) (domain.Game, error) {
    if bombs >= size*size {
      return domain.Game{}, errors.New("the number of bombs is invalid")
    }

    game := domain.NewGame(srv.uidGen.New(), name, size, bombs)

    if err := srv.gamesRepository.Save(game); err != nil {
      return domain.Game{}, errors.New("create game into repository has failed")
    }

    return game, nil
  }
  ```

### Adapters
- Now, it’s time to implement the adapters so the application can interact with the actors.
- Having all the components decoupled from each other give us the advantage to implement and test them in isolation or we can easily parallelize the work with the help of other members of the team.

### Driver adapter
- All the driver adapters will be placed in packages inside the directory ./internal/handlers.
- `http.go`
  ```go
  // ./internal/handlers/gamehdl/http.go

  package gamehdl

  type HTTPHandler struct {
    gamesService ports.GamesService
  }

  func NewHTTPHandler(gamesService ports.GamesService) *HTTPHandler {
    return &HTTPHandler{
      gamesService: gamesService,
    }
  }

  func (hdl *HTTPHandler) Get(c *gin.Context) {
    game, err := hdl.gamesService.Get(c.Param("id"))
    if err != nil {
      c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
      return
    }

    c.JSON(200, game)
  }
  ```

### Driven adapter
- `memkvs.go`
  ```go
  // ./internal/repositories/gamesrepo/memkvs.go

  package gamesrepo

  type memkvs struct {
    kvs map[string][]byte
  }

  func NewMemKVS() *memkvs {
    return &memkvs{kvs: map[string][]byte{}}
  }

  func (repo *memkvs) Get(id string) (domain.Game, error) {
    if value, ok := repo.kvs[id]; ok {
      game := domain.Game{}
      err := json.Unmarshal(value, &game)
      if err != nil {
        return domain.Game{}, errors.New("fail to get value from kvs")
      }

      return game, nil
    }

    return domain.Game{}, errors.New("game not found in kvs")
  }
  ```

### Serve the application
- `main.go`
  ```go
  // ./cmd/httpserver/main.go

  package main

  func main() {
    gamesRepository := gamesrepo.NewMemKVS()
    gamesService := gamesrv.New(gamesRepository, uidgen.New())
    gamesHandler := gamehdl.NewHTTPHandler(gamesService)

    router := gin.New()
    router.GET("/games/:id", gamesHandler.Get)
    router.POST("/games", gamesHandler.Create)

    router.Run(":8080")
  }
  ```

### Advantages
- Separation of concerns: each component (core, adapters, ports, etc) has a well-defined purpose and there is no doubt of their responsibilities.
- Focus on the business logic: delaying the technical details allows you to focus on what matters at the end, the business logic.
- Parallelization of work: once the ports are defined, it is easy to parallelize the work across mates. Having several members of the team working in different well-defined and decouple components can reduce the development time considerably.
- Tests in isolation: each component can be tested in isolation, and more important is that the core is self-tested.
- Easily change infrastructure: it is really easy to change the infrastructure. You can move from a mysql to an elastic-search database without an impact on the business logic.
- Self-guided process: the architecture itself guides you on how the development process steps should be taken. Starts from the core, continue with the ports and adapters and finally serve the application.

### Disadvantages
- Too complex for small or short-term projects: it is important to analyze if this architecture is appropriate for the desired project. For example, if the micro-service has only one specific task it could be overkill. Or if it is short-term project sometimes is better to keep it simple.This architecture is recommended for applications with real business domain problems.
- Performance overhead: adding extra components trigger extra calls to functions, therefore, in each of them we will be adding a very small overhead. This could be a disadvantage if our service has to be extremely performant.


**[⬆ back to top](#list-of-contents)**

</br>

---

## References:
- https://www.linkedin.com/pulse/hexagonal-software-architecture-implementation-using-golang-ramaboli/
- https://github.com/alramaboli/hexagonal-architecture
- https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3
- https://github.com/matiasvarela/minesweeper-hex-arch-sample