package main

import (
	context "context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "example.com/calculadora/proto"

	"google.golang.org/grpc"
)

var (
	// Puerto 10000 por defecto
	port = flag.Int("port", 10000, "The server port")
)

type calculadoraServer struct {
	pb.UnimplementedCalculadoraServiceServer
}

func (c *calculadoraServer) Sumar(ctx context.Context, input *pb.NumbersList) (*pb.ResultObject, error) {
	// Implementación de la operación sumar
	total := input.Numbers[0]

	for i := 1; i < len(input.Numbers); i++ {
		total += input.Numbers[i]
	}

	return &pb.ResultObject{
		Status: "OK",
		Result: total,
	}, nil
}

func (c *calculadoraServer) Restar(ctx context.Context, input *pb.NumbersList) (*pb.ResultObject, error) {
	// Implementación de la operación restar
	total := input.Numbers[0]

	for i := 1; i < len(input.Numbers); i++ {
		total -= input.Numbers[i]
	}

	return &pb.ResultObject{
		Status: "OK",
		Result: total,
	}, nil
}

func (c *calculadoraServer) Multiplicar(ctx context.Context, input *pb.NumbersList) (*pb.ResultObject, error) {
	// Implementación de la operación multiplicar
	total := input.Numbers[0]

	for i := 1; i < len(input.Numbers); i++ {
		total *= input.Numbers[i]
	}

	return &pb.ResultObject{
		Status: "OK",
		Result: total,
	}, nil
}

func (c *calculadoraServer) Dividir(ctx context.Context, input *pb.NumbersList) (*pb.ResultObject, error) {
	// Implementación de la operación dividir
	total := input.Numbers[0]

	for i := 1; i < len(input.Numbers); i++ {
		if input.Numbers[i] != 0 {
			total /= input.Numbers[i]
		} else {
			return &pb.ResultObject{
				Status: "ERROR - División entre cero",
				Result: 0,
			}, nil
		}
	}

	return &pb.ResultObject{
		Status: "OK",
		Result: total,
	}, nil
}

func main() {
	// Leemos el puerto desde una flag para poder modificarlo en caso de ser requerido
	flag.Parse()
	// Abrimos el puerto, escuchando conexiones TCP en 0.0.0.0:port
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Print("listening on port: ", *port)

	// Creamos el servidor grpc, en este caso sin opciones
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// Registramos el servicio dentro del servidor grpc
	pb.RegisterCalculadoraServiceServer(grpcServer, &calculadoraServer{})

	// Iniciamos el servidor
	grpcServer.Serve(lis)
}
