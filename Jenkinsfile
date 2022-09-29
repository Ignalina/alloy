pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh '''whoami
export PATH=/usr/local/go/bin:$PATH
sh /var/lib/jenkins/.cargo/env
go version
make
'''
      }
    }

  }
}