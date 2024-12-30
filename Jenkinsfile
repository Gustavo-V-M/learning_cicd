// Declarative //
pipeline {
    agent any

    tools {
        go 'mygo'
    }

    stages {
        stage('Build') {
            steps {
                sh 'go build'
            }
        }
        stage('Test') {
            steps {
                sh 'go test -v'
            }
            post {
                success {
                    archiveArtifacts artifact: 'http_example', fingerprint: true
                }
            }
        }
    }
}
