import { Box, Dialog, Typography } from "@mui/material";
import { Schedule } from "../types/types";

type ScheduleModalProps = {
  schedule: Schedule;
  open: boolean;
  onClose: () => void;
};

const ScheduleModal = (props: ScheduleModalProps): JSX.Element => {
  const { schedule, open, onClose } = props;

  return (
    <Dialog open={open} onClose={onClose} fullWidth maxWidth="md">
      <Box p="10%">
        <Typography variant="h5" component="h1">
          {schedule.title}
        </Typography>
        <Typography>
          {/*  TODO: この辺の表記を違和感がないように修正 */}
          {`${schedule.start.getMonth()}月${schedule.start.getDate()}日・`}
          {schedule.isAllDay
            ? "終日"
            : `${schedule.start.toLocaleTimeString([], {
                hour: "2-digit",
                minute: "2-digit",
              })}
           ~ 
          ${schedule.end.toLocaleTimeString([], {
            hour: "2-digit",
            minute: "2-digit",
          })}`}
        </Typography>
        <Box pt="10%">ここにdiaryを載せる</Box>
      </Box>
    </Dialog>
  );
};

export default ScheduleModal;
