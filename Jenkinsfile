pipeline {
  environment {
      imageName = "thenerdyhamster/demo-pipeline"
      registryCredentials = "docker-hub-credentials"
      dockerImage = ''
  }
  agent any
  stages {
    stage("Clone git repo") {
      steps {
        script {
          git([url: 'git@github.com:TheNerdyHamster/jenkins-docker-build.git', branch: 'main', credentialsId: 'git'])
        }
      }
    }
    stage("Build image") {
      steps {
        script {
          dockerImage = docker.build imageName
        }
      }
    }
    stage("Deploy image") {
      steps {
        script {
          docker.withRegistry('', registryCredentials) {
              dockerImage.push('$BUILD_NUMBER')
              dockerImage.push('latest')
          }
        }
      }
    }
    stage("Clean up build") {
      steps {
        script {
          sh "docker rmi $imageName:$BUILD_NUMBER"
          sh "docker rmi $imageName:latest"
        }
      }
    }
  }
}
