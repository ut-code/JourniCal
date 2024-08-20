import { Journal } from "../types/types";

export default function JournalEntry({ journal }: { journal: Journal }) {
  const parsedDate = new Date(journal.date);
  return (
    <div className="journal-entry" style={{ display: "flex" }}>
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
          {parsedDate.getDate()}
        </h3>
        <div>
          {parsedDate.getDay() == 0
            ? "Sun"
            : parsedDate.getDay() == 1
              ? "Mon"
              : parsedDate.getDay() == 2
                ? "Tue"
                : parsedDate.getDay() == 3
                  ? "Wed"
                  : parsedDate.getDay() == 4
                    ? "Thu"
                    : parsedDate.getDay() == 5
                      ? "Fri"
                      : "Sat"}
        </div>
      </div>
      <div></div>
      <div
        className="content"
        style={{ flexBasis: "85%", marginRight: "10px" }}
      >
        <h3>{journal.title}</h3>
        <div>{journal.content}</div>
      </div>
    </div>
  );
}
