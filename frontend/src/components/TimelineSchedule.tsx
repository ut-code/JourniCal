import { Box, Typography } from "@mui/material";
import { Duration, intervalToDuration } from "date-fns";

const durationToHours = (duration: Duration) => {
  const hours = duration.hours ? duration.hours : 0;
  const minutes = duration.minutes ? duration.minutes : 0;
  return hours + minutes / 60;
};

export type schedule = {
  isAllDay: boolean;
  start: Date;
  end: Date;
  title: string;
  color: string;
};

type TimelineScheduleProps = {
  schedule: schedule;
};

const TimelineSchedule = (props: TimelineScheduleProps): JSX.Element => {
  // TODO 変数名の付け方を統一する
  const schedule = props.schedule;

  const hoursBeforeStart =
    schedule.start.getHours() + schedule.start.getMinutes() / 60;

  const scheduleDuration = intervalToDuration({
    start: schedule.start,
    end: schedule.end,
  });

  // const modifiedTitle =
  //   schedule.title.length > 7
  //     ? schedule.title.substring(0, 7)
  //     : schedule.title;

  const scheduleDurationHours = durationToHours(scheduleDuration);
  //TODO 期間が短いと文字がはみ出るのをなんとかする
  return (
    <Box
      sx={{
        position: "absolute",
        paddingLeft: "10px",
        top: `${4 * (hoursBeforeStart + 1)}%`,
        width: "90%",
        height: `${4 * scheduleDurationHours}%`,
        minHeight: "25px",
        borderRadius: "5px",
        backgroundColor: schedule.color,
      }}
    >
      <Typography variant="caption">{schedule.title}</Typography>{" "}
      <Typography variant="caption">
        {schedule.start.toLocaleTimeString([], {
          hour: "2-digit",
          minute: "2-digit",
        })}
        {" ~ "}
        {schedule.end.toLocaleTimeString([], {
          hour: "2-digit",
          minute: "2-digit",
        })}
      </Typography>
    </Box>
  );
};

export default TimelineSchedule;
