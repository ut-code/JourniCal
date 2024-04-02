export default function DiaryEntry({
  date,
  title,
  content,
}: {
  date: Date;
  title: string;
  content: string;
}) {
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
          {date.getDate()}
        </h3>
        <div>
          {date.getDay() == 0
            ? "Sun"
            : date.getDay() == 1
              ? "Mon"
              : date.getDay() == 2
                ? "Tue"
                : date.getDay() == 3
                  ? "Wed"
                  : date.getDay() == 4
                    ? "Thu"
                    : date.getDay() == 5
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
