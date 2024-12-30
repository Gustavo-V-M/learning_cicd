// Declarative //
pipeline {
    agent any

    tools {
        go 'mygo'
        terraform 'myterraform'
        ansible 'myansible'
    }

    stages {
        stage ('Checkout ') {
            steps {
                checkout scm
            }
        }
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
                    archiveArtifacts artifacts: 'http_example', fingerprint: true
                    sh '/root/.local/share/pipx/venvs/ansible/bin/ansible --version'
                }
            }
        }
    }
}
