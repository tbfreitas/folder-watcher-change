import React from "react";
import logo from "./logo.svg";
import "./App.css";
import mqtt from "mqtt";

interface mod {
  topic: string;
  message: string;
}

function App() {
  const [modifications, setModifications] = React.useState([] as any[]);
  const [connectStatus, setConnectStatus] = React.useState("");
  const [client, setClient] = React.useState(null as any);

  const mqttConnect = () => {
    setConnectStatus("Connecting");
    setClient(mqtt.connect(process.env.REACT_APP_URL_MQTT));
  };

  React.useEffect(() => {
    if (client) {
      client.subscribe("topic/secret");
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
      client.on("message", (topic: string, message: string) => {
        debugger;
        const payload = { topic, message: message.toString() };
        setModifications([...modifications, payload]);
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
        <p>{process.env.REACT_APP_URL_MQTT}</p>
        <p>{process.env.REACT_APP_TOPIC_URL}</p>

        {modifications.map((md: mod) => (
          <p>{md.message}</p>
        ))}
      </header>
    </div>
  );
}

export default App;
