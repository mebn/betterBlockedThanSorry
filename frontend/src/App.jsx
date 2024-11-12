import { useEffect, useState } from "react";
import "./App.css";
import {
  StartBlocker,
  GetDaemonRunningStatus,
  GetEndTime,
} from "../wailsjs/go/main/App";

function App() {
  const getCurrentTime = () => Math.floor(Date.now() / 1000);

  const [blocktime, setBlocktime] = useState(60);
  const [blocklist, setBlocklist] = useState([
    "svt.se",
    "reddit.com",
    "youtube.com",
  ]);

  const [isRunning, setIsRunning] = useState(false);
  const [currentTime, setCurrentTime] = useState(getCurrentTime());
  const [endTime, setEndTime] = useState(0);

  useEffect(() => {
    let intervalId;

    const initialize = async () => {
      const daemonStatus = await GetDaemonRunningStatus();
      setIsRunning(daemonStatus);
      if (daemonStatus) {
        const fetchedEndTime = await GetEndTime();
        setEndTime(fetchedEndTime);

        intervalId = setInterval(() => {
          setCurrentTime(getCurrentTime());
        }, 1000);

        const timeRemaining = fetchedEndTime - getCurrentTime();
        setTimeout(() => {
          setIsRunning(false);
        }, timeRemaining * 1000);
      }
    };

    initialize();

    return () => {
      clearInterval(intervalId);
    };
  }, [isRunning]);

  const startBlocker = async () => {
    const preIsRunning = await GetDaemonRunningStatus();
    if (preIsRunning) {
      console.log("Program is already running");
      return;
    }

    const newEndTime = await StartBlocker(blocktime, blocklist);
    setCurrentTime(getCurrentTime());
    setEndTime(newEndTime);
    setIsRunning(await GetDaemonRunningStatus());

    if (isRunning && newEndTime !== 0) {
      const intervalId = setInterval(() => {
        setCurrentTime(getCurrentTime());
      }, 1000);

      const timeRemaining = newEndTime - getCurrentTime();
      setTimeout(() => {
        setIsRunning(false);
        clearInterval(intervalId);
      }, timeRemaining * 1000);
    }
  };

  if (isRunning) {
    return (
      <div id="App">
        <div>
          <h1>Program is already running.</h1>
          <h1>End time: {endTime}</h1>
          <h1>Current time: {currentTime}</h1>
          <h1>Time remaining: {endTime - currentTime}</h1>
        </div>
      </div>
    );
  } else {
    return (
      <div id="App">
        <div>
          <input
            id="blocklist"
            onChange={(e) => setBlocklist(e.target.value.split(","))}
            value={blocklist}
            autoComplete="off"
            name="blocklist"
            type="text"
          />
          <br />
          <input
            id="blocktime"
            onChange={(e) => setBlocktime(parseInt(e.target.value))}
            value={blocktime}
            autoComplete="off"
            name="blocktime"
            type="number"
          />
          <br />
          <button onClick={startBlocker}>Start blocker</button>
        </div>
      </div>
    );
  }
}

export default App;
