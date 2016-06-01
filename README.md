Hello World APP

HelloWorld.go is simple REST API Server written in Golang. It serves to only POST requests that are received to "/hi" endpoint.
For example, when a post request is sent with json, {firstName: "Deniz", lastName: "Mekik"}, the response will be a json, {body: "Hi Deniz Mekik"}

There is 3 different test cases to test the application. 
  - a post request with json proper filled fileds
  - a post request with json fields are emthy string
  - a post request without json
  
  
TO DO
More test cases, e.g. Type, name lenght checking

