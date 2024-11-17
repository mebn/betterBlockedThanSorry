export default ({ text, onClick, isRunning }) => {
  if (isRunning) {
    return (
      <button
        style={{
          padding: "20px",
          border: "none",
          cursor: "not-allowed",
          borderRadius: "10px",
          background: "#7E7E7E",
          color: "white",
          fontWeight: "bold",
          width: "100%",
        }}
      >
        <h2>Blocker started</h2>
      </button>
    );
  }
  return (
    <button
      style={{
        padding: "20px",
        border: "none",
        cursor: "pointer",
        borderRadius: "10px",
        background: "#4E67D6",
        color: "white",
        fontWeight: "bold",
        width: "100%",
      }}
      onClick={onClick}
    >
      <h2 style={{ cursor: "pointer" }}>{text}</h2>
    </button>
  );
};
