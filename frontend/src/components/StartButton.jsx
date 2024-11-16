export default ({ text }) => {
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
      }}
    >
      <h2>{text}</h2>
    </button>
  );
};
