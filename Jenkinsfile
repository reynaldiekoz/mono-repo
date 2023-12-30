pipeline {
    agent any

    environment {
        AWS_DEFAULT_REGION = 'ap-southeast-1' // Set your AWS region
        ECR_REGISTRY = '003866745935.dkr.ecr.ap-southeast-1.amazonaws.com/reynaldiekoz' // Set your ECR registry URL
        APP_NAME1 = 'go-service'
        APP_NAME2 = 'nodejs-service'
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
                    //  Build and tag Docker images
                    sh "cd go_service && docker build -t ${APP_NAME1}:latest ."                    
                    sh "cd node.js_service && docker build -t ${APP_NAME2}:latest ."
                    
                    
                }
            }
        }

        stage('Push to ECR') {
            steps {
                script {
                    // Push Docker images to ECR
                    sh """

                    #ECR Login
                    /var/lib/jenkins/script/login.sh
                    aws ecr get-login-password --region ${AWS_DEFAULT_REGION} | docker login --username AWS --password-stdin ${ECR_REGISTRY}                    
                    #Docker Push
                    docker tag $APP_NAME1:latest $ECR_REGISTRY:${APP_NAME1}
                    docker tag $APP_NAME2:latest $ECR_REGISTRY:${APP_NAME2}
                    docker push ${ECR_REGISTRY}:${APP_NAME1}
                    docker push ${ECR_REGISTRY}:${APP_NAME2}
                    """
                }
            }
        }

        stage('Deploy to Kubernetes Cluster') {
            steps {
                script {
                    // Apply Kubernetes manifests to EKS 
                    sh """
                       #kubectl config
                       /var/lib/jenkins/script/login.sh
                       aws eks --region ap-southeast-1 update-kubeconfig --name reynaldiekoz-eks
                       # Apply manifest
                       kubectl apply -f go_service/go-service.yaml
                       kubectl apply -f node.js_service/nodejs-service.yaml
                    """
                }
            }
        }
    }
}
