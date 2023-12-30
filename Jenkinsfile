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
                    sh 'cd node.js_service && npm install'
                }
            }
        }

        stage('Build Container Images') {
            steps {
                script {
                    // Authenticate Docker with AWS ECR && Build and tag Docker images
                    sh """
                    ./var/lib/jenkins/script/login.sh
                    aws ecr get-login-password --region ${AWS_DEFAULT_REGION} | docker login --username AWS --password-stdin ${ECR_REGISTRY}                    
                    cd go_service && docker build -t ${ECR_REGISTRY}:go-service:latest .
                    cd nodejs_service && docker build -t ${ECR_REGISTRY}:nodejs-service:latest .
                    """
                    
                }
            }
        }

        stage('Push to ECR') {
            steps {
                script {
                    // Push Docker images to ECR
                    sh """
                    ./var/lib/jenkins/script/login.sh
                    docker push ${ECR_REGISTRY}:go-service:latest
                    docker push ${ECR_REGISTRY}:nodejs-service:latest
                    """
                }
            }
        }

        stage('Deploy to Kubernetes Cluster') {
            steps {
                script {
                    // Apply Kubernetes manifests to EKS 
                    sh """
                       ./var/lib/jenkins/script/login.sh
                       aws eks --region ap-southeast-1 update-kubeconfig --name reynaldiekoz-eks
                       kubectl apply -f kubernetes/go-service.yaml
                       kubectl apply -f kubernetes/nodejs-service.yaml
                    """
                }
            }
        }
    }
}
