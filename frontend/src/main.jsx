import React from "react";
import { createRoot } from "react-dom/client";
import "./style.css";
import Home from "./screens/Home";

const container = document.getElementById("root");
const root = createRoot(container);

root.render(
  <React.StrictMode>
    <Home />
  </React.StrictMode>
);
