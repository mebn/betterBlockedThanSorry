import { useEffect, useState } from "react";
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
      if (daemonStatus) {
        setIsRunning(daemonStatus);
      }

      if (isRunning) {
        let fetchedEndTime = endTime;

        if (fetchedEndTime == 0) {
          fetchedEndTime = await GetEndTime();
          console.log("fetchedendtime", fetchedEndTime);
          setEndTime(fetchedEndTime);
        }

        intervalId = setInterval(() => {
          let time = getCurrentTime();
          setCurrentTime(time);

          if (fetchedEndTime - time <= 0) {
            setIsRunning(false);
          }
        }, 1000);
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
    setEndTime(newEndTime);

    const daemonStatus = await GetDaemonRunningStatus();
    if (daemonStatus) {
      setCurrentTime(getCurrentTime());
      setIsRunning(daemonStatus);
    }
  };

  if (isRunning) {
    return (
      <div style={{ color: "white" }}>
        <h1 className="title">BetterBlockedThanSorry</h1>
        <h2>The blocker will stop in {endTime - currentTime} seconds.</h2>
      </div>
    );
  } else {
    return (
      <div>
        <h1 className="title">BetterBlockedThanSorry</h1>
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
    );
  }
}

export default App;
