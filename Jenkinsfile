pipeline {
    agent { docker { image 'golang:1.10-alpine' } }
    stages {
        stage('Build') {
            steps {
                sh 'apk update && apk add --no-cache git make'
                sh 'wget -q https://raw.githubusercontent.com/golang/dep/master/install.sh'
                sh './install.sh'
                sh 'mkdir -p /go/src/github.com/eschudt/name-generator'
                sh 'cp -r * /go/src/github.com/eschudt/name-generator/'
                sh 'export GOPATH=/go'
                sh 'cd /go/src/github.com/eschudt/name-generator/ && dep ensure'
            }
        }
        stage('Test and Push') {
            steps {
                sh 'cd /go/src/github.com/eschudt/name-generator/ && make test && make build'
                sh 'docker push eschudt/name-generator'
            }
        }
        stage('Deploy - Dev') {
            steps {
                echo 'Deployed to Dev'
            }
        }
        stage('Sanity check') {
            steps {
                input "Deploy to production?"
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
