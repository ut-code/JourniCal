import { useEffect, useState } from "react";
import InfiniteScroll from "react-infinite-scroller";
import {
  Box,
  Table,
  TableCell,
  TableRow,
  Typography,
} from "@mui/material";
import TimelineSchedule from "./TimelineSchedule";
import { add } from "date-fns";
import { Schedule } from "../types/types";
import useSWR from "swr";

type ScheduleViewProps = {
  day: Date;
  today: Date;
};

type FetchedSchedule = {
  id: string;
  colorId?: string;
  start: {
    date?: string;
    dateTime?: string;
  };
  end: {
    date?: string;
    dateTime?: string;
  };
  summary: string;
};

const COLOR_DICT = [
  "#7986CB",
  "#33B679",
  "#8E24AA",
  "#E67C73",
  "#F6BF26",
  "#F4511E",
  "#039BE5",
  "#616161",
  "#3F51B5",
  "#0B8043",
  "#D50000",
];

const scheduleFromFetchedData = (
  fetchedSchedule: FetchedSchedule,
): Schedule => {
  if (
    fetchedSchedule.start.dateTime == undefined &&
    fetchedSchedule.start.date != undefined &&
    fetchedSchedule.end.dateTime == undefined &&
    fetchedSchedule.end.date != undefined
  ) {
    return {
      id: fetchedSchedule.id,
      isAllDay: true,
      start: new Date(fetchedSchedule.start.date),
      end: new Date(fetchedSchedule.end.date),
      title: fetchedSchedule.summary,
      color:
        COLOR_DICT[
          Number(fetchedSchedule.colorId)
            ? Number(fetchedSchedule.colorId) - 1
            : 6
        ],
    };
  }
  if (
    fetchedSchedule.start.dateTime != undefined &&
    fetchedSchedule.start.date == undefined &&
    fetchedSchedule.end.dateTime != undefined &&
    fetchedSchedule.end.date == undefined
  ) {
    return {
      id: fetchedSchedule.id,
      isAllDay: false,
      start: new Date(fetchedSchedule.start.dateTime),
      end: new Date(fetchedSchedule.end.dateTime),
      title: fetchedSchedule.summary,
      color:
        COLOR_DICT[
          Number(fetchedSchedule.colorId)
            ? Number(fetchedSchedule.colorId) - 1
            : 6
        ],
    };
  }
  throw new Error("invalid schedule format.");
};

const isEqualDay = (day1: Date, day2: Date) => {
  return (
    day1.getFullYear() === day2.getFullYear() &&
    day1.getMonth() === day2.getMonth() &&
    day1.getDate() === day2.getDate()
  );
};

// const style = {
//   position: "absolute",
//   top: "50%",
//   left: "50%",
//   transform: "translate(-50%, -50%)",
//   width: 400,
//   bgcolor: "background.paper",
//   border: "2px solid #000",
//   boxShadow: 24,
//   p: 4,
// };

// interface Post {
//   id: number;
//   title: string;
//   body: string;
//   color: string; // Added color property to Post interface
// }

