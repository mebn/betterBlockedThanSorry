export default ({ text, isRunning, onClick }) => {
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
          width: "30px",
          height: "30px",
          // marginRight: "10px",
          border: "none",
          borderRadius: "20px",
          cursor: "pointer",
          fontWeight: "bold",
        }}
        disabled={isRunning}
        onClick={onClick}
      >
        &#x1F5D1;
      </button>
    </div>
  );
};
