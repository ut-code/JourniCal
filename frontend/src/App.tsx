import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Home from "./pages/Home";
import { CssBaseline } from "@mui/material";
import Calendar from "./pages/Calendar";
import Journal from "./pages/journal";

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
      path: "/journal",
      element: <Journal />,
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
