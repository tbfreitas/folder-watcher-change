const clientId = "mqttjs_" + Math.random().toString(16).substr(2, 8);

const host = "ws://localhost:8883";
const msgs = [];
const options = {
  keepalive: 60,
  clientId: clientId,
  protocolId: "MQTT",
  protocolVersion: 4,
  clean: true,
  reconnectPeriod: 1000,
  connectTimeout: 30 * 1000,
  will: {
    topic: "WillMsg",
    payload: "Connection Closed abnormally..!",
    qos: 0,
    retain: false,
  },
};

console.log("Connecting mqtt client");
const client = mqtt.connect(host, options);

client.on("error", (err) => {
  console.log("Connection error: ", err);
  client.end();
});

client.on("reconnect", () => {
  console.log("Reconnecting...");
});

client.on("connect", () => {
  console.log("Client connected:" + clientId);
  // Subscribe
  client.subscribe("topic/secret", { qos: 0 });
});

client.on("message", (topic, message, packet) => {
  msgs.push(message.toString());
  msgs.map((msg, index) => {
    console.log(
      "Number : " +
        index +
        "    Received Message: " +
        msg +
        "\nOn topic: " +
        topic
    );
  });
});
