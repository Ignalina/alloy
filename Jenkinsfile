pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh '''export PATH=/usr/local/go/bin:$PATH
sh $HOME/.cargo/env
go version
make
'''
      }
    }

  }
}