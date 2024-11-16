export default ({ children }) => {
  return (
    <div
      style={{
        margin: "20px",
        display: "flex",
        flexFlow: "column",
        gap: "20px",
      }}
    >
      {children}
    </div>
  );
};
