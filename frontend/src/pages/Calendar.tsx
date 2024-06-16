import React, { useEffect, useState } from "react";
import TimelineView from "../components/TimelineVIew";
import { add, sub } from "date-fns";
import { Box, Button, MenuItem, Select } from "@mui/material";
import TopBar from "../components/TopBar";
import TimelineRowName from "../components/TimelineRowName";
import { Schedule } from "../types/types";

type ModeVariant = "schedule" | "day" | "3days" | "week";
type FetchedSchedule = {
  id: string;
  colorId?: string;
  start: {
    date?: string;
    dateTime?: string;
  };
  end: {
    date?: string;
    dateTime?: string;
  };
  summary: string;
};

const COLOR_DICT = [
  "#7986CB",
  "#33B679",
  "#8E24AA",
  "#E67C73",
  "#F6BF26",
  "#F4511E",
  "#039BE5",
  "#616161",
  "#3F51B5",
  "#0B8043",
  "#D50000",
];

const scheduleFromFetchedData = (
  fetchedSchedule: FetchedSchedule,
): Schedule => {
  if (
    fetchedSchedule.start.dateTime == undefined &&
    fetchedSchedule.start.date != undefined &&
    fetchedSchedule.end.dateTime == undefined &&
    fetchedSchedule.end.date != undefined
  ) {
    return {
      id: fetchedSchedule.id,
      isAllDay: true,
      start: new Date(fetchedSchedule.start.date),
      end: new Date(fetchedSchedule.end.date),
      title: fetchedSchedule.summary,
      color:
        COLOR_DICT[
          Number(fetchedSchedule.colorId)
            ? Number(fetchedSchedule.colorId) - 1
            : 6
        ],
    };
  }
  if (
    fetchedSchedule.start.dateTime != undefined &&
    fetchedSchedule.start.date == undefined &&
    fetchedSchedule.end.dateTime != undefined &&
    fetchedSchedule.end.date == undefined
  ) {
    return {
      id: fetchedSchedule.id,
      isAllDay: false,
      start: new Date(fetchedSchedule.start.dateTime),
      end: new Date(fetchedSchedule.end.dateTime),
      title: fetchedSchedule.summary,
      color:
        COLOR_DICT[
          Number(fetchedSchedule.colorId)
            ? Number(fetchedSchedule.colorId) - 1
            : 6
        ],
    };
  }
  throw new Error("invalid schedule format.");
};

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
  const [mode, setMode] = useState<ModeVariant>("day");

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

  const [weekSchedules, setWeekSchedules] = useState<Schedule[]>([]);

  // データフェッチ
  useEffect(() => {
    async function fetchData() {
      const response = await fetch(
        `http://localhost:3000/api/calendar/get-20-events-forward/1717250000`,
        {
          method: "GET",
          credentials: "include",
          mode: "cors",
        },
      );
      const data = await response.json();
      console.log(data);
      // 一週間の予定を格納
      setWeekSchedules(
        data.map((schedule: FetchedSchedule) =>
          scheduleFromFetchedData(schedule),
        ),
      );
    }
    fetchData();
  }, []);

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
          onChange={(e) => setMode(e.target.value as ModeVariant)}
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
