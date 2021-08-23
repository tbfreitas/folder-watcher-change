# folder-watcher-change

_FWC_ is a case study architeture that works as a datawarehouse. What it means ? There's a observable actor listening
any file change, or even a whole folder, just like in the sample. When it`s occurs, a payload is send throught a MQTT broker.
On the edge of the bus, we got a Javascript listener the actions, and using JQuery, manipulate the DOM to show it.

The Broker is a broadcaster software, thus making easy connect from a generic way any new application. Just connecting using the right protocol(tcp/wss) and port(8883:1883). Rather than create or modify the architetura on each new entry, just make the appointment to this address.

# Architeture

# Bring it up

# Run it locally
