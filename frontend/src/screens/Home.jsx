import Title from "../components/Title";

import {
  StartBlocker,
  GetDaemonRunningStatus,
} from "../../wailsjs/go/main/App";

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
    <div>
      <Title />

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
};
