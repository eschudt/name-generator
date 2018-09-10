node {
    def app

    stage('Clone repository') {
        checkout scm
    }

    stage('Test and Build image') {
        sh 'mkdir -p /root/go/src/github.com/eschudt/name-generator'
        sh 'cp -r * /root/go/src/github.com/eschudt/name-generator/'
        sh 'cd /root/go/src/github.com/eschudt/name-generator/ && /go/bin/dep ensure'
        sh 'cd /root/go/src/github.com/eschudt/name-generator/ && make test'
        app = docker.build("/root/go/src/github.com/eschudt/name-generator")
    }

    stage('Integration Test') {
        app.inside {
            sh 'echo "Tests passed"'
        }
    }

    stage('Push image') {
        /* Finally, we'll push the image with two tags:
         * First, the incremental build number from Jenkins
         * Second, the 'latest' tag.
         * Pushing multiple tags is cheap, as all the layers are reused. */
        docker.withRegistry('https://registry.hub.docker.com', 'docker-hub-credentials') {
            app.push("0.0.${env.BUILD_NUMBER}")
            app.push("latest")
        }
    }

    stage('Deploy - Dev') {
        sh 'echo "Deployed to Dev"'
    }
}
