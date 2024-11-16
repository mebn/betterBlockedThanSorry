export default ({ buttonTitle }) => {
  return (
    <div
      style={{
        color: "#4E67D6",
        lineHeight: "30px",
        fontWeight: "bold",
        display: "flex",
        flexFlow: "row",
        justifyContent: "space-between",
        alignItems: "start",
      }}
    >
      <h1>
        Better Blocked
        <br />
        Than Sorry
      </h1>

      <button
        style={{
          color: "#EFEFEF",
          background: "#4E67D6",
          padding: "10px 20px",
          border: "none",
          borderRadius: "20px",
          cursor: "pointer",
          fontWeight: "bold",
        }}
      >
        {buttonTitle}
      </button>
    </div>
  );
};
