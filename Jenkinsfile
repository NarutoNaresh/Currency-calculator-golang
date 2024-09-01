pipeline {
    agent any
    environment{
        def IMG_TAG = sh(script: "date +%y%m%d%H%M%S", returnStdout: true).trim()
    }

    stages {
        stage('Checkout Project Repo') {
            steps {
                checkout scmGit(branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/NarutoNaresh/Currency-calculator-golang.git']])
            }
        }
        stage('Build Docker Image & push to docker hub') {
            environment{
                DOCKER_IMAGE="narutonaresh/curcalcmain:${IMG_TAG}"
            }
            steps {
                script{
                withDockerRegistry(credentialsId: 'docker-cred') {
                sh 'docker build -t ${DOCKER_IMAGE} -f Dockermultistage .'
                sh 'docker run -d -p 80:5050 --name curcalcx$IMG_TAG $DOCKER_IMAGE'
                sh 'docker push ${DOCKER_IMAGE}'  
                }
                }
                
            }
        }
        stage('Update Deployment file for CD') {
            environment{
                NEW_IMG = 'narutonaresh/curcalcmain'
            }
            steps {
                checkout scmGit(branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/NarutoNaresh/Project-Manifests.git']])
                sh '''
                git checkout main
                git pull origin main
                '''
                sh 'sed -i "s@narutonaresh/curcalcmain:.*\\\$@$NEW_IMG:$IMG_TAG@" Currency-calculator-golang/Deploy.yml'
                
                sh '''
                git status
                git add .
                git status
                git commit -m "add new tag for new build"
                '''
            }
        }
        stage('Push the Changes to Manifest Repo') {
            steps {
                withCredentials([string(credentialsId: 'github-token', variable: 'GITHUB_TOKEN')]) {
                    script {
                        sh 'git config --global user.name "naresh kumar d"'
                        sh "git remote set-url origin https://x-access-token:${GITHUB_TOKEN}@github.com/NarutoNaresh/Project-Manifests.git"
                        sh 'git push origin main'
                    }
                }
            }
        }
        stage('Docker clean resources') {
            steps {
                echo 'Cleaning up Docker resources...'
                sh 'docker stop $(docker ps -a -q)'
                sh 'docker rm $(docker ps -a -q)'
            }
        }
    }
}
