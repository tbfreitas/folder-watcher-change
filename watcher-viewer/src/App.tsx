import React from "react";
import logo from "./logo.svg";
import "./App.css";
const Kafka = require("kafka-node");

function App() {
  var Consumer = Kafka.Consumer;
  var Client = Kafka.KafkaClient;
  var client = new Client({ kafkaHost: "localhost:9092" });
  // var topics = [{ topic: "t1", partition: 0 }];
  // var options = {
  //   autoCommit: false,
  //   fetchMaxWaitMs: 1000,
  //   fetchMaxBytes: 1024 * 1024,
  // };

  // var consumer = new Consumer(client, topics, options);
  // consumer.on("message", function (message: string) {
  //   console.log(message);
  // });

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
