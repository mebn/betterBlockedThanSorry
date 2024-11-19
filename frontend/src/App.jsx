import { useEffect, useRef, useState } from "react";
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
    days: "",
    hours: "",
    minutes: "",
    seconds: "",
  });

  const [blocklist, setBlocklist] = useState([
    "svt.se",
    "reddit.com",
    "youtube.com",
  ]);

  const [isRunning, setIsRunning] = useState(false);
  const [endTime, setEndTime] = useState(0);
  const [dialogWebsite, setDialogWebsite] = useState("");

  const dialogRef = useRef(null);

  const toggleDialog = () => {
    if (!dialogRef.current) {
      return;
    }

    setDialogWebsite("");

    dialogRef.current.hasAttribute("open")
      ? dialogRef.current.close()
      : dialogRef.current.showModal();
  };

  const resetBlocktime = () => {
    setBlocktime({
      days: "",
      hours: "",
      minutes: "",
      seconds: "",
    });
  };

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
            resetBlocktime();
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

  const mainLayoutStyle = {
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
  };

  return (
    <div>
      <dialog
        style={{
          position: "fixed",
          top: "50%",
          left: "50%",
          transform: "translate(-50%, -50%)",
          padding: "20px",
          border: "none",
          borderRadius: "10px",
          background: "white",
          maxWidth: "400px",
        }}
        ref={dialogRef}
        onClick={(e) => {
          if (e.currentTarget == e.target) {
            toggleDialog();
          }
        }}
      >
        <div
          style={{
            display: "flex",
            flexFlow: "column",
            gap: "20px",
          }}
        >
          <h3>Add new website to block</h3>
          <input
            style={{
              background: "#EFEFEF",
              border: "none",
              borderRadius: "10px",
              padding: "10px",
              fontSize: "16px",
              outline: "none",
            }}
            type="text"
            tabIndex={0}
            placeholder="www.website.com"
            value={dialogWebsite}
            onChange={(e) => setDialogWebsite(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === "Enter") {
                setBlocklist([...blocklist, dialogWebsite]);
                toggleDialog();
              }
            }}
          />

          <div
            style={{
              display: "flex",
              flexFlow: "row",
              justifyContent: "space-between",
            }}
          >
            <button
              style={{
                padding: "15px 20px",
                border: "none",
                cursor: "pointer",
                borderRadius: "10px",
                background: "#FF6B6B",
                color: "white",
                fontWeight: "bold",
              }}
              onClick={toggleDialog}
            >
              Cancel
            </button>
            <button
              style={{
                padding: "15px 20px",
                border: "none",
                cursor: "pointer",
                borderRadius: "10px",
                background: "#4E67D6",
                color: "white",
                fontWeight: "bold",
              }}
              onClick={() => {
                setBlocklist([...blocklist, dialogWebsite]);
                toggleDialog();
              }}
            >
              Add
            </button>
          </div>
        </div>
      </dialog>

      <div style={mainLayoutStyle}>
        <div style={{ gridArea: "top" }}>
          <Title buttonTitle="Give Feedback" />
        </div>

        <div style={{ gridArea: "blocktime" }}>
          <Column
            title="Blocktime"
            buttonTitle="Reset"
            isRunning={isRunning}
            onClick={resetBlocktime}
          >
            <Counter
              title="Days"
              value={blocktime.days}
              disabled={isRunning}
              blocktime={blocktime}
              blocktimeEntry={"days"}
              setBlocktime={setBlocktime}
              maxVal={7}
            />
            <Counter
              title="Hours"
              value={blocktime.hours}
              disabled={isRunning}
              blocktime={blocktime}
              blocktimeEntry={"hours"}
              setBlocktime={setBlocktime}
              maxVal={23}
            />
            <Counter
              title="Minutes"
              value={blocktime.minutes}
              disabled={isRunning}
              blocktime={blocktime}
              blocktimeEntry={"minutes"}
              setBlocktime={setBlocktime}
              maxVal={59}
            />
            <Counter
              title="Seconds"
              value={blocktime.seconds}
              disabled={isRunning}
              blocktime={blocktime}
              blocktimeEntry={"seconds"}
              setBlocktime={setBlocktime}
              maxVal={59}
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
            onClick={toggleDialog}
          >
            <div
              style={{
                display: "flex",
                flexFlow: "column",
                gap: "10px",
                overflowY: "auto",
              }}
            >
              {blocklist.map((value, key) => (
                <Entry
                  text={value}
                  isRunning={isRunning}
                  key={key}
                  onClick={() => {
                    // remove element from blocklist
                    const updatedBlocklist = blocklist.filter(
                      (_, i) => i !== key
                    );
                    setBlocklist(updatedBlocklist);
                  }}
                />
              ))}
            </div>
          </Column>
        </div>
      </div>
    </div>
  );
}

export default App;
