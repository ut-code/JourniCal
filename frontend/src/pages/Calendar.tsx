import React, { useEffect, useState } from "react";
import TimelineView from "../components/TimelineView";
import { add, sub } from "date-fns";
import { Box, Button, MenuItem, Select } from "@mui/material";
import TopBar from "../components/TopBar";
import TimelineRowName from "../components/TimelineRowName";
import ScheduleView from "../components/ScheduleView";

type ModeVariant = "schedule" | "day" | "3days" | "week";

const Calendar: React.FC = () => {
  const today = new Date(new Date().toDateString());
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
          <ScheduleView />
        ) : (
          <Box display={"flex"}>
            <TimelineRowName />
            {mode === "day" ? (
              <TimelineView
                key={baseDate.getTime()}
                day={baseDate}
                today={today}
              />
            ) : mode === "3days" ? (
              threeDays.map((day) => (
                <TimelineView key={day.getTime()} day={day} today={today} />
              ))
            ) : (
              week.map((day) => (
                <TimelineView key={day.getTime()} day={day} today={today} />
              ))
            )}
          </Box>
        )}
      </Box>
    </>
  );
};

export default Calendar;
