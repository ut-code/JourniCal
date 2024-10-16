import { TableCell, TableRow, Typography } from "@mui/material";
import { Duration, intervalToDuration } from "date-fns";
import { Schedule } from "../types/types";
import { useState } from "react";
import JournalEditDialog from "./JournalEditDialog";
import ScheduleModal from "./ScheduleModal";

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
  const [isScheduleModalOpen, setIsScheduleModalOpen] = useState(false);

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

  const [open, setOpen] = useState(false);
  const [currentSchedule, setCurrentSchedule] = useState<Schedule | null>(null);

  return (
    <>
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
          setCurrentSchedule(schedule);
          setOpen(true);
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
      {currentSchedule && (
        <JournalEditDialog
          open={open}
          handleClose={() => {
            setOpen(false);
          }}
          schedule={currentSchedule}
        />
      )}
      {schedule.isAllDay ? (
        <TableRow role="button" onClick={() => setIsScheduleModalOpen(true)}>
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
      ) : (
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
          role="button"
          onClick={() => setIsScheduleModalOpen(true)}
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
      )}
      <ScheduleModal
        schedule={schedule}
        open={isScheduleModalOpen}
        onClose={() => setIsScheduleModalOpen(false)}
      />
    </>
  );
};

export default TimelineSchedule;
