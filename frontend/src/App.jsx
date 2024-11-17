import { useEffect, useState } from "react";
import {
  GetDaemonRunningStatus,
  GetEndTime,
  StartBlocker,
} from "../wailsjs/go/main/App";
import StartButton from "./components/StartButton";
import Counter from "./components/Counter";
import Column from "./components/Column";
import Entry from "./components/Entry";
import Title from "./components/Title";

function App() {
  const getCurrentTime = () => Math.floor(Date.now() / 1000);

  const [blocktime, setBlocktime] = useState({
    days: null,
    hours: null,
    minutes: null,
    seconds: null,
  });
  const [blocklist, setBlocklist] = useState([
    "svt.se",
    "reddit.com",
    "youtube.com",
  ]);

  const [isRunning, setIsRunning] = useState(false);
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
          setEndTime(fetchedEndTime);
        }

        intervalId = setInterval(() => {
          let time = getCurrentTime();
          let timeLeft = fetchedEndTime - time;

          setBlocktime({
            days: Math.floor(timeLeft / 86400) || "",
            hours: Math.floor((timeLeft % 86400) / 3600) || "",
            minutes: Math.floor((timeLeft % 3600) / 60) || "",
            seconds: timeLeft % 60 || "",
          });

          if (timeLeft <= 0) {
            setBlocktime({
              days: "",
              hours: "",
              minutes: "",
              seconds: "",
            });

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

    let calculatedEndTime =
      (blocktime.days == "" ? 0 : blocktime.days * 24 * 60 * 60) +
      (blocktime.hours == "" ? 0 : blocktime.hours * 60 * 60) +
      (blocktime.minutes == "" ? 0 : blocktime.minutes * 60) +
      (blocktime.seconds == "" ? 0 : blocktime.seconds);

    const newEndTime = await StartBlocker(calculatedEndTime, blocklist);
    setEndTime(newEndTime);

    const daemonStatus = await GetDaemonRunningStatus();
    if (daemonStatus) {
      setIsRunning(daemonStatus);
    }
  };

  return (
    <div
      style={{
        display: "grid",
        gridTemplateAreas: `
          "top top"
          "blocktime blocklist"
          "start blocklist"
        `,
        gridTemplateColumns: "1fr 1fr",
        gap: "20px",
        padding: "20px",
        height: "100vh",
        boxSizing: "border-box",
      }}
    >
      <div style={{ gridArea: "top" }}>
        <Title buttonTitle="Give Feedback" />
      </div>

      <div style={{ gridArea: "blocktime" }}>
        <Column
          title="Blocktime"
          buttonTitle="Reset"
          isRunning={isRunning}
          onClick={() => {
            setBlocktime({
              days: "",
              hours: "",
              minutes: "",
              seconds: "",
            });
          }}
        >
          <Counter
            text="Days"
            value={blocktime.days}
            isRunning={isRunning}
            onChange={(e) => {
              let val = parseInt(e.target.value, 10);
              if (!isNaN(val)) {
                val = val > 7 ? 7 : val;
                val = val >= 0 ? val : 0;
              }
              setBlocktime({
                ...blocktime,
                days: val,
              });
            }}
          />
          <Counter
            text="Hours"
            value={blocktime.hours}
            isRunning={isRunning}
            onChange={(e) => {
              let val = parseInt(e.target.value, 10);
              if (!isNaN(val)) {
                val = val > 23 ? 23 : val;
                val = val >= 0 ? val : 0;
              }
              setBlocktime({
                ...blocktime,
                hours: val,
              });
            }}
          />
          <Counter
            text="Minutes"
            value={blocktime.minutes}
            isRunning={isRunning}
            onChange={(e) => {
              let val = parseInt(e.target.value, 10);
              if (!isNaN(val)) {
                val = val > 59 ? 59 : val;
                val = val >= 0 ? val : 0;
              }
              setBlocktime({
                ...blocktime,
                minutes: val,
              });
            }}
          />
          <Counter
            text="Seconds"
            value={blocktime.seconds}
            isRunning={isRunning}
            onChange={(e) => {
              let val = parseInt(e.target.value, 10);
              if (!isNaN(val)) {
                val = val > 59 ? 59 : val;
                val = val >= 0 ? val : 0;
              }
              setBlocktime({
                ...blocktime,
                seconds: val,
              });
            }}
          />
        </Column>
      </div>

      <div style={{ gridArea: "start" }}>
        <StartButton
          text="Start Blocker"
          onClick={startBlocker}
          isRunning={isRunning}
        />
      </div>

      <div
        style={{
          gridArea: "blocklist",
          overflowY: "auto",
        }}
      >
        <Column
          title="Blocklist"
          buttonTitle="Add"
          isRunning={isRunning}
          onClick={() => {}}
        >
          <div
            style={{
              display: "flex",
              flexFlow: "column",
              gap: "10px",
              overflowY: "auto",
            }}
          >
            <Entry text="reddit.com" isRunning={isRunning} />
            <Entry text="youtube.com" isRunning={isRunning} />
            <Entry text="reddit.com" isRunning={isRunning} />
            <Entry text="youtube.com" isRunning={isRunning} />
            <Entry text="reddit.com" isRunning={isRunning} />
            <Entry text="youtube.com" isRunning={isRunning} />
            <Entry text="reddit.com" isRunning={isRunning} />
            <Entry text="youtube.com" isRunning={isRunning} />
            <Entry text="reddit.com" isRunning={isRunning} />
            <Entry text="youtube.com" isRunning={isRunning} />
            <Entry text="reddit.com" isRunning={isRunning} />
            <Entry text="youtube.com" isRunning={isRunning} />
          </div>
        </Column>
      </div>
    </div>
  );
}

export default App;
