import React from "react";
import logo from "./logo.svg";
import "./App.css";
import mqtt from "mqtt";

function App() {
  const [modifications, setModifications] = React.useState([] as any[]);
  const [connectStatus, setConnectStatus] = React.useState("");
  const [client, setClient] = React.useState(null as any);

  const msg: any[] = [];
  const mqttConnect = () => {
    setConnectStatus("Connecting");
    setClient(mqtt.connect(process.env.REACT_APP_URL_MQTT));
  };

  React.useEffect(() => {
    if (client) {
      client.subscribe("topic/secret", { qos: 2 });
      client.on("connect", () => {
        setConnectStatus("Connected");
      });
      client.on("error", (err: string) => {
        console.error("Connection error: ", err);
        client.end();
      });
      client.on("reconnect", () => {
        setConnectStatus("Reconnecting");
      });
      client.on("message", (topic: string, message: string, packet: any) => {
        msg.push(message.toString());
        setModifications(msg);
      });
    }
  }, [client]);

  React.useEffect(() => {
    mqttConnect();
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        {connectStatus}
        <p>{process.env.REACT_APP_URL_MQTT}</p>
        <p>{process.env.REACT_APP_TOPIC_URL}</p>

        {modifications.map((mod: string) => (
          <p>{mod}</p>
        ))}
      </header>
    </div>
  );
}

export default App;
