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

type TimelineViewProps = {
  day: Date;
  today: Date;
};

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
  console.log(data);
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
            <TableRow key={schedule.id}>
              <TableCell
                padding="none"
                sx={{
                  paddingLeft: "10px",
                  border: "none",
                  borderRadius: "5px",
                  backgroundColor: schedule.color,
                }}
              >
                {schedule.title}
              </TableCell>
            </TableRow>
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
