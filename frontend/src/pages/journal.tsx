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
import AddIcon from "@mui/icons-material/Add";
import { useCallback, useEffect, useRef, useState } from "react";
import JournalEntry from "../components/JournalEntry";
import useJournal from "../hooks/useJournal";
import { add } from "date-fns";
import TopBar from "../components/TopBar";

function Journal() {
  const { journals, isLoading, error, fetchMoreJournalsAfter, fetchMoreJournalsBefore } =
    useJournal();
  const [open, setOpen] = useState(false);
  const [journalTitle, setJournalTitle] = useState("");
  const [journalContent, setJournalContent] = useState("");
  const [baseDate, setBaseDate] = useState(new Date());
  const [topDate, setTopDate] = useState<Date>(new Date(new Date("2024-09-17").toDateString()));
  const [bottomDate, setBottomDate] = useState<Date>(
    new Date(new Date("2024-09-17").toDateString())
  );
  const topTargetRef = useRef<HTMLDivElement>(null);
  const bottomTargetRef = useRef<HTMLDivElement>(null);

  const topScrollObserver = useCallback(
    () =>
      new IntersectionObserver(
        async (entries) => {
          if (entries[0].isIntersecting) {
            await fetchMoreJournalsBefore(topDate);
            setTopDate((prev) => add(prev, { days: -4 }));
          }
        },
        {
          root: null,
          rootMargin: "200px",
          threshold: 0.01,
        }
      ),
    [fetchMoreJournalsBefore, topDate]
  );

  const bottomScrollObserver = useCallback(
    () =>
      new IntersectionObserver(
        async (entries) => {
          if (entries[0].isIntersecting) {
            await fetchMoreJournalsAfter(bottomDate);
            setBottomDate((prev) => add(prev, { days: 4 }));
          }
        },
        {
          root: null,
          rootMargin: "200px",
          threshold: 0.01,
        }
      ),
    [fetchMoreJournalsAfter, bottomDate]
  );

  useEffect(() => {
    const topTarget = topTargetRef.current;
    if (topTarget) {
      const topObserver = topScrollObserver();
      topObserver.observe(topTarget);
      return () => {
        topObserver.unobserve(topTarget);
      };
    }
  }, [topScrollObserver, topTargetRef]);

  useEffect(() => {
    const bottomTarget = bottomTargetRef.current;
    if (bottomTarget) {
      const bottomObserver = bottomScrollObserver();
      bottomObserver.observe(bottomTarget);
      return () => {
        bottomObserver.unobserve(bottomTarget);
      };
    }
  }, [bottomScrollObserver, bottomTargetRef]);

  return (
    <div className="journal-app">
      <TopBar baseDate={baseDate} setBaseDate={setBaseDate} />
      <div ref={topTargetRef} style={{ marginTop: "20%" }} />
      {isLoading ? (
        <div>Loading...</div>
      ) : error ? (
        <div>Error: {error.message}</div>
      ) : (
        <div className="journal-entries">
          {journals?.map((journal, index) => <JournalEntry key={index} journal={journal} />)}
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
      <div ref={bottomTargetRef} />
    </div>
  );
}

export default Journal;
