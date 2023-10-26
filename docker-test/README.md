To deploy the app to AWS you docker-machine can be used as follow:

Note:
ignore AWS_PROFILE=YOUR_PROFILE_NAME if you are using a default aws config profile

1. Run docker-machine to deploy the app:

```shell
    AWS_PROFILE=YOUR_PROFILE_NAME docker-machine create -d amazonec2 \        
        --amazonec2-region eu-west-2 \
        --amazonec2-instance-type "t2.micro" \
        --amazonec2-ssh-keypath PATH_TO_KEY (e.g. ~/.ssh/id_rsa) \
        INSTANCE_NAME (e.g. aws-new-cli-test2)
```

The above is a basic command which can be extended to indicate a EC2 Security Group, ami, etc.

Once the command is succeeded you should have a running EC2 instance in AWS.

2. SSH into the running EC2 instance using docker-machine

```shell
    eval $(AWS_PROFILE=YOUR_PROFILE_NAME docker-machine env aws-new-cli-test2)
```

If the eval command succeeded the rest of the commands would be run inside the EC2 instance

3. If you want to run the app without your terminal staying open, run the following:

```shell
    docker run -d -p 8080:8081 docker-test
```

4. Run Docker build command to build the project container

```shell
    docker build -t docker-test .
```

5. Map the ports and run the container

```shell
    docker run -p 8080:8080 -it docker-test
```

or for detached mode:

```shell
    docker run -d -p 8080:8081 docker-test
```

6. To stop or remove the EC2 instance run

Stop:

```shell
    AWS_PROFILE=YOUR_PROFILE_NAME docker-machine stop aws-new-cli-test2
```

Remove:

```shell
    AWS_PROFILE=YOUR_PROFILE_NAME docker-machine rm INSTANCE_NAME
```

7. After stopping the machine, if you want to revert your local Docker client to point back to your local Docker daemon,
   you can unset the environment variables:

```shell
    eval $(docker-machine env -u)
```