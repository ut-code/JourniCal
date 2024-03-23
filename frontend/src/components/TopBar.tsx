import { Box, Typography } from "@mui/material";
import React from "react";
import { Link } from "react-router-dom";
import MenuIcon from "@mui/icons-material/Menu";
import CalendarMonthIcon from "@mui/icons-material/CalendarMonth";
import SearchIcon from "@mui/icons-material/Search";
import CheckCircleOutlineIcon from "@mui/icons-material/CheckCircleOutline";
import AutoStoriesIcon from "@mui/icons-material/AutoStories";
import DateRangeIcon from "@mui/icons-material/DateRange";
import { useLocation } from "react-router-dom";

const TopBar: React.FC<{
  journalPathName: string;
  calendarPathName: string;
}> = (props) => {
  const currentPathName = useLocation().pathname;
  const iconCommonSxProps = { mx: 1.5, color: "primary.contrastText" };
  const linkIconCommonStyleProps = { paddingTop: "3%" };

  return (
    <Box
      width={"100%"}
      display={"flex"}
      alignItems={"center"}
      justifyContent={"space-between"}
      sx={{ bgcolor: "primary.main" }}
    >
      <MenuIcon sx={{ mx: 2, color: "primary.contrastText" }} />
      <Typography
        variant="h5"
        component="div"
        sx={{ my: 2, color: "primary.contrastText" }}
      >
        Month
      </Typography>
      <Box display={"flex"} alignItems={"center"} sx={{ ml: "auto" }}>
        <CalendarMonthIcon sx={iconCommonSxProps} />
        <SearchIcon sx={iconCommonSxProps} />
        <CheckCircleOutlineIcon sx={iconCommonSxProps} />
        {currentPathName === props.journalPathName ? (
          <Link to={props.calendarPathName} style={linkIconCommonStyleProps}>
            <DateRangeIcon sx={iconCommonSxProps} />
          </Link>
        ) : (
          currentPathName === props.calendarPathName && (
            <Link to={props.journalPathName} style={linkIconCommonStyleProps}>
              <AutoStoriesIcon sx={iconCommonSxProps} />
            </Link>
          )
        )}
      </Box>
    </Box>
  );
};

export default TopBar;
