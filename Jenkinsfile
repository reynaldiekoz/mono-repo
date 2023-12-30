pipeline {
    agent any

    environment {
        AWS_DEFAULT_REGION = 'ap-southeast-1' // Set your AWS region
        ECR_REGISTRY = '003866745935.dkr.ecr.ap-southeast-1.amazonaws.com/reynaldiekoz' // Set your ECR registry URL
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build and Test') {
            steps {
                script {
                    // Add build and test steps if needed
                    sh 'cd go_service && /usr/local/go/bin/go build'
                    sh 'cd node_service && npm install'
                }
            }
        }

        stage('Build Container Images') {
            steps {
                script {
                    // Authenticate Docker with AWS ECR
                    sh "aws ecr get-login-password --region ${AWS_DEFAULT_REGION} | docker login --username AWS --password-stdin ${ECR_REGISTRY}"

                    // Build and tag Docker images
                    sh 'cd go_service && docker build -t ${ECR_REGISTRY}:go-service:v1 .'
                    sh 'cd nodejs_service && docker build -t ${ECR_REGISTRY}:nodejs-service:v1 .'
                }
            }
        }

        stage('Push to ECR') {
            steps {
                script {
                    // Push Docker images to ECR
                    sh 'docker push ${ECR_REGISTRY}:go-service:v1'
                    sh 'docker push ${ECR_REGISTRY}:nodejs-service:v1'
                }
            }
        }

        stage('Deploy to Kubernetes Cluster') {
            steps {
                script {
                    // Apply Kubernetes manifests to EKS
                    sh 'kubectl apply -f kubernetes/go-service.yaml'
                    sh 'kubectl apply -f kubernetes/nodejs-service.yaml'
                }
            }
        }
    }
}
