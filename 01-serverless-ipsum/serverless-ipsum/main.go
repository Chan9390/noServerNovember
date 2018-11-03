package main

import (
	"bytes"
	"math/rand"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var dictionary = []string{"auto-scaling",
	"zero-maintenance",
	"pay-per-execution",
	"serverdeath",
	"function",
	"event",
	"handler",
	"cloud",
	"NoOps",
	"Lambda",
	"microservices",
	"monitoring"}

var htmlContentStart = `
<!doctype html>
<html>
<head>
  <meta charset="utf-8">

  <title></title>

  <script type="text/javascript">
  </script>

  <style type="text/css">
    div {
      margin: auto;
      margin-top: 50px;
      width: 50%;
      font-family: "Courier New", monospace;
    }
    h1 {
      text-align: center;
    }
  </style>

</head>

<body>
  <div>
    <h1> Serverless Ipsum </h1>
	`

var htmlContentEnd = `
  </div>
</body>

</html>`

func getParagraph() string {
	var buffer bytes.Buffer
	words := rand.Intn(99) // get a random number for words
	if words < 25 {
		words = 25
	}

	for i := 0; i < words; i++ {
		r := rand.Intn(len(dictionary)) // get a random word
		buffer.WriteString(dictionary[r] + " ")
	}

	return buffer.String()
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       htmlContentStart + getParagraph() + htmlContentEnd,
		Headers: map[string]string{
			"Content-Type": "text/html",
			"Project":      "noServerNovember",
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
