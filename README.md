Twitter contest winner API.

1. go build .
2. go run main.go
3. Send GET request in the form of localhost:8080/winner?id=i_kanganaranaut&tweetCount=15&rtCount=15
4. Receive JSON output as winner
5. Send GET request in the form of localhost:8080/winners to see list of previous winners and their counts.