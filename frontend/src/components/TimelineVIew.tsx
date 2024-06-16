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
import { Schedule } from "../types/types";

type TimelineViewProps = {
  day: Date;
  today: Date;
  daySchedules: Schedule[];
};

const TimelineView = (props: TimelineViewProps): JSX.Element => {
  const { day, today, daySchedules } = props;
  const DAYOFWEEK = ["日", "月", "火", "水", "木", "金", "土"];
  const isToday =
    day.getFullYear() === today.getFullYear() &&
    day.getMonth() === today.getMonth() &&
    day.getDate() === today.getDate();

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
        <Box mt="10px">
          {daySchedules
            .filter((schedule) => schedule.isAllDay)
            .map((schedule) => (
              <Box
                key={schedule.id}
                sx={{
                  paddingLeft: "10px",
                  width: "90%",
                  borderRadius: "5px",
                  backgroundColor: schedule.color,
                }}
              >
                {schedule.title}
              </Box>
            ))}
        </Box>
      </TableHead>
      <TableBody sx={{ position: "relative" }}>
        {[...Array(25).keys()].map((i) => (
          <TableRow key={i}>
            <TableCell></TableCell>
          </TableRow>
        ))}
        {daySchedules
          .filter((schedule) => !schedule.isAllDay)
          .map((schedule) => (
            <TimelineSchedule key={schedule.id} schedule={schedule} />
          ))}
      </TableBody>
    </Table>
  );
};
export default TimelineView;
