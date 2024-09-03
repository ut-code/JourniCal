import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  Fab,
  TextField,
} from "@mui/material";
import JournalEntry from "../components/JournalEntry";
import useJournal from "../hooks/useJournal";
import AddIcon from "@mui/icons-material/Add";
import { useState } from "react";

function Journal() {
  const { journals, isLoading, error } = useJournal();
  const [open, setOpen] = useState(false);
  const [journalTitle, setJournalTitle] = useState("");
  const [journalContent, setJournalContent] = useState("");
  return (
    <div className="journal-app">
      {isLoading ? (
        <div>Loading...</div>
      ) : error ? (
        <div>Error: {error.message}</div>
      ) : (
        <div className="journal-entries">
          {journals?.map((journal, index) => (
            <JournalEntry key={index} journal={journal} />
          ))}
        </div>
      )}
      <Box
        sx={{
          position: "fixed",
          bottom: 16,
          right: 16,
        }}
      >
        <Fab color="primary" aria-label="add" onClick={() => setOpen(true)}>
          <AddIcon />
        </Fab>
      </Box>
      <Dialog
        open={open}
        onClose={() => {
          setOpen(false);
        }}
      >
        <DialogTitle>ジャーナルの追加</DialogTitle>
        <DialogContent>
          <DialogContentText>ジャーナルを追加します。</DialogContentText>
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
              setOpen(false);
            }}
          >
            キャンセル
          </Button>
          <Button
            type="submit"
            onClick={async () => {
              fetch(`http://localhost:3000/api/journals/`, {
                method: "POST",
                headers: {
                  "Content-Type": "application/json",
                },
                body: JSON.stringify({
                  title: journalTitle,
                  content: journalContent,
                }),
              });
              setOpen(false);
            }}
          >
            登録
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}

export default Journal;
