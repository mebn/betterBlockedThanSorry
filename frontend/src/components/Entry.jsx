export default ({ text, isRunning }) => {
  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        justifyContent: "space-between",
        borderRadius: "10px",
      }}
    >
      <p>{text}</p>
      <button
        style={{
          visibility: isRunning ? "hidden" : "visible",
          background: "#EFEFEF",
          color: "#4E67D6",
          padding: "10px 20px",
          border: "none",
          borderRadius: "20px",
          cursor: "pointer",
          fontWeight: "bold",
        }}
        disabled={isRunning}
      >
        &#x1F5D1;
      </button>
    </div>
  );
};
