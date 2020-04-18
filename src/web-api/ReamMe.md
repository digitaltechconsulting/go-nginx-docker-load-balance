# Load Balancing Go Applications with NGINX and Docker

```docker 
docker build -t nginx-as-load-balancer-web-api:dev .
```
> Above command will generate image named **web-api-load-balance:dev**

Now we will create two containers api1 and api2 exposing two differnt ports **4001** and **4002** to external world.

```docker
docker run --name apiserver1 -d -p 4001:4000 -e DOCKER_INSTANCE_ID='server1'   web-api-load-balance:dev
```
Simillarly, run second instance of the docker container on port **4002**
```docker
docker run --name apiserver2 -d -p 4002:4000 -e DOCKER_INSTANCE_ID='server2'   web-api-load-balance:dev
```
With this we have our two worker container ready to be handled by our load balancer.  
Below docker command will create a docker imange named **nginxbalancer:001** 

```docker
docker build -t nginxbalancer:001 .
```

Now we need to create another container which will host our **nginx** server, our load balancer.

```docker 
docker run --name mybalancer -d -p 5000:80 nginxbalancer:001
```

> Note that in our **nginx.conf** file we have hardcoded ip address **172.17.0.1**, you can find more on this [here](https://auth0.com/blog/load-balancing-nodejs-applications-with-nginx-and-docker/)

