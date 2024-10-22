import { Box, Button, Dialog, Typography } from "@mui/material";
import AddIcon from "@mui/icons-material/Add";
import { Journal, Schedule } from "../types/types";
import useSWR from "swr";
import { useEffect, useState } from "react";
import JournalEditDialog from "./JournalEditDialog";

type JournalPresenterProps = {
  journal: Journal;
};

const JournalPresenter = ({ journal }: JournalPresenterProps): JSX.Element => {
  return (
    <>
      <Typography variant="h6" component="h2">
        {journal.title}
      </Typography>
      <Typography>{journal.content}</Typography>
    </>
  );
};

type ScheduleModalProps = {
  schedule: Schedule;
  open: boolean;
  onClose: () => void;
};

const ScheduleModal = (props: ScheduleModalProps): JSX.Element => {
  const { schedule, open, onClose } = props;
  const [journal, setJournal] = useState<Journal | undefined>(undefined);
  const [journalEditDialogOpen, setJournalEditDialogOpen] = useState(false);

  // diaryを取得
  const { data, error } = useSWR(
    `http://localhost:3000/api/journals/event/${schedule.id}`,
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
      setJournal(data);
    }
  }, [data]);

  return (
    <>
      <Dialog open={open} onClose={onClose} fullWidth maxWidth="md">
        <Box p="10%">
          <Typography variant="h5" component="h1">
            {schedule.title}
          </Typography>
          <Typography>
            {/*  TODO: この辺の表記を違和感がないように修正 */}
            {`${schedule.start.getMonth() + 1}月${schedule.start.getDate()}日・`}
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
          <Box pt="10%">
            {journal ? (
              <JournalPresenter journal={journal} />
            ) : (
              <Box>
                <Typography>ジャーナルがありません</Typography>
                <Button
                  startIcon={<AddIcon />}
                  onClick={() => {
                    setJournalEditDialogOpen(true);
                  }}
                >
                  追加する
                </Button>
              </Box>
            )}
          </Box>
        </Box>
      </Dialog>
      {schedule && (
        <JournalEditDialog
          open={journalEditDialogOpen}
          handleClose={() => {
            setJournalEditDialogOpen(false);
          }}
          schedule={schedule}
        />
      )}
    </>
  );
};

export default ScheduleModal;
