pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh '''whoami
export PATH=/usr/local/go/bin:$PATH
sh $HOME/.cargo/env
go version
make
'''
      }
    }

  }
}