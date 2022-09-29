pipeline {
  agent any
  stages {
    stage('build') {
      steps {
        sh '''$HOME/.cargo/env
make
'''
      }
    }

  }
}