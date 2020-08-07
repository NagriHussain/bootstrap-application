node {
    checkout scm
    def dockerfile = 'build/Dockerfile'
    def customImage = docker.build("nagrihussain/bootstrap-application:${env.BUILD_ID}",  "-f ${dockerfile} .")

    customImage.inside {
        sh 'make test'
    }
}