const ScheduleView = (props: ScheduleViewProps): JSX.Element => {
  // const [items, setItems] = useState<Post[]>([]);
  // const [page, setPage] = useState<number>(1);
  // const [selectedEvent, setSelectedEvent] = useState<Post | null>(null);
  // const [editedTitle, setEditedTitle] = useState<string>("");
  // const [baseDate, setBaseDate] = useState<Date>(new Date());

  const { day, today } = props;
  const DAYOFWEEK = ["日", "月", "火", "水", "木", "金", "土"];
  const isToday = isEqualDay(day, today);
  const [schedules, setSchedules] = useState<Schedule[]>([]);
  const [currentDate, setCurrentDate] = useState<Date>(day);

  // fetch data and increase page number
  const startUnixTime = Math.floor(day.getTime() / 1000);
  const endUnixTime = Math.floor(add(day, { days: 1 }).getTime() / 1000);
  const { data, error } = useSWR(
    `http://localhost:3000/api/calendar/get-events-in-range/${startUnixTime}/${endUnixTime}`,
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
    console.log(currentDate);
    if (data) {
      const newSchedules = data.map((schedule: FetchedSchedule) =>
        scheduleFromFetchedData(schedule),
      );
      console.log(newSchedules);
      setSchedules((previousSchedule) => [...previousSchedule, ...newSchedules]);
      setCurrentDate((prevDage) => add(prevDage, { days: 1 }));
    }
  }, [data, currentDate]);
  const loadMore = () => {
    setCurrentDate((prevDate) => add(prevDate, { days: 1 }));
  };

  // const fetchData = async (__page: number) => {
  //   console.log(__page);
  //   const response = await axios.get<Post[]>(
  //     `https://jsonplaceholder.typicode.com/posts?_page=${page}&_limit=10`,
  //   );
  //   setItems([
  //     ...items,
  //     ...response.data.map((item) => ({
  //       ...item,
  //       color: getEventColor(item.title),
  //     })),
  //   ]);
  //   setPage(page + 1);
  // };

  // const currentDate = new Date("2024-04-16");

  // const getEventColor = (title: string) => {
  //   const firstLetter = title.charAt(0).toLowerCase();
  //   if (firstLetter >= "a" && firstLetter <= "j") {
  //     return "#ff69b4"; // Pink
  //   } else if (firstLetter >= "k" && firstLetter <= "t") {
  //     return "#90ee90"; // Light green
  //   } else {
  //     return "#ffa500"; // Orange
  //   }
  // };

  // const handleOpenModal = (event: Post) => {
  //   setSelectedEvent(event);
  //   setEditedTitle(event.title);
  // };

  // const handleCloseModal = () => {
  //   setSelectedEvent(null);
  // };

  // const handleSaveTitle = () => {
  //   if (selectedEvent) {
  //     setSelectedEvent({ ...selectedEvent, title: editedTitle });
  //     setItems(
  //       items.map((item) =>
  //         item.id === selectedEvent.id ? { ...item, title: editedTitle } : item,
  //       ),
  //     );
  //     handleCloseModal();
  //   }
  // };

  return (
    <Table
      sx={{
        height: "90vh",
        minHeight: "1150px",
        borderRight: "1px solid gainsboro",
      }}
    >
      <InfiniteScroll
          style={{ flexGrow: 1 }}
          pageStart={0}
          loadMore={loadMore}
          hasMore={true}
          loader={
            <div className="loader" key={0}>
              Loading ...
            </div>
          }
        >
        <Box
          display={"flex"}
          flexDirection={"column"}
          justifyContent={"center"}
          alignItems={"center"}
          height={"3vh"}
        >
          <Typography>{DAYOFWEEK[day.getDay()]}</Typography>
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
                {day.getDate()}
              </Typography>
            </Box>
          ) : (
            <Typography variant="h5">{day.getDate()}</Typography>
          )}
        </Box>
        {schedules
          .filter((schedule) => schedule.isAllDay)
          .map((schedule) => (
            <TableRow key={schedule.id}>
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
          ))}
        {schedules
          .filter((schedule) => !schedule.isAllDay)
          .map((schedule) => (
            <TimelineSchedule key={schedule.id} schedule={schedule} />
          ))}
      </InfiniteScroll>
    </Table>
  );
  // return (
  //   <div style={{ display: "flex", flexDirection: "column", height: "100vh" }}>
  //     <TopBar baseDate={baseDate} setBaseDate={setBaseDate} />
  //     <div style={{ display: "flex", flexGrow: 1 }}>
  //       <div style={{ marginRight: "20px" }}>
  //         <Typography variant="h6" style={{ color: "gray" }}>
  //           {currentDate.toLocaleDateString(undefined, { weekday: "short" })}
  //           <br />
  //           {currentDate.toLocaleDateString(undefined, { day: "numeric" })}
  //         </Typography>
  //       </div>
  //     </div>
  //     {selectedEvent && (
  //       <Modal
  //         open={true}
  //         onClose={handleCloseModal}
  //         aria-labelledby="modal-modal-title"
  //         aria-describedby="modal-modal-description"
  //       >
  //         <Box sx={{ ...style, bgcolor: selectedEvent.color }}>
  //           <TextField
  //             label="Title"
  //             value={editedTitle}
  //             onChange={(e) => setEditedTitle(e.target.value)}
  //             fullWidth
  //             sx={{ mb: 2 }}
  //           />
  //           <Button
  //             onClick={handleSaveTitle}
  //             variant="contained"
  //             color="primary"
  //           >
  //             Save
  //           </Button>
  //         </Box>
  //       </Modal>
  //     )}
  //   </div>
  // );
};

export default ScheduleView;
