# Demo application that responds with a JSON payload upon receiving a request.


## Jenkins Setup 

docker network create jenkins  
docker volume create jenkins-docker-certs  
docker volume create jenkins-data  

```docker container run   --name jenkins-docker   --rm   --detach   --privileged   --network jenkins   --network-alias docker   --env DOCKER_TLS_CERTDIR=/certs   --volume jenkins-docker-certs:/certs/client   --volume jenkins-data:/var/jenkins_home   --publish 2376:2376   docker:dind```

```docker container run   --name jenkins-blueocean   --rm   --detach   --network jenkins   --env DOCKER_HOST=tcp://docker:2376   --env DOCKER_CERT_PATH=/certs/client   --env DOCKER_TLS_VERIFY=1   --publish 8080:8080   --publish 50000:50000   --volume jenkins-data:/var/jenkins_home   --volume jenkins-docker-certs:/certs/client:ro   jenkinsci/blueocean```

```docker exec jenkins-blueocean cat /var/jenkins_home/secrets/initialAdminPassword```

copy the output  

go to - `http://localhost:8080` and paste the output in the Jenkins UI to unlock Jenkins. Choose default setup and done.


## Set a credential with ID `dockerRegistryCredentials` which is the credentials of Your docker hub account

```$ docker exec jenkins-blueocean /bin/bash -c 'mkdir -p $JENKINS_HOME/jobs/Application\ Build\ Pipeline'```
```$ docker cp build/jenkins/config.xml jenkins-blueocean:/var/jenkins_home/jobs/Application\ Build\ Pipeline/```

Then visit `http://localhost:8080/restart`

