pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh '''whoami
export PATH=$PATH:/usr/local/go/bin:~/.cargo/bin
#sh "$HOME/.cargo/env"
go version
make
'''
      }
    }

  }
}