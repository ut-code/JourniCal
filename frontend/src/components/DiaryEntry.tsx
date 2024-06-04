import React from "react";

function DiaryEntry({ date, title, content, onEdit }: { date: Date; title: string; content: string; onEdit: () => void; }) {
  const parsedContent = content.split("\n").map((line, index) => (
    <React.Fragment key={index}>
      {line}
      <br />
    </React.Fragment>
  ));

  return (
    <div className="diary-entry" style={{ display: "flex" ,alignItems: "baseline"  }}>
      <div
        className="date"
        style={{
          flexBasis: "15%",
          textAlign: "center",
          marginRight: "10px",
          marginLeft: "10px",
        }}
      >
        <h3 style={{ marginBottom: "0%", fontSize: "2em" }}>
          {date.getDate()}
        </h3>
        <div>
          {date.getDay() === 0
            ? "Sun"
            : date.getDay() === 1
              ? "Mon"
              : date.getDay() === 2
                ? "Tue"
                : date.getDay() === 3
                  ? "Wed"
                  : date.getDay() === 4
                    ? "Thu"
                    : date.getDay() === 5
                      ? "Fri"
                      : "Sat"}
        </div>
      </div>
      <div
        className="content"
        style={{ flexBasis: "85%", marginRight: "10px" }}
      >
        <h3>{title}</h3>
        <div>{parsedContent}</div>
      </div>
      <button onClick={onEdit} style={{
        marginTop: "1%", 
        marginRight: "1%",
        width: "40px",
        height: "40px",
        borderRadius: "50%",
        backgroundColor: "glay",
        border: "none",
        outline: "none",
        cursor: "pointer"
      }}>✏️</button>
    </div>
  );
}

export default DiaryEntry;

