import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Home from "./pages/Home";
import { CssBaseline } from "@mui/material";
import Calendar from "./pages/WeekViewTest";
import Diary from "./pages/diary";

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <Home />,
    },
    {
      path: "/calendar",
      element: <Calendar />,
    },
    {
      path: "/diary",
      element: <Diary />,
    },
  ]);

  return (
    <>
      <CssBaseline />
      <RouterProvider router={router} />
    </>
  );
}

export default App;
