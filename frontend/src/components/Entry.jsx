import RoundButton from "./RoundButton";

export default ({
  title,
  buttonText,
  onClick,
  disabled,
  hidden,
  monochrome,
}) => {
  const containerStyle = {
    display: "flex",
    alignItems: "center",
    justifyContent: "space-between",
    borderRadius: "10px",
  };

  return (
    <div style={containerStyle}>
      <p>{title}</p>

      <RoundButton
        text={buttonText}
        onClick={onClick}
        disabled={disabled}
        hidden={hidden}
        monochrome={monochrome}
      />
    </div>
  );
};
