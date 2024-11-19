export default ({
  title,
  value,
  disabled,
  blocktime,
  blocktimeEntry,
  setBlocktime,
  maxVal,
}) => {
  const handleOnChange = (e) => {
    let val = e.target.value;
    const number = parseInt(val, 10);

    if (val == "") {
      setBlocktime({
        ...blocktime,
        [blocktimeEntry]: "",
      });
    } else if (!isNaN(number) && number >= 0) {
      setBlocktime({
        ...blocktime,
        [blocktimeEntry]: Math.min(number, maxVal),
      });
    }
  };

  const containerStyle = {
    display: "flex",
    alignItems: "center",
    justifyContent: "left",
    gap: "20px",
  };

  const inputStyle = {
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
  };

  return (
    <div style={containerStyle}>
      <input
        style={inputStyle}
        type="text"
        placeholder="0"
        value={value}
        disabled={disabled}
        onChange={handleOnChange}
      />

      {title}
    </div>
  );
};
