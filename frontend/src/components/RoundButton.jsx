export default ({ text, onClick, disabled, hidden, monochrome }) => {
  const monochromeBackground = monochrome ? "#EFEFEF" : "#4E67D6";

  const buttonStyle = {
    visibility: hidden ? "hidden" : "visible",
    color: monochrome ? "black" : "#EFEFEF",
    background: disabled ? "#7E7E7E" : monochromeBackground,
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
