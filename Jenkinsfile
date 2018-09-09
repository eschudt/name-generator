pipeline {
    stages {
        agent { docker { image 'golang:1.10-alpine' } }
        stage('build') {
            steps {
                sh 'go --version'
            }
        }
    }
}
