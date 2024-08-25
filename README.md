## Book Management Application
![Screenshot from 2024-08-25 17-37-19](https://github.com/user-attachments/assets/5aa577ec-4f0c-4822-a95d-cc7a61ac9b62)

### Overview

This repository contains a web application built with Golang for managing a book library. Users can create, read, update, and delete books.

## Hosted Link : https://books-api-8tv3.onrender.com/

### Technologies Used

    Backend: Golang
    Database: (Specify the database you're using, e.g., PostgreSQL, MySQL)
    Frontend: HTML, CSS, JavaScript
    Getting Started

### Clone the repository:

    git clone https://github.com/vijaymehrotra/Books-API.git

### Install dependencies:

    cd Books-API
    npm install // For frontend dependencies
    go mod tidy // For backend dependencies
### Set up database: 

Create a database instance and configure the connection details in your project's environment variables or configuration file.

### Run the application:

    go run main.go
    
### API Enndpoints

    GET /books: Retrieves a list of all books.
    POST /books: Creates a new book.
    GET /books/:id: Retrieves a specific book by ID.
    PUT /books/:id: Updates a book by ID.
    DELETE /books/:id: Deletes a book by ID.

