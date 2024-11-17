export default ({ title, buttonTitle, isRunning, onClick, children }) => {
  return (
    <div
      style={{
        display: "flex",
        flexFlow: "column",
        gap: "10px",
        background: "#FEFEFE",
        borderRadius: "10px",
        padding: "20px",
        overflowY: "hidden",
        height: "100%",
        boxSizing: "border-box",
        overflow: "hidden",
      }}
    >
      {/* info and button */}
      <div
        style={{
          display: "flex",
          flexFlow: "row",
          justifyContent: "space-between",
          alignItems: "center",
        }}
      >
        <h3 style={{ color: "#7E7E7E" }}>{title}</h3>
        <button
          style={{
            visibility:
              buttonTitle == "Reset" && isRunning ? "hidden" : "visible",
            background: "#4E67D6",
            color: "#EFEFEF",
            padding: "10px 20px",
            border: "none",
            borderRadius: "20px",
            cursor: "pointer",
            fontWeight: "bold",
          }}
          onClick={onClick}
        >
          {buttonTitle}
        </button>
      </div>

      {children}
    </div>
  );
};
