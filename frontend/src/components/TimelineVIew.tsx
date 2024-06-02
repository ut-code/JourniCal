import {
  Box,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";
import TimelineSchedule, { schedule } from "./TimelineSchedule";

type TimelineViewProps = {
  day: Date;
  today: Date;
  daySchedules: schedule[];
};

const TimelineView = (props: TimelineViewProps): JSX.Element => {
  const DAYOFWEEK = ["日", "月", "火", "水", "木", "金", "土"];
  const isToday =
    props.day.getFullYear() === props.today.getFullYear() &&
    props.day.getMonth() === props.today.getMonth() &&
    props.day.getDate() === props.today.getDate();

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
          <Typography>{DAYOFWEEK[props.day.getDay()]}</Typography>
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
                {props.day.getDate()}
              </Typography>
            </Box>
          ) : (
            <Typography variant="h5">{props.day.getDate()}</Typography>
          )}
        </Box>
      </TableHead>
      <TableBody sx={{ position: "relative" }}>
        {[...Array(25).keys()].map((i) => (
          <TableRow key={i}>
            <TableCell></TableCell>
          </TableRow>
        ))}
        {props.daySchedules.map((schedule) => (
          <TimelineSchedule schedule={schedule} />
        ))}
      </TableBody>
    </Table>
  );
};
export default TimelineView;