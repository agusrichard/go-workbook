# Hexagonal Architecture in Go

</br>

## List of Contents:
### 1. [A Hexagonal Software Architecture Implementation using Golang and MongoDB](#content-1)


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

## References:
- https://www.linkedin.com/pulse/hexagonal-software-architecture-implementation-using-golang-ramaboli/
- https://github.com/alramaboli/hexagonal-architecture