# Mock Api Example

[credit for this tutorial goes to](https://www.youtube.com/watch?v=t96hBT53S4U)


*the following tutorial assumes you have golang installed, also you'll need postman installed in order to test out this api_*

---

## Testing and building the api manually on your local machine. 

-  install postman [postman](https://www.getpostman.com/)
-  clone this gitrepo and browse to the *mock-api* folder
-  if you decide to edit the main.go file please rebuild with `go build && mock-api`
-  in order to run this api please use the following command `./mock-api`
-  this api listens on port *1111*

---

## Making a portable build using docker

- set the following variables *BUILDTIME_IMAGE, BUILDTIME_VERSION, DKR_IMG_NAME*
- run this `docker build --build-arg $BUILDTIME_IMAGE --build-arg $BUILDTIME_VERSION --name $DKR_IMG_NAME .`
- Start the docker image `docker run $DKR_IMAGE_NAME -p 1111:1111`
- The contaier should now be accessible at localhost:1111
---
