export default function DiaryEntry({
  date,
  title,
  content,
}: {
  date: string;
  title: string;
  content: string;
}) {
  const parsedDate = new Date(date);
  return (
    <div className="diary-entry" style={{ display: "flex" }}>
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
        <h3>{title}</h3>
        <div>{content}</div>
      </div>
    </div>
  );
}
