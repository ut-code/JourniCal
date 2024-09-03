import { TableCell, TableRow, Typography } from "@mui/material";
import { Duration, intervalToDuration } from "date-fns";
import { Schedule } from "../types/types";
import useJournal from "../hooks/useJournal";

const durationToHours = (duration: Duration) => {
  const hours = duration.hours ? duration.hours : 0;
  const minutes = duration.minutes ? duration.minutes : 0;
  return hours + minutes / 60;
};

type TimelineScheduleProps = {
  schedule: Schedule;
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

  const { createJournal } = useJournal();

  return (
    <TableRow
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
      onClick={async () => {
        await createJournal({
          title: "tmp",
          content:
            "あのイーハトーヴォのすきとおった風、夏でも底に冷たさをもつ青いそら、うつくしい森で飾られたモリーオ市、郊外のぎらぎらひかる草の波。",
          eventId: schedule.id,
          date: schedule.start.toISOString(),
        });
      }}
    >
      <TableCell padding="none" sx={{ border: "none" }}>
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
      </TableCell>
    </TableRow>
  );
};

export default TimelineSchedule;
