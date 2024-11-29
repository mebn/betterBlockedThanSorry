import RoundButton from "./RoundButton";
import logo from "../assets/images/icon.png";

export default ({ buttonText }) => {
  const containerStyle = {
    color: "#4E67D6",
    lineHeight: "30px",
    fontWeight: "bold",
    display: "flex",
    flexFlow: "row",
    justifyContent: "space-between",
    alignItems: "center",
  };

  return (
    <div style={containerStyle}>
      <div
        style={{
          display: "flex",
          flexFlow: "row",
          alignItems: "center",
          gap: "20px",
        }}
      >
        <img
          src={logo}
          alt="Logo"
          width={50}
          height={50}
          draggable={false}
          onContextMenu={() => false}
        />
        <h1>
          Better Blocked
          <br />
          Than Sorry (old)
        </h1>
      </div>

      {buttonText && <RoundButton text={buttonText} />}
    </div>
  );
};
