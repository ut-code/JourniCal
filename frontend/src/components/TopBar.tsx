import { Box, Button, Collapse, Typography } from "@mui/material";
import React, { useState } from "react";
import { Link } from "react-router-dom";
import MenuIcon from "@mui/icons-material/Menu";
import CalendarMonthIcon from "@mui/icons-material/CalendarMonth";
import ExpandLessIcon from "@mui/icons-material/ExpandLess";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import SearchIcon from "@mui/icons-material/Search";
import CheckCircleOutlineIcon from "@mui/icons-material/CheckCircleOutline";
import AutoStoriesIcon from "@mui/icons-material/AutoStories";
import DateRangeIcon from "@mui/icons-material/DateRange";
import { useLocation } from "react-router-dom";
import { CALENDAR_PATH_NAME, DIARY_PATH_NAME } from "../consts/consts";
import { DateCalendar, LocalizationProvider } from "@mui/x-date-pickers";
import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFnsV3";

const TopBar: React.FC<{
  baseDate: Date;
  setBaseDate: React.Dispatch<React.SetStateAction<Date>>;
}> = (props) => {
  const { baseDate, setBaseDate } = props;
  const [isTopCalendarOpen, setIsTopCalendarOpen] = useState(false);
  const currentPathName = useLocation().pathname;
  const iconCommonSxProps = { mx: 1.5, color: "primary.contrastText" };
  const linkIconCommonStyleProps = { paddingTop: "3%" };

  return (
    <Box width="100vw" sx={{ position: "fixed", zIndex: 1, top: 0 }}>
      <Box
        width={"100%"}
        display={"flex"}
        alignItems={"center"}
        justifyContent={"space-between"}
        sx={{ bgcolor: "primary.main" }}
      >
        <MenuIcon sx={{ mx: 2, color: "primary.contrastText" }} />
        <Button
          variant="contained"
          color="primary"
          onClick={async () => {
            try {
              await fetch("http://localhost:3000/api/diaries/", {
                method: "POST",
                headers: {
                  "Content-Type": "application/json",
                },
                body: JSON.stringify({
                  title: "some title",
                  content: "some content",
                  eventId: 1,
                }),
              });
            } catch (error) {
              console.error(error);
            }
          }}
        >
          テスト: 新規作成
        </Button>
        <Box
          display={"flex"}
          alignItems={"center"}
          onClick={() => setIsTopCalendarOpen(!isTopCalendarOpen)}
        >
          <Typography
            variant="h5"
            component="div"
            sx={{ my: 2, color: "primary.contrastText" }}
          >
            {baseDate.getMonth() + 1}月
          </Typography>
          {isTopCalendarOpen ? (
            <ExpandLessIcon sx={{ color: "primary.contrastText" }} />
          ) : (
            <ExpandMoreIcon sx={{ color: "primary.contrastText" }} />
          )}
        </Box>
        <Box display={"flex"} alignItems={"center"} sx={{ ml: "auto" }}>
          <CalendarMonthIcon sx={iconCommonSxProps} />
          <SearchIcon sx={iconCommonSxProps} />
          <CheckCircleOutlineIcon sx={iconCommonSxProps} />
          {currentPathName === DIARY_PATH_NAME ? (
            <Link to={CALENDAR_PATH_NAME} style={linkIconCommonStyleProps}>
              <DateRangeIcon sx={iconCommonSxProps} />
            </Link>
          ) : (
            currentPathName === CALENDAR_PATH_NAME && (
              <Link to={DIARY_PATH_NAME} style={linkIconCommonStyleProps}>
                <AutoStoriesIcon sx={iconCommonSxProps} />
              </Link>
            )
          )}
        </Box>
      </Box>
      <Collapse in={isTopCalendarOpen} sx={{ backgroundColor: "gainsboro" }}>
        <LocalizationProvider dateAdapter={AdapterDateFns}>
          <DateCalendar
            value={baseDate}
            onChange={(value) => setBaseDate(value)}
          />
        </LocalizationProvider>
      </Collapse>
    </Box>
  );
};

export default TopBar;
