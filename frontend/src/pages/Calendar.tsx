import React, { useEffect, useState } from "react";
import TimelineView from "../components/TimelineVIew";
import { add, sub } from "date-fns";
import { Box, Button, MenuItem, Select } from "@mui/material";
import TopBar from "../components/TopBar";
import TimelineRowName from "../components/TimelineRowName";
import { schedule } from "../components/TimelineSchedule";

type modeVariant = "schedule" | "day" | "3days" | "week";

const isEqualDay = (day1: Date, day2: Date) => {
  return (
    day1.getFullYear() === day2.getFullYear() &&
    day1.getMonth() === day2.getMonth() &&
    day1.getDate() === day2.getDate()
  );
};

const Calendar: React.FC = () => {
  const today = new Date();
  const [baseDate, setBaseDate] = useState(
    sub(today, { days: today.getDay() }),
  );

  const [mode, setMode] = useState<modeVariant>("day");

  const [threeDays, setThreeDays] = useState(
    [...Array(3).keys()].map((i) => add(baseDate, { days: i })),
  );
  const [week, setWeek] = useState(
    [...Array(7).keys()].map((i) =>
      add(sub(baseDate, { days: baseDate.getDay() }), { days: i }),
    ),
  );
  useEffect(() => {
    setWeek(
      [...Array(7).keys()].map((i) =>
        add(sub(baseDate, { days: baseDate.getDay() }), { days: i }),
      ),
    );
    setThreeDays([...Array(3).keys()].map((i) => add(baseDate, { days: i })));
  }, [baseDate]);

  // 一週間の予定を格納
  const weekSchedules: schedule[] = [
    {
      title: "工学部ガイダンス",
      start: new Date("2024-04-03T10:30"),
      end: new Date("2024-04-03T17:00"),
      color: "mediumpurple",
    },
    {
      title: "サーオリ手伝い",
      start: new Date("2024-04-04T14:00"),
      end: new Date("2024-04-04T18:00"),
      color: "mediumseagreen",
    },
    {
      title: "2限",
      start: new Date("2024-04-05T10:25"),
      end: new Date("2024-04-05T12:10"),
      color: "mediumseagreen",
    },
    {
      title: "3限",
      start: new Date("2024-04-05T13:00"),
      end: new Date("2024-04-05T14:45"),
      color: "mediumseagreen",
    },
    {
      title: "4限",
      start: new Date("2024-04-05T14:55"),
      end: new Date("2024-04-05T16:40"),
      color: "mediumseagreen",
    },
    {
      title: "5限",
      start: new Date("2024-04-05T16:50"),
      end: new Date("2024-04-05T18:35"),
      color: "mediumseagreen",
    },
    {
      title: "mayFes mtg",
      start: new Date("2024-04-05T21:00"),
      end: new Date("2024-04-05T22:00"),
      color: "dodgerblue",
    },
    {
      title: "1限",
      start: new Date("2024-04-08T08:30"),
      end: new Date("2024-04-08T10:15"),
      color: "mediumseagreen",
    },
    {
      title: "2限",
      start: new Date("2024-04-08T10:25"),
      end: new Date("2024-04-08T12:10"),
      color: "mediumseagreen",
    },
    {
      title: "3限",
      start: new Date("2024-04-08T13:00"),
      end: new Date("2024-04-08T14:45"),
      color: "mediumseagreen",
    },
    {
      title: "4限",
      start: new Date("2024-04-08T14:55"),
      end: new Date("2024-04-08T16:40"),
      color: "mediumseagreen",
    },
    {
      title: "5限",
      start: new Date("2024-04-08T16:50"),
      end: new Date("2024-04-08T18:35"),
      color: "mediumseagreen",
    },
    {
      title: "braille mtg",
      start: new Date("2024-04-08T22:00"),
      end: new Date("2024-04-08T23:00"),
      color: "dodgerblue",
    },
    {
      title: "2限",
      start: new Date("2024-04-09T10:25"),
      end: new Date("2024-04-09T12:10"),
      color: "mediumseagreen",
    },
    {
      title: "3限",
      start: new Date("2024-04-09T13:00"),
      end: new Date("2024-04-09T14:45"),
      color: "mediumseagreen",
    },
    {
      title: "4限",
      start: new Date("2024-04-09T14:55"),
      end: new Date("2024-04-09T16:40"),
      color: "mediumseagreen",
    },
    {
      title: "5限",
      start: new Date("2024-04-09T16:50"),
      end: new Date("2024-04-09T18:35"),
      color: "mediumseagreen",
    },
    {
      title: "journal mtg",
      start: new Date("2024-04-09T21:00"),
      end: new Date("2024-04-09T22:00"),
      color: "dodgerblue",
    },
    {
      title: "長い名前の予定長い名前の予定",
      start: new Date("2024-04-10T21:00"),
      end: new Date("2024-04-10T21:01"),
      color: "dodgerblue",
    },
  ];

  return (
    <>
      <TopBar baseDate={baseDate} setBaseDate={setBaseDate} />
      <Box mt="20%">
        {mode === "day" ? (
          <>
            <Button onClick={() => setBaseDate(sub(baseDate, { days: 1 }))}>
              前の日
            </Button>
            <Button onClick={() => setBaseDate(add(baseDate, { days: 1 }))}>
              次の日
            </Button>
          </>
        ) : mode === "3days" ? (
          <>
            <Button onClick={() => setBaseDate(sub(baseDate, { days: 3 }))}>
              前の3日間
            </Button>
            <Button onClick={() => setBaseDate(add(baseDate, { days: 3 }))}>
              次の3日間
            </Button>
          </>
        ) : (
          mode === "week" && (
            <>
              <Button onClick={() => setBaseDate(sub(baseDate, { days: 7 }))}>
                前の週
              </Button>
              <Button onClick={() => setBaseDate(add(baseDate, { days: 7 }))}>
                次の週
              </Button>
            </>
          )
        )}

        <Select
          value={mode}
          onChange={(e) => setMode(e.target.value as modeVariant)}
        >
          <MenuItem value={"schedule"}>スケジュール</MenuItem>
          <MenuItem value={"day"}>日</MenuItem>
          <MenuItem value={"3days"}>3日間</MenuItem>
          <MenuItem value={"week"}>一週間</MenuItem>
        </Select>

        {mode === "schedule" ? (
          <>ここにスケジュールビューを配置</>
        ) : (
          <Box display={"flex"}>
            <TimelineRowName />
            {mode === "day" ? (
              <TimelineView
                key={baseDate.getTime()}
                day={baseDate}
                today={today}
                daySchedules={weekSchedules.filter((schedule) =>
                  isEqualDay(baseDate, schedule.start),
                )}
              />
            ) : mode === "3days" ? (
              threeDays.map((day) => (
                <TimelineView
                  key={day.getTime()}
                  day={day}
                  today={today}
                  daySchedules={weekSchedules.filter((schedule) =>
                    isEqualDay(day, schedule.start),
                  )}
                />
              ))
            ) : (
              week.map((day) => (
                <TimelineView
                  key={day.getTime()}
                  day={day}
                  today={today}
                  daySchedules={weekSchedules.filter((schedule) =>
                    isEqualDay(day, schedule.start),
                  )}
                />
              ))
            )}
          </Box>
        )}
      </Box>
    </>
  );
};

export default Calendar;
