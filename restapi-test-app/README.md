# RestAPI Application with Unit Test using Testify and Mockery

If you want to run this application locally, you can download/clone this repo entirely and keep this folder only.
By running this application, you'll install all dependencies needed to run this application and it will spawn a server running on port 9090 (the default port I set).

## How to run:
1. Using docker container
   1. Create your own .env file with keys listed in sample.env, then you can fill the values.
   2. Set the specified values of build_args in docker-compose file. Fill it with your linux user id, user group id, and username. The purpose of doing this is to change default user inside docker container.
   3. Run `docker-compose up` to fetch and build images need for this project.
   
2. The ol' go way:
   1. Make sure you have a database running.
   2. Create tweets table using sql script inside sql folder.
   3. Create and fill the .env file with the specified values (you can look on sample.env)
   4. Then run `go run main.go`. This will install the dependencies and run the server.
   
PS: There are several endpoints without tests. So I think you can apply your newly acquired knowledge here :) 

Feel free to contact me if you have any feedbacks or questions through my email [agus.richard21@gmail.com](mailto:agus.richard21@gmail.com)