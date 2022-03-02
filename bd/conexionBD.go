package bd

import  (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
//MongoCN es el objeto de conexion a la base de datos
var MongoCN=ConectarBD()
var clienteOptions=options.Client().ApplyURI("mongodb+srv://admin:MongoDB@cluster0.euzag.mongodb.net/myFirstDatabase?retryWrites=true&w=majority\n")

//ConectarBD es la funcion que permite la conexion a la base de datos
func ConectarBD() *mongo.Client{
	client, err:=mongo.Connect(context.TODO(),clienteOptions)
if err != nil{
	log.Fatal(err)
	return client
}
	err= client.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa con la BD")
	return client
}
//ChequeoConection es el ping a la base de datos
func ChequeoConection() int {
	err:=MongoCN.Ping(context.TODO(), nil)
	if err != nil{
		return 0
	}
	return 1
}