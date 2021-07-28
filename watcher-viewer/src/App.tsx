import React from "react";
import logo from "./logo.svg";
import "./App.css";
import mqtt from "mqtt";

function App() {
  React.useEffect(() => {
    const client = mqtt.connect("ws://localhost:8883");
    client.subscribe("topic/secret");

    client.on("message", (topic, payload, packet) => {
      console.log(payload.toString());
    });
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
