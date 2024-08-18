pipeline {
    agent any
    environment {
        DOCKER_IMAGE = 'golang:1.23.0'
        POSTGRES_IMAGE = 'postgres:13'
        DB_USER = 'postgres'
        DB_PASSWORD = 'postgres'
        DB_NAME = 'books_db'
    }
    stages {
        stage('Build') {
            steps {
                echo 'Building..'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}