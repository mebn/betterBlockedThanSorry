export default ({ text, onClick, disabled, hidden }) => {
  const buttonStyle = {
    visibility: hidden ? "hidden" : "visible",
    color: "#EFEFEF",
    background: disabled ? "#7E7E7E" : "#4E67D6",
    cursor: disabled ? "not-allowed" : "pointer",
    padding: "10px 20px",
    border: "none",
    borderRadius: "20px",
    fontWeight: "bold",
  };

  return (
    <button style={buttonStyle} onClick={onClick} disabled={disabled}>
      {text}
    </button>
  );
};
