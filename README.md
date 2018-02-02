# Breakfast_docker
Just another virtual environment to learning and testing Docker and Docker Compose

## How to start?

To launch the environment you should check to have a vagrant provider e.g Virtualbox. 

```shell
vagrant up
```

Then you can connecto to virtual machine.

```shell
vagrant ssh
```

The machine has Go, Docker and Docker-compose. To launch the Docker-compose.

```shell
docker-compose up
```

You will see the logs of the two containers to stop the process Ctrl+C. You maybe want to clean your memory so you can remove your containers.

```shell
docker rm $(docker ps -aq)
```