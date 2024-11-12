import {
  Box,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";
import TimelineSchedule from "./TimelineSchedule";
import { add } from "date-fns";
import { Schedule } from "../types/types";
import { useEffect, useState } from "react";
import useSWR from "swr";
import { FetchedSchedule, scheduleFromFetchedData } from "../utils/calendar";

type TimelineViewProps = {
  day: Date;
  today: Date;
};

const isEqualDay = (day1: Date, day2: Date) => {
  return (
    day1.getFullYear() === day2.getFullYear() &&
    day1.getMonth() === day2.getMonth() &&
    day1.getDate() === day2.getDate()
  );
};

const TimelineView = (props: TimelineViewProps): JSX.Element => {
  const { day, today } = props;
  const DAYOFWEEK = ["日", "月", "火", "水", "木", "金", "土"];
  const isToday = isEqualDay(day, today);
  const [schedules, setSchedules] = useState<Schedule[]>([]);

  // データフェッチ
  const startUnixTime = Math.floor(day.getTime() / 1000);
  const endUnixTime = Math.floor(add(day, { days: 1 }).getTime() / 1000);
  const { data, error } = useSWR(
    `http://localhost:3000/api/calendar/get-events-in-range/${startUnixTime}/${endUnixTime}`,
    (url) =>
      fetch(url, {
        method: "GET",
        credentials: "include",
        mode: "cors",
      }).then((r) => r.json()),
  );
  if (error) {
    console.error(error);
  }
  useEffect(() => {
    if (data) {
      setSchedules(
        data.map((schedule: FetchedSchedule) =>
          scheduleFromFetchedData(schedule),
        ),
      );
    }
  }, [data]);

  return (
    <Table
      sx={{
        height: "90vh",
        minHeight: "1150px",
        borderRight: "1px solid gainsboro",
      }}
    >
      <TableHead>
        <Box
          display={"flex"}
          flexDirection={"column"}
          justifyContent={"center"}
          alignItems={"center"}
          height={"3vh"}
        >
          <Typography>{DAYOFWEEK[day.getDay()]}</Typography>
          {isToday ? (
            <Box
              display={"flex"}
              justifyContent={"center"}
              sx={{
                backgroundColor: "primary.main",
                borderRadius: "50px",
                width: "2rem",
                height: "2rem",
              }}
            >
              <Typography variant="h5" color={"primary.contrastText"}>
                {day.getDate()}
              </Typography>
            </Box>
          ) : (
            <Typography variant="h5">{day.getDate()}</Typography>
          )}
        </Box>
        {schedules
          .filter((schedule) => schedule.isAllDay)
          .map((schedule) => (
            <TimelineSchedule key={schedule.id} schedule={schedule} />
          ))}
      </TableHead>
      <TableBody sx={{ position: "relative" }}>
        {[...Array(25).keys()].map((i) => (
          <TableRow key={i}>
            <TableCell></TableCell>
          </TableRow>
        ))}
        {schedules
          .filter((schedule) => !schedule.isAllDay)
          .map((schedule) => (
            <TimelineSchedule key={schedule.id} schedule={schedule} />
          ))}
      </TableBody>
    </Table>
  );
};
export default TimelineView;
