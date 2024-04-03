import React, { useEffect, useState } from "react";
import TimelineView from "../components/TimelineVIew";
import { add, sub } from "date-fns";
import { Box, Button } from "@mui/material";
import TopBar from "../components/TopBar";
import TimelineRowName from "../components/TimelineRowName";
import { schedule } from "../components/TimelineSchedule";

const isEqualDay = (day1: Date, day2: Date) => {
  return (
    day1.getFullYear() === day2.getFullYear() &&
    day1.getMonth() === day2.getMonth() &&
    day1.getDate() === day2.getDate()
  );
};

const WeekViewTest: React.FC = () => {
  const today = new Date();
  const [baseDate, setBaseDate] = useState(
    sub(today, { days: today.getDay() }),
  );
  const [week, setWeek] = useState(
    [...Array(7).keys()].map((i) => add(baseDate, { days: i })),
  );
  useEffect(() => {
    setWeek([...Array(7).keys()].map((i) => add(baseDate, { days: i })));
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
  ];

  return (
    <Box>
      <TopBar journalPathName="/page1" calendarPathName="/page2" />
      <Button onClick={() => setBaseDate(sub(baseDate, { days: 7 }))}>
        前の週
      </Button>
      <Button onClick={() => setBaseDate(add(baseDate, { days: 7 }))}>
        次の週
      </Button>
      <Box display={"flex"}>
        <TimelineRowName />
        {week.map((day) => (
          <TimelineView
            key={day.getTime()}
            day={day}
            today={today}
            daySchedules={weekSchedules.filter((schedule) =>
              isEqualDay(day, schedule.start),
            )}
          />
        ))}
      </Box>
    </Box>
  );
};

export default WeekViewTest;
