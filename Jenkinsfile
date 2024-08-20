pipeline {
    agent any
    stages {
        stage('Checkout') {
            steps {
                git url: "https://github.com/vijaymehrotra/Books-API.git", branch: 'main'
            }
        }
        stage('Build') {
            steps {
                echo 'Building..'
                sh 'docker build -t go-books .'
            }
        }
        stage('Pushing to docker hub') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'dockerHub', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD')]) {
                    sh 'docker login -u $USERNAME -p $PASSWORD'
                    sh 'docker tag go-books $USERNAME/go-books'
                    sh 'docker push $USERNAME/go-books'
                }
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}