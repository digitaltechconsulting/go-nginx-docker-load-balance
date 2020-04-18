# Setting up API container(s)

```docker 
docker build -t nginx-as-load-balancer-web-api:dev .
```
> Above command will generate image named **web-api-load-balance:dev**

Now we will create two containers **api1** and **api2** exposing two different ports **4001** and **4002** to external world.  These two containers will be our two workers.

```docker
docker run --name apiserver1 -d -p 4001:4000 -e DOCKER_INSTANCE_ID='server1'   web-api-load-balance:dev
```
Similarly, run second command to create another instance of the same docker container listening on port **4002**
```docker
docker run --name apiserver2 -d -p 4002:4000 -e DOCKER_INSTANCE_ID='server2'   web-api-load-balance:dev
```
With this we have our two worker containers ready to be handled by our load balancer.  Now it is time to [setup](../nginx/Readme.md) our **NGINX** server as load balancer.




