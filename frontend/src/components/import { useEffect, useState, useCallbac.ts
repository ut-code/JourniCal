import { useEffect, useState, useCallback } from "react";
import InfiniteScroll from "react-infinite-scroller";
import {
  Box,
  Table,
  TableCell,
  TableRow,
  Typography,
} from "@mui/material";
import TimelineSchedule from "./TimelineSchedule";
import { add } from "date-fns";
import { Schedule } from "../types/types";
import useSWR from "swr";

type ScheduleViewProps = {
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

// Refactor to handle both all-day and time-bound schedules
const scheduleFromFetchedData = (fetchedSchedule: FetchedSchedule): Schedule => {
  const color = COLOR_DICT[Number(fetchedSchedule.colorId) - 1] || COLOR_DICT[6];
  if (fetchedSchedule.start.date && fetchedSchedule.end.date) {
    return {
      id: fetchedSchedule.id,
      isAllDay: true,
      start: new Date(fetchedSchedule.start.date),
      end: new Date(fetchedSchedule.end.date),
      title: fetchedSchedule.summary,
      color,
    };
  } else if (fetchedSchedule.start.dateTime && fetchedSchedule.end.dateTime) {
    return {
      id: fetchedSchedule.id,
      isAllDay: false,
      start: new Date(fetchedSchedule.start.dateTime),
      end: new Date(fetchedSchedule.end.dateTime),
      title: fetchedSchedule.summary,
      color,
    };
  } else {
    throw new Error("Invalid schedule format.");
  }
};

// Utility for day comparison
const isEqualDay = (day1: Date, day2: Date) =>
  day1.getFullYear() === day2.getFullYear() &&
  day1.getMonth() === day2.getMonth() &&
  day1.getDate() === day2.getDate();

const ScheduleView = ({ day, today }: ScheduleViewProps): JSX.Element => {
  const DAYOFWEEK = ["日", "月", "火", "水", "木", "金", "土"];
  const isToday = isEqualDay(day, today);
  const [schedules, setSchedules] = useState<Schedule[]>([]);
  const [currentDate, setCurrentDate] = useState<Date>(day);
  const [hasMore, setHasMore] = useState<boolean>(true);

  // check from here
  const fetchSchedules = useCallback(async (date: Date) => {
    const startUnixTime = Math.floor(date.getTime() / 1000);
    const endUnixTime = Math.floor(add(date, { days: 1 }).getTime() / 1000);

    const response = await fetch(
      `http://localhost:3000/api/calendar/get-events-in-range/${startUnixTime}/${endUnixTime}`,
      { method: "GET", credentials: "include", mode: "cors" }
    );
    if (!response.ok) throw new Error("Failed to fetch data.");
    
    const fetchedData: FetchedSchedule[] = await response.json();
    return fetchedData.map(scheduleFromFetchedData);
  }, []);

  useEffect(() => {
    const loadInitialSchedules = async () => {
      try {
        const initialSchedules = await fetchSchedules(currentDate);
        setSchedules(initialSchedules);
      } catch (error) {
        console.error(error);
        setHasMore(false);
      }
    };
    loadInitialSchedules();
  }, [currentDate, fetchSchedules]);

  const loadMore = async () => {
    try {
      const nextDate = add(currentDate, { days: 1 });
      const newSchedules = await fetchSchedules(nextDate);
      
      if (newSchedules.length > 0) {
        setSchedules((prevSchedules) => [...prevSchedules, ...newSchedules]);
        setCurrentDate(nextDate);
      } else {
        setHasMore(false); // No more data to load
      }
    } catch (error) {
      console.error(error);
      setHasMore(false); // Error fetching more schedules
    }
  };

  return (
    <Table
      sx={{
        height: "90vh",
        minHeight: "1150px",
        borderRight: "1px solid gainsboro",
      }}
    >
      <InfiniteScroll
        pageStart={0}
        loadMore={loadMore}
        hasMore={hasMore}
        loader={
          <div className="loader" key={0}>
            Loading...
          </div>
        }
      >
        <Box
          display="flex"
          flexDirection="column"
          justifyContent="center"
          alignItems="center"
          height="3vh"
        >
          <Typography>{DAYOFWEEK[day.getDay()]}</Typography>
          {isToday ? (
            <Box
              display="flex"
              justifyContent="center"
              sx={{
                backgroundColor: "primary.main",
                borderRadius: "50px",
                width: "2rem",
                height: "2rem",
              }}
            >
              <Typography variant="h5" color="primary.contrastText">
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
        {schedules
          .filter((schedule) => !schedule.isAllDay)
          .map((schedule) => (
            <TimelineSchedule key={schedule.id} schedule={schedule} />
          ))}
      </InfiniteScroll>
    </Table>
  );
};

export default ScheduleView;
