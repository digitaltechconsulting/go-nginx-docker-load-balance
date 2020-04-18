# Setting up **NGINX**
I have used [this](https://hub.docker.com/_/nginx) image of **NGINX** in this example.

Run below command to create an image on your local machine

> Before setting up this environment, make sure that you have setup [API containers](../web-api/ReadMe.md) first.

```docker
docker build -t nginxbalancer:001 .
```
Once image is ready, now its time to spin up a container which will act as our **loadbalancer**

```docker 
docker run --name mybalancer -d -p 5000:80 nginxbalancer:001
```
> Note that in our **nginx.conf** file we have hardcoded IP address **172.17.0.1**, you can find more on this [here](https://auth0.com/blog/load-balancing-nodejs-applications-with-nginx-and-docker/)

With this you have your NGINX load balancer ready.  To verify the setup open browser of your choice and navigate to http://localhost:5000 

You will see either **Request is served by container : server1** OR **Request is served by container : server2** message depending upon which container is serving the request.