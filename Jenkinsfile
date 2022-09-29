pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh '''whoami
export PATH=/usr/local/go/bin:$PATH
ls -l  /var/lib/jenkins/.cargo
sh /var/lib/jenkins/.cargo/env
go version
make
'''
      }
    }

  }
}