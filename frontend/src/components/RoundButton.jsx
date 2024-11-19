export default ({ text }) => {
  const buttonStyle = {
    color: "#EFEFEF",
    background: "#4E67D6",
    padding: "10px 20px",
    border: "none",
    borderRadius: "20px",
    cursor: "pointer",
    fontWeight: "bold",
  };

  return <button style={buttonStyle}>{text}</button>;
};
