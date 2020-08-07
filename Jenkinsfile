node {
    checkout scm

    def customImage = docker.build("nagrihussain/bootstrap-application:${env.BUILD_ID}", "./build")

    customImage.inside {
        sh 'make test'
    }
}