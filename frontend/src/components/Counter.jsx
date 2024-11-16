export default ({ text }) => {
  return (
    <div
      style={{
        display: "flex",
        alignItems: "center",
        justifyContent: "left",
        gap: "20px",
      }}
    >
      <input
        style={{
          width: "50px",
          height: "30px",
          background: "#EFEFEF",
          color: "#4E67D6",
          textAlign: "center",
          border: "none",
          borderRadius: "10px",
          padding: "10px",
          fontSize: "22px",
          outline: "none",
        }}
        type="number"
        name={text}
        id={text}
        placeholder="0"
      />
      <h4>{text}</h4>
    </div>
  );
};
