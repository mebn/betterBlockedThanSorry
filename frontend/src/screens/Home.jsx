import { useEffect, useState } from "react";
import {
  GetEndtimeDB,
  SetBlocklistDB,
  GetBlocklistDB,
  StartBlocker,
  // GetCurrentTime,
} from "../../wailsjs/go/main/App";
import StartButton from "../components/StartButton";
import Counter from "../components/Counter";
import Column from "../components/Column";
import Entry from "../components/Entry";
import Title from "../components/Header";
import Dialog from "../components/Dialog";

function App() {
  const [isDialogOpen, setIsDialogOpen] = useState(false);
  const [endTime, setEndTime] = useState(0);
  const [isRunning, setIsRunning] = useState(true);
  const [blocklist, setBlocklist] = useState([]);
  const [blocktime, setBlocktime] = useState({
    days: "",
    hours: "",
    minutes: "",
    seconds: "",
  });

  const resetBlocktime = () => {
    setBlocktime({ days: "", hours: "", minutes: "", seconds: "" });
  };

  // used to get blocklist
  useEffect(() => {
    const f = async () => {
      const newBlocklist = await GetBlocklistDB();
      setBlocklist(newBlocklist);
    };

    f();
  }, []);

  useEffect(() => {
    let intervalId;

    if (!isRunning) {
      return;
    }

    const initialize = async () => {
      let fetchedEndTime = await GetEndtimeDB();
      let currentTime = Math.floor(Date.now() / 1000);

      setIsRunning(fetchedEndTime >= currentTime);

      if (!isRunning) {
        return;
      }

      setEndTime(fetchedEndTime);

      // can this be async?
      intervalId = setInterval(() => {
        let currentTime = Math.floor(Date.now() / 1000);
        // let currentTime = await GetCurrentTime();
        let timeLeft = fetchedEndTime - currentTime;

        setBlocktime({
          days: Math.floor(timeLeft / 86400) || "",
          hours: Math.floor((timeLeft % 86400) / 3600) || "",
          minutes: Math.floor((timeLeft % 3600) / 60) || "",
          seconds: timeLeft % 60 || "",
        });

        if (timeLeft <= 0) {
          resetBlocktime();
          setEndTime(0);
          setIsRunning(false);
          clearInterval(intervalId);
        }
      }, 1000);
    };

    initialize();

    return () => {
      clearInterval(intervalId);
    };
  }, [isRunning]);

  const startBlocker = async () => {
    let calculatedEndTime =
      (blocktime.days == "" ? 0 : blocktime.days * 24 * 60 * 60) +
      (blocktime.hours == "" ? 0 : blocktime.hours * 60 * 60) +
      (blocktime.minutes == "" ? 0 : blocktime.minutes * 60) +
      (blocktime.seconds == "" ? 0 : blocktime.seconds);

    const newEndTime = await StartBlocker(calculatedEndTime, blocklist);

    if (newEndTime == 0) {
      return;
    }

    setEndTime(newEndTime);
    setIsRunning(true);
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
      <Dialog
        isOpen={isDialogOpen}
        onClose={() => setIsDialogOpen(false)}
        onAddWebsite={(website) => {
          const newBlocklist = [...blocklist, website];
          SetBlocklistDB(newBlocklist);
          setBlocklist(newBlocklist);
        }}
      />

      <div style={mainLayoutStyle}>
        <div style={{ gridArea: "top" }}>
          <Title buttonText="Give Feedback" />
        </div>

        <div style={{ gridArea: "blocktime" }}>
          <Column
            title="Blocktime"
            buttonText="Reset"
            disabled={isRunning}
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
            textEnabled="Start Blocker"
            textDisabled="Blocker started"
            onClick={startBlocker}
            disabled={isRunning}
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
            buttonText="Add"
            onClick={() => setIsDialogOpen(true)}
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
                  title={value}
                  buttonText="&#x1F5D1;"
                  hidden={isRunning}
                  disabled={isRunning}
                  monochrome={true}
                  key={key}
                  onClick={() => {
                    // remove element from blocklist
                    const newBlocklist = blocklist.filter((_, i) => i !== key);
                    SetBlocklistDB(newBlocklist);
                    setBlocklist(newBlocklist);
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
