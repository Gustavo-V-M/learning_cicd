node {
    checkout scm
    pipeline {
        agent any

        stages {
            stage('Build') {
                steps {
                    go build 
                    archiveArtifacts artifacts: 'http_example', fingerprint: true
                }
            }
            stage('Test') {
                steps {
                    go test
                }
            }
            stage('Deploy') {
                steps {
                    echo "TODO: implement deploy with ansible"
                }
            }
        }
    }
}