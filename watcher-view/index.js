const clientId = "mqttjs_" + Math.random().toString(16).substr(2, 8);

const host = "ws://localhost:8883";
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
  client.subscribe("topic/secret", { qos: 0 });
});

client.on("message", (topic, message, packet) => {
  const msg = JSON.parse(message.toString());
  debugger;
  const date = new Date();
  $("#main-container").append(
    `<p>Message received at ${date.getHours()}:${date.getMinutes()}</p>`
  );
  $("#main-container").append(`<p>Topic => ${topic}</p>`);
  $("#main-container").append(`<p>Action => ${message.toString()}</p>`);
  $("#main-container").append(`</br>`);
});
