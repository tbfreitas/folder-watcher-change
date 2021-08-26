# folder-watcher-change

_FWC_ is a case study architeture that works as a datawarehouse. What it means ? There's a observable actor listening any file change, or even a whole folder, just like in sample. When it`s occurs, a payload is send throught a MQTT broker.
On the edge of the bus, we got a Javascript listening the actions, and manipulating the DOM using **Jquery** to show the changes.
<br/>
The Broker is a broadcaster software, thus making easy connect from a generic way any new application. Just connect using the right protocol - tcp/wss- and port -8883:1883-. Rather than create or modify the architeture on each new entry, just make the appointment to this address.

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
