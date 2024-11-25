import { useState, useRef, useEffect } from "react";

export default ({ isOpen, onClose, onAddWebsite }) => {
  const [website, setWebsite] = useState("");
  const dialogRef = useRef(null);

  useEffect(() => {
    if (isOpen) {
      dialogRef.current.showModal();
    } else if (dialogRef.current.open) {
      dialogRef.current.close();
    }
  }, [isOpen]);

  const handleAdd = () => {
    if (website.trim()) {
      onAddWebsite(website.trim());
      setWebsite("");
      onClose();
    }
  };

  const dialogStyle = {
    position: "fixed",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    padding: "20px",
    border: "none",
    borderRadius: "10px",
    background: "white",
    maxWidth: "400px",
  };

  const Button = ({ text, onClick, background }) => (
    <button
      style={{
        padding: "15px 20px",
        border: "none",
        cursor: "pointer",
        borderRadius: "10px",
        background,
        color: "white",
        fontWeight: "bold",
      }}
      onClick={onClick}
    >
      {text}
    </button>
  );

  return (
    <dialog
      ref={dialogRef}
      style={dialogStyle}
      onClick={(e) => {
        if (e.currentTarget === e.target) {
          setWebsite("");
          onClose();
        }
      }}
    >
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "20px",
        }}
      >
        <h3
          style={{
            color: "#7E7E7E",
          }}
        >
          Add new website to blocklist
        </h3>

        {/* <p style={{}}>
          You don't need "http://", "https://", or "www".
          <br />
          You
        </p> */}

        <input
          style={{
            background: "#EFEFEF",
            border: "none",
            borderRadius: "10px",
            padding: "10px",
            fontSize: "16px",
            outline: "none",
            cursor: "auto",
          }}
          type="text"
          placeholder="www.website.com"
          value={website}
          onChange={(e) => setWebsite(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === "Enter") {
              handleAdd();
            }
          }}
        />

        <div
          style={{
            display: "flex",
            justifyContent: "space-between",
          }}
        >
          <Button
            text="Cancel"
            onClick={() => {
              setWebsite("");
              onClose();
            }}
            background="#FF6B6B"
          />

          <Button text="Add" onClick={handleAdd} background="#4E67D6" />
        </div>
      </div>
    </dialog>
  );
};
