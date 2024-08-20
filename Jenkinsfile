pipeline {
    agent any
    environment {
        DOCKER_IMAGE = 'golang:1.23.0'
        POSTGRES_IMAGE = 'postgres:13'
        DB_USER = 'postgres'
        DB_PASSWORD = 'postgres'
        DB_NAME = 'books_db'
        DB_HOST = 'postgres-db'
        DB_PORT = '5432'
    }
    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/vijaymehrotra/Books-API.git'
            }
        }
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'go mod download && go build -o main'
            }
        }
        stage('Run PostgreSQL') {
            steps {
                echo 'Initilizing the DB..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}