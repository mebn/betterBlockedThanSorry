import Title from "../components/Title";

import {
  StartBlocker,
  GetDaemonRunningStatus,
} from "../../wailsjs/go/main/App";
import StartButton from "../components/StartButton";
import Counter from "../components/Counter";
import Column from "../components/Column";
import Entry from "../components/Entry";

export default ({
  blocklist,
  setBlocklist,
  blocktime,
  setBlocktime,
  setEndTime,
  setCurrentTime,
  setIsRunning,
}) => {
  const getCurrentTime = () => Math.floor(Date.now() / 1000);

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
        <Column title="Blocktime" buttonTitle="Reset">
          <Counter text="Days" />
          <Counter text="Hours" />
          <Counter text="Minutes" />
          <Counter text="Seconds" />
        </Column>
      </div>

      <div style={{ gridArea: "start" }}>
        <StartButton text="Start Blocker" />
      </div>

      <div
        style={{
          gridArea: "blocklist",
          overflowY: "auto",
        }}
      >
        <Column title="Blocklist" buttonTitle="Add">
          <div
            style={{
              display: "flex",
              flexFlow: "column",
              gap: "10px",
              overflowY: "auto",
            }}
          >
            <Entry text="reddit.com" />
            <Entry text="youtube.com" />
            <Entry text="reddit.com" />
            <Entry text="youtube.com" />
            <Entry text="reddit.com" />
            <Entry text="youtube.com" />
            <Entry text="reddit.com" />
            <Entry text="youtube.com" />
            <Entry text="reddit.com" />
            <Entry text="youtube.com" />
            <Entry text="reddit.com" />
            <Entry text="youtube.com" />
          </div>
        </Column>
      </div>

      {/* <input
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
      
      <button onClick={startBlocker}>Start blocker</button> */}
    </div>
  );
};
