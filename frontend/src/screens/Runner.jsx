import Title from "../components/Title";

export default ({ endTime, currentTime }) => {
  return (
    <div style={{ color: "white" }}>
      <Title />
      <h2>The blocker will stop in {endTime - currentTime} seconds.</h2>
    </div>
  );
};
