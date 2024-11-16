import Title from "../components/Title";

import {
  StartBlocker,
  GetDaemonRunningStatus,
} from "../../wailsjs/go/main/App";
import StartButton from "../components/StartButton";
import Counter from "../components/Counter";
import Container from "../components/Container";
import Column from "../components/Column";

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
    <Container>
      <Title buttonTitle="Give Feedback" />

      {/* the two columns in a row */}
      <div
        style={{
          display: "flex",
          flexFlow: "row",
          columnGap: "20px",
        }}
      >
        {/* left side */}
        <div
          style={{
            display: "flex",
            flexFlow: "column",
            flexGrow: "1",
            rowGap: "20px",
            flex: "2",
          }}
        >
          <Column title="Blocktime" buttonTitle="Reset">
            <Counter text="Days" />
            <Counter text="Hours" />
            <Counter text="Minutes" />
            <Counter text="Seconds" />
          </Column>

          <StartButton text="Start Blocker" />
        </div>

        {/* right side */}
        <Column title="Blocklist" buttonTitle="Add">
          {/* The list */}
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
    </Container>
  );
};
