package main

import (
	"test/db"
	"test/log"
	"test/tracing"
	"test/web"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func main() {

	tp, tpErr := tracing.JaegerTraceProvider()
	if tpErr != nil {
		log.Log.Fatal(tpErr)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	Client := db.Helper{
		DB: db.ConnectDB(),
	}

	web.WebDB{}.SetupWebServer(Client.DB)

	// If you want to create the COllection first uncomment the below line
	//Client.CreateCollection("todo", "todos")

	// collection := Client.GetCollection("todo", "todos")

	// cur, _ := collection.Find(context.Background(), bson.D{})
	// var models []db.Model

	// err := cur.All(context.Background(), &models)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, model := range models {
	// 	jsonData, err := json.Marshal(model)
	// 	if err != nil {
	// 		// handle error
	// 	}
	// 	fmt.Println(string(jsonData))
	// }

}
