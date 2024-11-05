import { useState } from "react";
import "./App.css";
import { StartBlocker } from "../wailsjs/go/main/App";

function App() {
  const [blocklist, setBlocklist] = useState([
    "svt.se",
    "reddit.com",
    "youtube.com",
  ]);

  // in secs
  const [blocktime, setBLocktime] = useState(60);

  const updateBlocklist = (e) => {
    setBlocklist(e.target.value.split(","));
  };

  function startBlocker() {
    StartBlocker(blocktime, blocklist).then((res) => {
      console.log(res);
    });
  }

  return (
    <div id="App">
      <div>
        <input
          id="blocklist"
          onChange={updateBlocklist}
          value={blocklist}
          autoComplete="off"
          name="blocklist"
          type="text"
        />

        <br />

        <input
          id="blocktime"
          onChange={(e) => setBLocktime(e.target.value)}
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

export default App;
