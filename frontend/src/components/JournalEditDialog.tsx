import {
  Dialog,
  DialogTitle,
  DialogContent,
  TextField,
  DialogActions,
  Button,
} from "@mui/material";
import { useEffect, useState } from "react";
import { Journal, Schedule } from "../types/types";
import useJournal from "../hooks/useJournal";
import { fetchJournalByEventId } from "../api/journal";

interface Props {
  open: boolean;
  handleClose: () => void;
  schedule: Schedule;
}

export default function JournalEditDialog(props: Props) {
  const { open, handleClose, schedule } = props;

  const { createJournal, updateJournal } = useJournal();

  const [existingJournal, setExistingJournal] = useState<Journal | null>(null);

  const [journalTitle, setJournalTitle] = useState("");
  const [journalContent, setJournalContent] = useState("");

  useEffect(() => {
    (async () => {
      const journalOnEvent = await fetchJournalByEventId(schedule.id);
      setExistingJournal(journalOnEvent);
      if (journalOnEvent) {
        setJournalTitle(journalOnEvent.title);
        setJournalContent(journalOnEvent.content);
      }
    })();
  }, [schedule, open]);

  return (
    <Dialog open={open} onClose={handleClose}>
      <DialogTitle>ジャーナルの{existingJournal ? "編集" : "追加"}</DialogTitle>
      <DialogContent>
        <TextField
          value={journalTitle}
          onChange={(e) => setJournalTitle(e.target.value)}
          autoFocus
          required
          margin="dense"
          id="title"
          name="title"
          label="タイトル"
          type="email"
          fullWidth
          variant="standard"
        />
        <TextField
          value={journalContent}
          onChange={(e) => setJournalContent(e.target.value)}
          autoFocus
          required
          margin="dense"
          id="content"
          name="content"
          label="内容"
          type="text"
          fullWidth
          multiline
          variant="standard"
        />
      </DialogContent>
      <DialogActions>
        <Button
          onClick={() => {
            if (confirm("編集内容は破棄されます。本当にキャンセルしますか？")) {
              setJournalTitle("");
              setJournalContent("");
              handleClose();
            }
          }}
        >
          キャンセル
        </Button>
        <Button
          type="submit"
          onClick={async () => {
            if (!existingJournal) {
              await createJournal({
                title: journalTitle,
                content: journalContent,
                eventId: schedule.id,
                date: schedule.start.toISOString(),
              });
            } else {
              await updateJournal({
                ...existingJournal,
                title: journalTitle,
                content: journalContent,
              });
            }
            setJournalTitle("");
            setJournalContent("");
            handleClose();
          }}
        >
          完了
        </Button>
      </DialogActions>
    </Dialog>
  );
}
