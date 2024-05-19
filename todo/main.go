// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Todo struct {
// 	ID    string `json:"id"`
// 	Title string `json:"title"`
// }

// var todosCollection *mongo.Collection

// func main() {
// 	// Set up MongoDB client
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, err := mongo.NewClient(clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(ctx)

// 	// Accessing a collection
// 	todosCollection = client.Database("todosDB").Collection("todos")

// 	// Define the HTTP server address
// 	addr := ":4000"

// 	// Register HTTP routes
// 	http.HandleFunc("/todos", getTodos)
// 	http.HandleFunc("/todos/add", addTodos)
// 	http.HandleFunc("/todos/delete", deleteTodos)

// 	// Create a new HTTP server with the given address
// 	server := &http.Server{
// 		Addr:    addr,
// 		Handler: http.DefaultServeMux,
// 	}

// 	// Start the HTTP server and handle any errors
// 	log.Printf("Server listening on %s", addr)
// 	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 		log.Fatalf("Server error: %v", err)
// 	}
// }

// func getTodos(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Todos: ")
// 	w.Header().Set("Content-Type", "application/json")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	cursor, err := todosCollection.Find(ctx, bson.M{})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		fmt.Println("Error retrieving todos from MongoDB:", err)
// 		return
// 	}
// 	defer cursor.Close(ctx)

// 	var todos []Todo
// 	for cursor.Next(ctx) {
// 		var todo Todo
// 		if err := cursor.Decode(&todo); err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			fmt.Println("Error decoding todo from MongoDB:", err)
// 			return
// 		}
// 		todos = append(todos, todo)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		fmt.Println("Error with cursor in MongoDB:", err)
// 		return
// 	}

// 	if err := json.NewEncoder(w).Encode(todos); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		fmt.Println("Error encoding JSON:", err)
// 		return
// 	}
// }

// func addTodos(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var todo Todo
// 	_ = json.NewDecoder(r.Body).Decode(&todo)

// 	_, err := todosCollection.InsertOne(context.Background(), todo)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		fmt.Println("Error inserting todo into MongoDB:", err)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(todo)
// }

// func deleteTodos(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodDelete {
// 		http.Error(w, "Method not allowed", http.StatusNotAcceptable)
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	var todoID string
// 	if err := json.NewDecoder(r.Body).Decode(&todoID); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	_, err := todosCollection.DeleteOne(context.Background(), bson.M{"id": todoID})
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		fmt.Println("Error deleting todo from MongoDB:", err)
// 		return
// 	}

// 	fmt.Println("Successfully Deleted")
// }


