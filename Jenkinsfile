node {
    checkout scm
    docker.withRegistry('', 'dockerRegistryCredentials'){
    def dockerfile = 'build/Dockerfile'
    def appImage = docker.build("nagrihussain/bootstrap-application:${env.BUILD_ID}",  "-f ${dockerfile} .")

    customImage.push()
    }
}