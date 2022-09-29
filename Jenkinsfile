pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh '''export PATH=$PATH:/usr/local/go/bin
$HOME/.cargo/env
go version
make
'''
      }
    }

  }
}