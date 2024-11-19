import RoundButton from "./RoundButton";

export default ({ title, buttonText, disabled, hidden, onClick, children }) => {
  const containerStyle = {
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
  };

  const headerStyle = {
    display: "flex",
    flexFlow: "row",
    justifyContent: "space-between",
    alignItems: "center",
  };

  const titleStyle = { color: "#7E7E7E" };

  return (
    <div style={containerStyle}>
      <div style={headerStyle}>
        <h3 style={titleStyle}>{title}</h3>
        {buttonText && (
          <RoundButton
            text={buttonText}
            onClick={onClick}
            disabled={disabled}
            hidden={hidden}
          />
        )}
      </div>

      {children}
    </div>
  );
};
