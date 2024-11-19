export default ({ textEnabled, textDisabled, onClick, disabled }) => {
  const buttonStyle = {
    padding: "20px",
    border: "none",
    cursor: disabled ? "not-allowed" : "pointer",
    borderRadius: "10px",
    background: disabled ? "#7E7E7E" : "#4E67D6",
    color: "white",
    fontWeight: "bold",
    width: "100%",
  };

  const textStyle = {
    cursor: disabled ? "not-allowed" : "pointer",
  };

  return (
    <button style={buttonStyle} onClick={onClick}>
      <h2 style={textStyle}>{disabled ? textDisabled : textEnabled}</h2>
    </button>
  );
};
