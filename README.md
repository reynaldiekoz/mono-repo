# Mono Repo Use Case

## Use Case
Develop a simple Node.js and Go service, implementing a "Hello" endpoint as the simplest example. Implement CI/CD pipelines to deploy the code on Kubernetes.

## Approach
- **CI/CD:** Jenkins
- **Kubernetes:** AWS EKS
- **Registry:** AWS ECR
- **Others:** AWS Load Balancer, AWS Route 53

## Solutions
1. Develop a simple "Hello" service for both Node.js and Golang and publish it on a public repository.
2. Prepare the infrastructure for this use case. Deploy the services on AWS EKS and use AWS ECR as the Docker image registry.
3. Set up the Jenkins instance for continuous integration and deployment.
4. **Jenkins CI/CD Stage Workflow:**
   - Jenkins checks out the code from GitHub using a GitHub webhook, detecting pushes, and initiating the deployment.
   - Build the code.
   - Build Docker images using a multistage process.
   - Push Docker images to AWS ECR.
   - Deploy services on EKS using ECR images, exposing the EKS service with an AWS Load Balancer.
5. Configure Route 53 for DNS access to the services. This enables accessing the services using DNS.
