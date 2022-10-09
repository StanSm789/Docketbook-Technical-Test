This project is a simple REST API written in Go. It fetches data from Docketbook’s API and transforms the full payload into a smaller payload. 
In detail, by using two authentication headers (db-shared-key-id and db-shared-key) the program fetches data from both Dockets List URL and Single Docket URL. Then, it transforms the received data before returning it via both /dockets/ /dockets/:id enpoints. 

Steps to run the program:
1) Download the project
2) Add values to db-shared-key-id and db-shared-key keys in .env file
3) Using Terminal, go to the project folder and then type "go run main.go"
4) Open another Terminal window and type "curl localhost:8080/dockets" to fetch the data about all dockets or "curl localhost:8080/dockets/32603b07-d46f-4935-a070-993be281e970" to fetch the data about an individual docket using its id.
