import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Home from "./pages/Home";
import Page1 from "./pages/page1";
import Page2 from "./pages/page2";
import { CssBaseline } from "@mui/material";
import WeekViewTest from "./pages/WeekViewTest";
import Diary from "./pages/diary";

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <Home />,
    },
    {
      path: "/page1",
      element: <Page1 />,
    },
    {
      path: "/page2",
      element: <Page2 />,
    },
    {
      path: "/week_view_test",
      element: <WeekViewTest />,
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
