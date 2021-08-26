# folder-watcher-change

Rewriting after Denucci just over my wrote english.

### Architeture

![Arch](./imgs/arch.png)

### Bring it up

As simple as it can, just do on console:

    docker-compose up

So, you gonna download 3 docker images. When it's done, a ngnix will be exposing port 80. Docker will make a fast forward from port 8000 to 80. Thus, open your browser on address:

    localhost:8000

### Run it locally

### Disclaimer

The main goal of the project is to test the architeture. So, optimize the build images was **not** a concern in here. We could make better _Dockerfile_, using a image to build and a lightweight with a only one purpose, run de binary. Same thing for the frontend structure, which is pure HTML injecting a javascript file(MQTT lib and JQuery Lib).
