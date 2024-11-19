import RoundButton from "./RoundButton";

export default ({ buttonText }) => {
  const containerStyle = {
    color: "#4E67D6",
    lineHeight: "30px",
    fontWeight: "bold",
    display: "flex",
    flexFlow: "row",
    justifyContent: "space-between",
    alignItems: "start",
  };

  return (
    <div style={containerStyle}>
      <h1>
        Better Blocked
        <br />
        Than Sorry
      </h1>

      {buttonText ? <RoundButton text={buttonText} /> : <></>}
    </div>
  );
};
