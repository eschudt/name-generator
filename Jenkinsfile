pipeline {
    agent { docker { image 'golang:1.10-alpine' } }
    stages {
        stage('Build') {
            steps {
                sh 'go version'
                sh 'pwd'
                sh 'ls -l'
                sh 'apk update && apk add --no-cache git make'
                sh 'wget https://raw.githubusercontent.com/golang/dep/master/install.sh'
                sh 'chmod 777 install.sh'
                sh './install.sh'
                sh 'mkdir -p /go/src/github.com/eschudt/name-generator'
                sh 'cp -r * /go/src/github.com/eschudt/name-generator/'
                sh 'cd /go/src/github.com/eschudt/name-generator/'
                sh 'pwd'
                sh 'ls -l'
                sh 'dep ensure'
            }
        }
        stage('Test') {
            steps {
                sh 'make test'
            }
        }
        stage('Deploy - Dev') {
            steps {
                echo 'Deployed to Dev'
            }
        }
        stage('Sanity check') {
            steps {
                input "Does the Dev environment look ok?"
            }
        }
        stage('Deploy - Production') {
            steps {
                echo 'Deployed to Prod'
            }
        }
    }
    post {
        always {
            echo 'Job finished'
        }
        success {
            echo 'I succeeeded!'
        }
        unstable {
            echo 'I am unstable :/'
        }
        failure {
            echo 'I failed :('
        }
        changed {
            echo 'Things were different before...'
        }
    }
}
