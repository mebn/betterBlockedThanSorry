export default ({ text }) => {
  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        justifyContent: "space-between",
        // gap: "10px",
        borderRadius: "10px",
        // padding: "5px 20px",
      }}
    >
      <p>{text}</p>
      <button
        style={{
          background: "#EFEFEF",
          color: "#4E67D6",
          padding: "10px 20px",
          border: "none",
          borderRadius: "20px",
          cursor: "pointer",
          fontWeight: "bold",
        }}
      >
        &#x1F5D1;
      </button>
    </div>
  );
};
