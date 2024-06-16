import React, { useState } from "react";
import InfiniteScroll from "react-infinite-scroller";
import axios from "axios";
import { Box, Typography, Button, Modal, TextField } from "@mui/material";
import TopBar from "./TopBar";

const style = {
  position: "absolute",
  top: "50%",
  left: "50%",
  transform: "translate(-50%, -50%)",
  width: 400,
  bgcolor: "background.paper",
  border: "2px solid #000",
  boxShadow: 24,
  p: 4,
};

interface Post {
  id: number;
  title: string;
  body: string;
  color: string; // Added color property to Post interface
}

const ScheduleView: React.FC = () => {
  const [items, setItems] = useState<Post[]>([]);
  const [page, setPage] = useState<number>(1);
  const [selectedEvent, setSelectedEvent] = useState<Post | null>(null);
  const [editedTitle, setEditedTitle] = useState<string>("");
  const [baseDate, setBaseDate] = useState<Date>(new Date());

  const fetchData = async (__page: number) => {
    console.log(__page);
    const response = await axios.get<Post[]>(
      `https://jsonplaceholder.typicode.com/posts?_page=${page}&_limit=10`
    );
    setItems([
      ...items,
      ...response.data.map((item) => ({
        ...item,
        color: getEventColor(item.title),
      })),
    ]);
    setPage(page + 1);
  };

  const currentDate = new Date("2024-04-16");

  const getEventColor = (title: string) => {
    const firstLetter = title.charAt(0).toLowerCase();
    if (firstLetter >= "a" && firstLetter <= "j") {
      return "#ff69b4"; // Pink
    } else if (firstLetter >= "k" && firstLetter <= "t") {
      return "#90ee90"; // Light green
    } else {
      return "#ffa500"; // Orange
    }
  };

  const handleOpenModal = (event: Post) => {
    setSelectedEvent(event);
    setEditedTitle(event.title);
  };

  const handleCloseModal = () => {
    setSelectedEvent(null);
  };

  const handleSaveTitle = () => {
    if (selectedEvent) {
      setSelectedEvent({ ...selectedEvent, title: editedTitle });
      setItems(
        items.map((item) =>
          item.id === selectedEvent.id ? { ...item, title: editedTitle } : item
        )
      );
      handleCloseModal();
    }
  };

  return (
    <div style={{ display: "flex", flexDirection: "column", height: "100vh" }}>
      <TopBar baseDate={baseDate} setBaseDate={setBaseDate} />
      <div style={{ display: "flex", flexGrow: 1 }}>
        <div style={{ marginRight: "20px" }}>
          <Typography variant="h6" style={{ color: "gray" }}>
            {currentDate.toLocaleDateString(undefined, { weekday: "short" })}
            <br />
            {currentDate.toLocaleDateString(undefined, { day: "numeric" })}
          </Typography>
        </div>
        <InfiniteScroll
          style={{ flexGrow: 1 }}
          pageStart={0}
          loadMore={() => fetchData(page)}
          hasMore={true}
          loader={
            <div className="loader" key={0}>
              Loading ...
            </div>
          }
        >
          {items.map((item) => (
            <Box
              key={item.id}
              sx={{
                padding: "5px",
                width: "90%",
                borderRadius: "5px",
                backgroundColor: item.color,
                color: "white",
                marginBottom: "10px",
                display: "flex",
                flexDirection: "column",
                alignItems: "flex-start",
              }}
            >
              <Typography variant="caption">{item.title}</Typography>
              <Typography variant="caption">9:00 ~ 10:00</Typography>
              <Button onClick={() => handleOpenModal(item)}>
                View Details
              </Button>
            </Box>
          ))}
        </InfiniteScroll>
      </div>
      {selectedEvent && (
        <Modal
          open={true}
          onClose={handleCloseModal}
          aria-labelledby="modal-modal-title"
          aria-describedby="modal-modal-description"
        >
          <Box sx={{ ...style, bgcolor: selectedEvent.color }}>
            <TextField
              label="Title"
              value={editedTitle}
              onChange={(e) => setEditedTitle(e.target.value)}
              fullWidth
              sx={{ mb: 2 }}
            />
            <Button
              onClick={handleSaveTitle}
              variant="contained"
              color="primary"
            >
              Save
            </Button>
          </Box>
        </Modal>
      )}
    </div>
  );
};

export default ScheduleView;
