import { Box, Table, TableBody, TableRow, Typography } from "@mui/material";

const TimelineRowName = (): JSX.Element => {
  return (
    // TODO 時刻がずれるのをなおす
    <Table
      sx={{
        mx: 1,
        width: "10%",
        height: "89.7vh",
        minHeight: "1150px",
        transform: "translate(0, -0.8%)",
      }}
    >
      <Box height="3vh" />
      <TableBody>
        {[...Array(24).keys()].map((i) => (
          <TableRow key={i}>
            <Typography variant="caption">{i}:00</Typography>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};

export default TimelineRowName;
