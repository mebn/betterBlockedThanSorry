import { useEffect, useState } from "react";
import { GetDaemonRunningStatus, GetEndTime } from "../wailsjs/go/main/App";
import Runner from "./screens/Runner";
import Home from "./screens/Home";

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

  if (isRunning) {
    return <Runner currentTime={currentTime} endTime={endTime} />;
  }

  return (
    <Home
      blocklist={blocklist}
      setBlocklist={setBlocklist}
      blocktime={blocktime}
      setBlocktime={setBlocktime}
      setEndTime={setEndTime}
      setCurrentTime={setCurrentTime}
      setIsRunning={setIsRunning}
    />
  );
}

export default App;
